package dtc

import (
    "os"
    log "github.com/sirupsen/logrus"
    "fmt"
    "net"
    "time"
    "io"
    "sync"
    "bufio"
    "errors"
    "encoding/binary"
    "encoding/json"
    "reflect"
    "github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/encoding/protojson"
    //"google.golang.org/protobuf/reflect/protoregistry"
    "strings"
    "github.com/iancoleman/strcase"

    "github.com/RileyR387/sc-dtc-client/dtcproto"
    //"github.com/RileyR387/sc-dtc-client/marketdata"
    "github.com/RileyR387/sc-dtc-client/securities"
    ttr "github.com/RileyR387/sc-dtc-client/termtrader"
)

const DTC_CLIENT_HEARTBEAT_SECONDS = 10

type DtcConnection struct {
    connArgs ConnectArgs
    connUri string
    conn  net.Conn
    connected bool
    terminated bool
    reader io.Reader
    requestId int32
    clientClose chan int
    listenClose chan int
    lastHeartbeatResponse int64
    heartbeatUpdate chan *Heartbeat
    marketData chan securities.MarketDataUpdate
    subscribers []*ttr.TermTraderPlugin
    securityMap map[int32] *securities.Security
    securityMapMutex sync.RWMutex
    accountMap  map[string] *AccountBalance
    keepingAlive bool
    listening bool
    loggedOn bool
    backOffRate int
}

func init() {
}

func (d *DtcConnection) _Listen() {
    d.listenClose = make(chan int)
    d.listening = true
    log.Infof("Client listener started")
    for {
        select {
        case <-d.listenClose:
            d.listening = false
            log.Warn("Closing client listener")
            return
        default:
            d._RouteMessage( d._NextMessage() )
        }
    }
}

func (d *DtcConnection) Connect( c ConnectArgs ) (error){
    log.Infof("Connecting: %s@%s:%s\n", c.Username, c.Host, c.Port )
    uri := net.JoinHostPort(c.Host, c.Port)
    dialer := net.Dialer{Timeout: 4*time.Second}

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
    d.reader = bufio.NewReader(d.conn)
    d.connArgs = c
    d.connUri = uri

    d.securityMap = make(map[int32]*securities.Security)
    d.securityMapMutex = sync.RWMutex{}
    d.accountMap = make(map[string]*AccountBalance)

    d._SetEncoding()

    err = d._Logon()
    if err != nil {
        d.Disconnect()
        return err
    }
    d.loggedOn = true
    d.connected = true
    go d._Listen()
    go d.keepAlive()
    if d.backOffRate < 2 {
        d.backOffRate = 1
        log.Trace("Starting hearbeat listener...")
        go d._ReceiveHeartbeat()
    }
    go d.startSubscriptionRouter()
    //go d.initTrading()

    return nil
}
/**
* Test bed for now.
*/
func (d *DtcConnection) initTrading() (error) {

    time.Sleep( 2*time.Second )
    err := d.LoadAccounts()
    if err != nil {
        log.Fatalf("Failed to load accounts with error: %v", err)
    }

    time.Sleep( 2*time.Second )
    err = d.AccountBlanaceRefresh()
    if err != nil {
        log.Fatalf("Failed to load account balances with error: %v", err)
    }

    time.Sleep( 2*time.Second )
    err = d.HistoricalFills()
    if err != nil {
        log.Fatalf("Failed to load historical fills with error: %v", err)
    }

    return err
}

func (d *DtcConnection) addSecurity(def *dtcproto.SecurityDefinitionResponse) {
    log.Infof("Added security %v from exchange %v as %v with ID: %v", def.ExchangeSymbol, def.Exchange, def.Symbol, def.RequestID)
    d.securityMapMutex.Lock()
    d.securityMap[def.RequestID] = &securities.Security{Definition: def}
    d.securityMapMutex.Unlock()
}

func (d *DtcConnection) Disconnect() {
    if d.keepingAlive {
        d.clientClose <-0
    }
    if d.loggedOn {
        d._Logoff()
    }

    if d.connected {
        d.conn.Close()
        log.Info("Connection closed")
    }

    if d.listening {
        go d.closeListener()
        time.Sleep( time.Second )
    }

    d.connected = false
    if !d.connArgs.Reconnect {
        d.terminated = true
    }
    log.Info("Disconnected")
}

func (d *DtcConnection) closeListener() {
        d.listening = false
        d.listenClose <-0
}

func (d *DtcConnection) keepAlive() {
    d.clientClose = make(chan int)
    d.keepingAlive = true
    for {
        select {
        case <-d.clientClose:
            d.keepingAlive = false
            log.Debugf("Client Heartbeat closed\n")
            return
        default:
            time.Sleep( DTC_CLIENT_HEARTBEAT_SECONDS * time.Second )
            if d.connected {
                d._SendHeartbeat()
            } else {
                d.clientClose <-0
            }
        }
    }
}

