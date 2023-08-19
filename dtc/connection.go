package dtc

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	//"encoding/json"
	"reflect"

	"google.golang.org/protobuf/proto" // TODO: use the other one! above^

	//"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/encoding/protojson"
	//"google.golang.org/protobuf/reflect/protoregistry"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/RileyR387/sc-dtc-client/dtcproto"
	//"github.com/RileyR387/sc-dtc-client/marketdata"
	"github.com/RileyR387/sc-dtc-client/accounts"
	"github.com/RileyR387/sc-dtc-client/securities"
	ttr "github.com/RileyR387/sc-dtc-client/termtrader"
	wsp "github.com/RileyR387/sc-dtc-client/web"
)

const DTC_CLIENT_HEARTBEAT_SECONDS = 10

// const CONN_BUFFER_SIZE = 4096
// const CONN_BUFFER_SIZE = 4096*2
const CONN_BUFFER_SIZE = 131072 // 4096*32

type DtcConnection struct {
	connArgs              ConnectArgs
	connUri               string
	conn                  net.Conn
	socketReadMutex       sync.Mutex
	connected             bool
	reconnecting          bool
	terminated            bool
	reader                io.Reader
	requestId             int32
	clientClose           chan int
	listenClose           chan int
	lastHeartbeatResponse int64
	heartbeatMtx          sync.Mutex
	heartbeatUpdate       chan *Heartbeat
	marketData            chan securities.MarketDataUpdate
	subscribers           []ClientPlugin
	securityStore         *securities.SecurityStore
	accountStore          *accounts.AccountStore
	keepingAlive          bool
	keepingAliveMtx       sync.Mutex
	listening             bool
	loggedOn              bool
	backOffRate           int
}

func init() {
}

func (d *DtcConnection) _Listen() {
	d.listening = true
	log.Infof("Client listener started")
	for {
		select {
		case <-d.listenClose:
			d.listening = false
			log.Warn("Closing client listener")
			return
		default:
			d._RouteMessage(d._NextMessage())
		}
	}
}

func (d *DtcConnection) Connect(c ConnectArgs) error {
	log.Infof("Connecting: %s@%s:%s\n", c.Username, c.Host, c.Port)
	uri := net.JoinHostPort(c.Host, c.Port)
	dialer := net.Dialer{Timeout: 4 * time.Second}

	conn, err := dialer.Dial("tcp", uri)

	if err != nil {
		if c.Reconnect && d.backOffRate > 1 {
			log.Warnf("Failed to connect to DTC server: %v\n", err)
			return err
		} else {
			log.Fatalf("Failed to connect to DTC server: %v\n", err)
			os.Exit(1)
		}
	}

	d.conn = conn
	d.reader = bufio.NewReaderSize(d.conn, CONN_BUFFER_SIZE)
	d.connArgs = c
	d.connUri = uri

	d._SetEncoding()
	err = d._Logon()
	if err != nil {
		d.Disconnect()
		return err
	}
	d.loggedOn = true
	d.connected = true

	d.securityStore = securities.NewSecurityStore()
	d.accountStore = accounts.NewAccountStore()
	d.clientClose = make(chan int)
	d.listenClose = make(chan int)
	d.heartbeatUpdate = make(chan *Heartbeat)
	d.marketData = make(chan securities.MarketDataUpdate)
	d.keepingAlive = true
	if d.backOffRate < 2 {
		d.backOffRate = 1
	}
	go d._ReceiveHeartbeat()
	go d._Listen()
	go d.keepAlive()
	go d.startSubscriptionRouter()
	go d.initTrading()
	d.reconnecting = false

	return nil
}

/**
* Test bed for now.
 */
func (d *DtcConnection) initTrading() error {
	time.Sleep(5 * time.Second)
	err := d.LoadAccounts()
	if err != nil {
		log.Fatalf("Failed to load accounts with error: %v", err)
	}
	time.Sleep(5 * time.Second)
	err = d.AccountBlanaceRefresh(os.Getenv("SC_DTC_CLIENT__DEFAULT_TRADE_ACCOUNT"))
	if err != nil {
		log.Fatalf("Failed to load account balances with error: %v", err)
	}
	return err
}

func (d *DtcConnection) _Reconnect() error {
	return nil
}

