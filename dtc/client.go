package dtc

import (
    "os"
    "log"
    "fmt"
    "net"
    "time"
    "io"
    "bufio"
    "errors"
    "encoding/binary"
    "reflect"
    proto "github.com/golang/protobuf/proto"
    //"google.golang.org/protobuf/reflect/protoregistry"
    "strings"
    "github.com/iancoleman/strcase"
)

const DTC_CLIENT_HEARTBEAT_SECONDS = 5

type ConnectArgs struct {
    Host string
    Port string
    HistPort string
    Username string
    Password string
}

type DtcConnection struct {
    connArgs ConnectArgs
    connUri string
    conn  net.Conn
    clientClose chan int
    listenClose chan int
    heartbeatUpdate chan proto.Message
    Connected bool `default: false`
}

func init() {
    log.SetPrefix("[dtc] ")
}

func Connect( c ConnectArgs ) (*DtcConnection, error){
    d := &DtcConnection{}
    err := d.Connect(c)
    return d, err
}

func (d *DtcConnection) Connect( c ConnectArgs ) (error){
    log.Printf("Connecting: %s@%s:%s\n", c.Username, c.Host, c.Port )
    uri := net.JoinHostPort(c.Host, c.Port)
    conn, err := net.Dial("tcp", uri)
    if err != nil {
        log.Fatalf("Failed to connect to DTC server: %v\n", err)
        os.Exit(1)
    }

    d.conn = conn
    d.connArgs = c
    d.connUri = uri
    d.clientClose = make(chan int)
    d.listenClose = make(chan int)
    d.heartbeatUpdate = make(chan proto.Message)

    d._SetEncoding()

    err = d._Logon()
    if err != nil {
        d.Disconnect()
        return err;
    }
    d.Connected = true
    go d._Listen()
    d._KeepAlive()
    return nil
}

func (d *DtcConnection) Disconnect() {
    log.Printf("Client Disconnecting...")
    d.clientClose <- 0
    if d.Connected {
        d.Connected = false
        d._Logoff()
        time.Sleep( DTC_CLIENT_HEARTBEAT_SECONDS * time.Second )
        d.listenClose <- 0
    }
}

func (d *DtcConnection) _Listen() {
    log.Printf("Client listener started")
    for {
        select {
        case <-d.listenClose:
            log.Printf("Terminating Client listener")
            return
        default:
            //d._NextMessage()
            msg, _, mTypeStr := d._NextMessage()
            if mTypeStr == "" {
                describe(msg)
            }
            if msg != nil {
                fmt.Println(msg.String())
            }
        }
    }
}

func (d *DtcConnection) _KeepAlive() {
    var m proto.Message
    for {
        select {
        case m = <-d.heartbeatUpdate:
            log.Printf(m.String())
        case <-d.clientClose:
            d.conn.Close()
            log.Printf("Connection closed")
            return
        default:
            time.Sleep( DTC_CLIENT_HEARTBEAT_SECONDS * time.Second )
            if d.Connected {
                d._SendHeartbeat()
            } else {
                d.clientClose <- 0
            }
        }
    }
}

func (d *DtcConnection) _Logon() error {
    logonRequest := LogonRequest{
        Username: d.connArgs.Username,
        Password: d.connArgs.Password,
        Integer_1: 2,
        TradeMode: 0,
        HeartbeatIntervalInSeconds: DTC_CLIENT_HEARTBEAT_SECONDS+1,
        ClientName: "go-dtc",
        ProtocolVersion: DTCVersion_value["CURRENT_VERSION"],
    }
    //describe( logonRequest.ProtoReflect().Descriptor().FullName() )

    msg, err := proto.Marshal( &logonRequest )
    if err != nil {
        log.Fatalf("Failed to marshal LogonRequest message: %v\n", err)
        os.Exit(1)
    }

    log.Printf("Sending LOGON_REQUEST")
    d.conn.Write( PackMessage( msg, DTCMessageType_value["LOGON_REQUEST"] ))

    resp, _ := d._GetMessage()

    logonResponse := LogonResponse{}
    if err := proto.Unmarshal(resp, &logonResponse); err != nil {
        log.Fatalln("Failed to parse LogonResponse:", err)
    }
    if logonResponse.Result != LogonStatusEnum_LOGON_SUCCESS {
        log.Fatalln("Logon Failed with result %v and text %v", logonResponse.Result, logonResponse.ResultText)
        return errors.New("Logon Failure")
    }
    log.Printf("Logon response: %v", logonResponse.ResultText)
    return nil
}