func (d *DtcConnection) startSubscriptionRouter(){
    var msg securities.MarketDataUpdate
    d.marketData = make(chan securities.MarketDataUpdate)
    d.subscribers = []*ttr.TermTraderPlugin{ ttr.New(&d.securityMap, d.securityMapMutex ) }

    for {
        select {
        case msg = <-d.marketData:
            //d.lastHeartbeatResponse = time.Now().Unix()
            var mktDataI interface{}
            // TODO: I shouldn't need to go to string before a map right?
            var lastMsgJson = protojson.Format((msg.Msg).(protoreflect.ProtoMessage))
            err := json.Unmarshal([]byte(lastMsgJson), &mktDataI)
            dmm := mktDataI.(map[string]interface{})

            if symID := int32( dmm["SymbolID"].(float64) ); err == nil {
                //symbolDesc := d.securityMap)[symID].Definition.Symbol
                //log.Tracef("Update for: %v", symbolDesc)
                d.securityMapMutex.Lock()
                d.securityMap[symID].AddData(msg)
                d.securityMapMutex.Unlock()
            }
            // Distribute Market Data
            for _, subscriber := range d.subscribers {
                subscriber.ReceiveData <-msg
            }
        }
    }
}

func (d *DtcConnection) _ReceiveHeartbeat() {
    var m *Heartbeat
    d.heartbeatUpdate = make(chan *Heartbeat)
    d.lastHeartbeatResponse = time.Now().Unix()
    for {
        select {
        case m = <-d.heartbeatUpdate:
            if d.lastHeartbeatResponse < m.CurrentDateTime-30 {
                log.Warnf("Received late heartbeat, span from last: %v", m.CurrentDateTime-d.lastHeartbeatResponse)
            }
            d.lastHeartbeatResponse = m.CurrentDateTime
        default:
            //time.Sleep( DTC_CLIENT_HEARTBEAT_SECONDS * time.Millisecond )
            time.Sleep( time.Duration(d.backOffRate) * 500 * time.Millisecond )
            if time.Now().Unix() - d.lastHeartbeatResponse > DTC_CLIENT_HEARTBEAT_SECONDS*2 {
                log.Warnf("No server heartbeat received in %v seconds", time.Now().Unix()-d.lastHeartbeatResponse)
                if d.listening && !d.connArgs.Reconnect {
                    d.Disconnect()
                }
                if !d.listening && !d.connArgs.Reconnect {
                    log.Warn("No longer listening, closing hearbeat listening thread")
                    return
                } else {
                    log.Warn("Attempting reconnect...")
                    d.backOffRate = d.backOffRate * 2
                    d.Connect( d.connArgs )
                }
            }
        }
    }
}

func (d *DtcConnection) _Logon() error {
    logonRequest := dtcproto.LogonRequest{
        Username: d.connArgs.Username,
        Password: d.connArgs.Password,
        Integer_1: 2,
        //TradeMode: dtcproto.TradeModeEnum_TRADE_MODE_UNSET,
        TradeMode: dtcproto.TradeModeEnum_TRADE_MODE_LIVE,
        //TradeMode: dtcproto.TradeModeEnum_TRADE_MODE_SIMULATED,
        HeartbeatIntervalInSeconds: DTC_CLIENT_HEARTBEAT_SECONDS+1,
        ClientName: "go-dtc",
        ProtocolVersion: dtcproto.DTCVersion_value["CURRENT_VERSION"],
    }
    //describe( logonRequest.ProtoReflect().Descriptor().FullName() )
    fmt.Println( protojson.Format(&logonRequest) )

    msg, err := proto.Marshal( &logonRequest )
    if err != nil {
        log.Fatalf("Failed to marshal LogonRequest message: %v\n", err)
        os.Exit(1)
    }

    log.Debug("Sending LOGON_REQUEST")
    d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["LOGON_REQUEST"] ))

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
    d.lastHeartbeatResponse = time.Now().Unix()
    log.Debugf("Received %v result: %v", dtcproto.DTCMessageType_name[mTypeId], logonResponse.ResultText)
    fmt.Println( protojson.Format(&logonResponse) )
    return nil
}

func (d *DtcConnection) _Logoff() {
    logoff := dtcproto.Logoff{
        Reason: "Done",
        DoNotReconnect: 1,
    }
    msg, err := proto.Marshal( &logoff )
    if err != nil {
        log.Fatalf("Failed to marshal LogonRequest message: %v\n", err)
        os.Exit(1)
    }

    log.Debug("Sending LOGOFF")
    d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["LOGOFF"] ))
    log.Trace("Logoff request sent")
}