func (d *DtcConnection) addSecurity(def *dtcproto.SecurityDefinitionResponse) {
	log.Infof("Added security %v from exchange %v as %v with ID: %v", def.ExchangeSymbol, def.Exchange, def.Symbol, def.RequestID)
	d.securityStore.AddSecurity(securities.FromDef(def))
}

func (d *DtcConnection) Disconnect() {

	d.keepingAliveMtx.Lock()
	if d.keepingAlive {
		d.clientClose <- 0
	}
	d.keepingAliveMtx.Unlock()
	if d.loggedOn {
		d._Logoff()
	}

	if d.connected {
		d.conn.Close()
		log.Info("Connection closed")
	}

	if d.listening {
		go d.closeListener()
		time.Sleep(time.Second)
	}

	d.connected = false
	if !d.connArgs.Reconnect {
		d.terminated = true
	}
	log.Info("Disconnected")
}

func (d *DtcConnection) closeListener() {
	d.listenClose <- 0
}

func (d *DtcConnection) keepAlive() {
	lastSend := time.Now()
	for {
		select {
		case <-d.clientClose:
			d.keepingAliveMtx.Lock()
			defer d.keepingAliveMtx.Unlock()
			d.keepingAlive = false
			log.Debugf("Client Heartbeat closed\n")
			return
		default:
			time.Sleep(200 * time.Millisecond)
			if time.Since(lastSend).Seconds() > DTC_CLIENT_HEARTBEAT_SECONDS {
				if d.connected {
					d._SendHeartbeat()
				} else {
					d.clientClose <- 0
				}
				lastSend = time.Now()
			}
		}
	}
}

func (d *DtcConnection) startSubscriptionRouter() {
	var msg securities.MarketDataUpdate
	d.subscribers = []ClientPlugin{
		ttr.New(d.securityStore, d.accountStore),
		wsp.New(d.securityStore, d.accountStore),
	}

	for {
		select {
		case msg = <-d.marketData:
			d._UpdateLastHeartBeat()
			d.securityStore.AddData(msg)
			// Distribute Market data to other subscribers
			for _, subscriber := range d.subscribers {
				go subscriber.ReceiveData(msg)
			}
		}
	}
}

func (d *DtcConnection) _UpdateLastHeartBeat() {
	d.heartbeatMtx.Lock()
	defer d.heartbeatMtx.Unlock()
	d.lastHeartbeatResponse = time.Now().Unix()
}

func (d *DtcConnection) _GetLastHeartbeat() int64 {
	d.heartbeatMtx.Lock()
	defer d.heartbeatMtx.Unlock()
	return d.lastHeartbeatResponse
}

func (d *DtcConnection) _ReceiveHeartbeat() {
	var m *Heartbeat
	d._UpdateLastHeartBeat()
	for {
		select {
		case m = <-d.heartbeatUpdate:
			lastHb := d._GetLastHeartbeat()
			if m.CurrentDateTime-lastHb > DTC_CLIENT_HEARTBEAT_SECONDS+5 {
				log.Warnf("Received late heartbeat, span from last: %v", m.CurrentDateTime-lastHb)
			} else {
				//log.Trace("Received sane heartbeat")
			}
			d._UpdateLastHeartBeat()
			//d.lastHeartbeatResponse = time.Now().Unix()
			//d.lastHeartbeatResponse = m.CurrentDateTime
			//log.Tracef("Updated last heartbeat to %v", m.CurrentDateTime)
		default:
			lastHb := d._GetLastHeartbeat()
			time.Sleep(time.Duration(d.backOffRate) * 400 * time.Millisecond)
			if time.Now().Unix()-lastHb > DTC_CLIENT_HEARTBEAT_SECONDS*3 {
				log.Warnf("No server heartbeat received in %v seconds", time.Now().Unix()-lastHb)
				if d.listening && !d.connArgs.Reconnect {
					d.Disconnect()
				} else {
					if d.loggedOn {
						d._Logoff()
					}
					if d.listening {
						go d.closeListener()
						time.Sleep(time.Second)
					}
				}
				if !d.listening && !d.connArgs.Reconnect {
					log.Warn("No longer listening, closing hearbeat listening thread")
					return
				} else {
					log.Warn("Awaiting logoff for reconnect...")
					time.Sleep(time.Duration(d.backOffRate*10) * time.Second)
					log.Warn("Attempting reconnect...")
					d.reconnecting = true
					d.backOffRate = d.backOffRate * 2
					d.Connect(d.connArgs)
					return
				}
			}
		}
	}
}