func (d *DtcConnection) _Logoff() {
    logoff := Logoff{
        Reason: "Done",
        DoNotReconnect: 1,
    }
    msg, err := proto.Marshal( &logoff )
    if err != nil {
        log.Fatalf("Failed to marshal LogonRequest message: %v\n", err)
        os.Exit(1)
    }

    log.Printf("Sending LOGOFF")
    d.conn.Write( PackMessage( msg, DTCMessageType_value["LOGOFF"] ))
}

func (d *DtcConnection) _SetEncoding() {
    encodingReq := EncodingRequest{
        ProtocolVersion: DTCVersion_value["CURRENT_VERSION"],
        Encoding: EncodingEnum_PROTOCOL_BUFFERS,
        ProtocolType: "DTC",
    }
    msg := dtc_bin_encoder( encodingReq )
    log.Printf("Sending ENCODING_REQUEST")
    d.conn.Write( PackMessage(msg, DTCMessageType_value["ENCODING_REQUEST"] ))
    d._GetMessage()
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
    d.conn.Write( PackMessage(msg, DTCMessageType_value["HEARTBEAT"] ))
}

func PackMessage(msg []byte, mTypeId int32) ([]byte){
    length := 4 + len(msg)
    header := make([]byte, 4)
    binary.LittleEndian.PutUint16(header[0:2], uint16(length))
    binary.LittleEndian.PutUint16(header[2:4], uint16(mTypeId))
    message := append(header, msg...)
    return message
}

func (d *DtcConnection) _NextMessage() (proto.Message, reflect.Type, string) {
    bytes, mTypeId := d._GetMessage()
    m, t := d._ParseMessage( bytes, mTypeId )
    return m, t, DTCMessageType_name[mTypeId]
}

func (d *DtcConnection) _GetMessage() ([]byte, int32) {
    r := bufio.NewReader(d.conn)

    length, mTypeId := _ParseHeaderBytes(r)

    resp := make([]byte, length)
    _, err := io.ReadFull(r, resp)

    if err == io.EOF {
        log.Printf("Received end of file on communication channel. Exiting...\n")
        d.Disconnect();
        return nil, 0
    }
    if err != nil {
        log.Printf("Message didn't fill buffer of %d bytes with error: %v\n", length, err)
        return nil, 0
    }

    log.Printf("Received %v(%v) with byte length %v", DTCMessageType_name[mTypeId], mTypeId, length )
    if DTCMessageType_name[mTypeId] == "ENCODING_RESPONSE" {
        // binary encoding... nbd for now
    }
    return resp, mTypeId
}

func (d *DtcConnection) _ParseMessage(bMsg []byte, mTypeId int32) (proto.Message, reflect.Type) {
    var msg proto.Message
    pbtype := proto.MessageType( "DTC_PB." + strcase.ToCamel( strings.ToLower( DTCMessageType_name[mTypeId] ) ) )
    //if pbtype != nil && err == nil {
    if pbtype != nil {
        msg = reflect.New(pbtype.Elem()).Interface().(proto.Message)
        //msg := reflect.New(pbtype).Interface().(proto.Message)
        //msg := reflect.New((pbtype.(reflect.Type)).Elem()).Interface().(proto.Message)
        proto.Unmarshal(bMsg, msg)
    }
    return msg, pbtype
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
        case EncodingRequest:
            bMsg := make([]byte, 8)
            binary.LittleEndian.PutUint32(bMsg[0:4], uint32(v.ProtocolVersion))
            binary.LittleEndian.PutUint32(bMsg[4:8], uint32(v.Encoding))
            bMsg = append([]byte(bMsg),[]byte(v.ProtocolType)... )
            bMsg = append([]byte(bMsg), 0x00)
            return bMsg
        default:
            log.Printf("Don't know how to bin encode type %T!\n", v)
    }
    return nil
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