func (d *DtcConnection) _SetEncoding() {
    encodingReq := dtcproto.EncodingRequest{
        ProtocolVersion: dtcproto.DTCVersion_value["CURRENT_VERSION"],
        Encoding: dtcproto.EncodingEnum_PROTOCOL_BUFFERS,
        ProtocolType: "DTC",
    }
    msg := dtc_bin_encoder( encodingReq )
    log.Debug("Sending ENCODING_REQUEST")
    d.conn.Write( PackMessage(msg, dtcproto.DTCMessageType_value["ENCODING_REQUEST"] ))
    respBin, mTypeId := d._GetMessage()
    log.Debugf("Received %v(%v) with bytes %v", dtcproto.DTCMessageType_name[mTypeId], mTypeId, respBin )
    /**
     * TODO: Handle binary encoding response for log purposes
    resp := d._GetMessage()
    */
}

func (d *DtcConnection) _SendHeartbeat() {
    heartbeat := Heartbeat{
        NumDroppedMessages: 0,
        CurrentDateTime: time.Now().Unix(),
    }
    msg, err := proto.Marshal( &heartbeat )
    if err != nil {
        log.Fatalf("Failed to marshal Heartbeat message: %v\n", err)
        os.Exit(1)
    }
    d.conn.Write( PackMessage(msg, dtcproto.DTCMessageType_value["HEARTBEAT"] ))
}

func PackMessage(msg []byte, mTypeId int32) ([]byte){
    length := 4 + len(msg)
    header := make([]byte, 4)
    binary.LittleEndian.PutUint16(header[0:2], uint16(length))
    binary.LittleEndian.PutUint16(header[2:4], uint16(mTypeId))
    message := append(header, msg...)
    if dtcproto.DTCMessageType_name[mTypeId] != "HEARTBEAT" {
        log.Tracef("Packed message with TypeID: %v (%v) with length %v and contents: 0x%x", mTypeId, dtcproto.DTCMessageType_name[mTypeId], length, message)
    }
    return message
}

func (d *DtcConnection) _NextMessage() (proto.Message, reflect.Type, int32) {
    bytes, mTypeId := d._GetMessage()
    m, t := d._ParseMessage( bytes, mTypeId )
    return m, t, mTypeId
}

func (d *DtcConnection) _GetMessage() ([]byte, int32) {

    length, mTypeId := _ParseHeaderBytes(d.reader)

    if length == 0  {
        log.Warnf("Received %v(%v) with byte length %v", dtcproto.DTCMessageType_name[mTypeId], mTypeId, length )
    } else if log.GetLevel() == log.TraceLevel {
        //log.Tracef("Received %v(%v) with byte length %v", dtcproto.DTCMessageType_name[mTypeId], mTypeId, length )
    }

    resp := make([]byte, length)
    _, err := io.ReadFull(d.reader, resp)

    if err != nil {
        switch t := err.(type) {
        case *net.OpError:
            if t.Op == "read" {
                log.Info("Reader closed")
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
    pbtype := proto.MessageType( "DTC_PB." + strcase.ToCamel( strings.ToLower( dtcproto.DTCMessageType_name[mTypeId] ) ) )
    //if pbtype != nil && err == nil {
    if pbtype != nil {
        msg = reflect.New(pbtype.Elem()).Interface().(proto.Message)
        //msg := reflect.New(pbtype).Interface().(proto.Message)
        //msg := reflect.New((pbtype.(reflect.Type)).Elem()).Interface().(proto.Message)
        proto.Unmarshal(bMsg, msg)
    }
    return msg, pbtype
}

func (d *DtcConnection) nextRequestID() (int32){
    d.requestId++
    return d.requestId
}

func _ParseHeaderBytes(r io.Reader) (uint16, int32){
    hBuf := make([]byte, 4)
    io.ReadFull(r, hBuf)
    mLength := binary.LittleEndian.Uint16(hBuf[0:2])
    mTypeId := binary.LittleEndian.Uint16(hBuf[2:4])
    return mLength-4, int32(mTypeId)
}

func dtc_bin_encoder( m interface{} ) ([]byte) {
    switch v := m.(type) {
        case dtcproto.EncodingRequest:
            bMsg := make([]byte, 8)
            binary.LittleEndian.PutUint32(bMsg[0:4], uint32(v.ProtocolVersion))
            binary.LittleEndian.PutUint32(bMsg[4:8], uint32(v.Encoding))
            bMsg = append([]byte(bMsg),[]byte(v.ProtocolType)... )
            bMsg = append([]byte(bMsg), 0x00)
            return bMsg
        default:
            log.Warnf("Don't know how to bin encode type %T!\n", v)
    }
    return nil
}