func (d *DtcConnection) _Logon() error {
	logonRequest := dtcproto.LogonRequest{
		Username:                   d.connArgs.Username,
		Password:                   d.connArgs.Password,
		Integer_1:                  2,
		HeartbeatIntervalInSeconds: DTC_CLIENT_HEARTBEAT_SECONDS,
		ClientName:                 "go-dtc",
		ProtocolVersion:            dtcproto.DTCVersion_value["CURRENT_VERSION"],
	}
	//describe( logonRequest.ProtoReflect().Descriptor().FullName() )
	fmt.Println(protojson.Format(&logonRequest))

	msg, err := proto.Marshal(&logonRequest)
	if err != nil {
		log.Fatalf("Failed to marshal LogonRequest message: %v\n", err)
		os.Exit(1)
	}

	log.Debug("Sending LOGON_REQUEST")
	d.conn.Write(PackMessage(msg, dtcproto.DTCMessageType_value["LOGON_REQUEST"]))

	resp, mTypeId := d._GetMessage()

	logonResponse := dtcproto.LogonResponse{}
	if err := proto.Unmarshal(resp, &logonResponse); err != nil {
		log.Fatalln("Failed to parse LogonResponse:", err)
	}
	if logonResponse.Result != dtcproto.LogonStatusEnum_LOGON_SUCCESS {
		/*
		   log.WithFields(log.Fields{
		       "result": logonResponse.Result,
		       "desc": logonResponse.ResultText,
		   }).Fatal("Logon Failed")
		*/
		log.Fatalf("Logon Failed with result %v and text %v", logonResponse.Result, logonResponse.ResultText)
		return errors.New("Logon Failure")
	}
	d._UpdateLastHeartBeat()
	log.Debugf("Received %v result: %v", dtcproto.DTCMessageType_name[mTypeId], logonResponse.ResultText)
	fmt.Println(protojson.Format(&logonResponse))
	return nil
}

func (d *DtcConnection) _Logoff() {
	logoff := dtcproto.Logoff{
		Reason:         "Done",
		DoNotReconnect: 1,
	}
	msg, err := proto.Marshal(&logoff)
	if err != nil {
		log.Fatalf("Failed to marshal LogonRequest message: %v\n", err)
		os.Exit(1)
	}

	log.Debug("Sending LOGOFF")
	d.conn.Write(PackMessage(msg, dtcproto.DTCMessageType_value["LOGOFF"]))
	log.Trace("Logoff request sent")
}

func (d *DtcConnection) _SetEncoding() {
	encodingReq := dtcproto.EncodingRequest{
		ProtocolVersion: dtcproto.DTCVersion_value["CURRENT_VERSION"],
		Encoding:        dtcproto.EncodingEnum_PROTOCOL_BUFFERS,
		ProtocolType:    "DTC",
	}
	msg := dtc_bin_encoder(encodingReq)
	log.Debug("Sending ENCODING_REQUEST")
	d.conn.Write(PackMessage(msg, dtcproto.DTCMessageType_value["ENCODING_REQUEST"]))
	respBin, mTypeId := d._GetMessage()
	log.Debugf("Received %v(%v) with bytes %v", dtcproto.DTCMessageType_name[mTypeId], mTypeId, respBin)
	/**
	   * TODO: Handle binary encoding response for log purposes
	  resp := d._GetMessage()
	*/
}

func (d *DtcConnection) _SendHeartbeat() {
	heartbeat := Heartbeat{
		NumDroppedMessages: 0,
		CurrentDateTime:    time.Now().Unix(),
	}
	msg, err := proto.Marshal(&heartbeat)
	if err != nil {
		log.Fatalf("Failed to marshal Heartbeat message: %v\n", err)
		os.Exit(1)
	}
	d.conn.Write(PackMessage(msg, dtcproto.DTCMessageType_value["HEARTBEAT"]))
}

func PackMessage(msg []byte, mTypeId int32) []byte {
	length := 4 + len(msg)
	header := make([]byte, 4)
	binary.LittleEndian.PutUint16(header[0:2], uint16(length))
	binary.LittleEndian.PutUint16(header[2:4], uint16(mTypeId))
	message := append(header, msg...)
	if dtcproto.DTCMessageType_name[mTypeId] != "HEARTBEAT" && dtcproto.DTCMessageType_name[mTypeId] != "LOGON_REQUEST" {
		log.Tracef("Packed message with TypeID: %v (%v) with length %v and contents: 0x%x", mTypeId, dtcproto.DTCMessageType_name[mTypeId], length, message)
	}
	return message
}

func (d *DtcConnection) _NextMessage() (proto.Message, reflect.Type, int32) {
	bytes, mTypeId := d._GetMessage()
	m, t := d._ParseMessage(bytes, mTypeId)
	return m, t, mTypeId
}

func (d *DtcConnection) _GetMessage() ([]byte, int32) {
	d.socketReadMutex.Lock()
	defer d.socketReadMutex.Unlock()
	length, mTypeId := _ParseHeaderBytes(d.reader)

	if length == 0 {
		log.Warnf("Received %v(%v) with byte length %v", dtcproto.DTCMessageType_name[mTypeId], mTypeId, length)
	} else if log.GetLevel() == log.TraceLevel {
		//log.Tracef("Received %v(%v) with byte length %v", dtcproto.DTCMessageType_name[mTypeId], mTypeId, length )
	}

	resp := make([]byte, length)
	_, err := io.ReadFull(d.reader, resp)

	if err != nil {
		switch t := err.(type) {
		case *net.OpError:
			if t.Op == "read" {
				log.Trace("Reader closed")
			} else {
				log.Errorf("Message didn't fill buffer of %d bytes with error: %v\n", length, err)
			}
			return nil, 0
		default:
			if err == io.EOF {
				log.Warn("Received end of file on communication channel. Exiting...\n")
				d.Disconnect()
				return nil, 0
			} else {
				log.Errorf("Message didn't fill buffer of %d bytes with error: %v\n", length, err)
				return nil, 0
			}
		}
	}

	//log.Tracef("Received %v(%v) with byte length %v", dtcproto.DTCMessageType_name[mTypeId], mTypeId, length )
	if dtcproto.DTCMessageType(mTypeId) == dtcproto.DTCMessageType_ENCODING_RESPONSE {
		// binary encoding... nbd for now
	}
	return resp, mTypeId
}

func (d *DtcConnection) _ParseMessage(bMsg []byte, mTypeId int32) (proto.Message, reflect.Type) {
	var msg proto.Message
	pbtype := proto.MessageType("DTC_PB." + strcase.ToCamel(strings.ToLower(dtcproto.DTCMessageType_name[mTypeId])))
	//if pbtype != nil && err == nil {
	if pbtype != nil {
		msg = reflect.New(pbtype.Elem()).Interface().(proto.Message)
		//msg := reflect.New(pbtype).Interface().(proto.Message)
		//msg := reflect.New((pbtype.(reflect.Type)).Elem()).Interface().(proto.Message)
		proto.Unmarshal(bMsg, msg)
	}
	return msg, pbtype
}

func (d *DtcConnection) nextRequestID() int32 {
	d.requestId++
	return d.requestId
}

func _ParseHeaderBytes(r io.Reader) (uint16, int32) {
	hBuf := make([]byte, 4)
	io.ReadFull(r, hBuf)
	mLength := binary.LittleEndian.Uint16(hBuf[0:2])
	mTypeId := binary.LittleEndian.Uint16(hBuf[2:4])
	return mLength - 4, int32(mTypeId)
}

func dtc_bin_encoder(m interface{}) []byte {
	switch v := m.(type) {
	case dtcproto.EncodingRequest:
		bMsg := make([]byte, 8)
		binary.LittleEndian.PutUint32(bMsg[0:4], uint32(v.ProtocolVersion))
		binary.LittleEndian.PutUint32(bMsg[4:8], uint32(v.Encoding))
		bMsg = append([]byte(bMsg), []byte(v.ProtocolType)...)
		bMsg = append([]byte(bMsg), 0x00)
		return bMsg
	default:
		log.Warnf("Don't know how to bin encode type %T!\n", v)
	}
	return nil
}
