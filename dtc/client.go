package dtc

import (
    "os"
    "log"
    "fmt"
    "net"
    "bufio"
    "encoding/binary"
    proto "github.com/golang/protobuf/proto"
)

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
    fmt.Printf("%s@%s:%s\n", c.Username, c.Host, c.Port )
    uri := net.JoinHostPort(c.Host, c.Port)

    conn, err := net.Dial("tcp", uri)
    if err != nil {
        log.Fatalf("Failed to connect to DTC server: %v\n", err)
        os.Exit(1)
    }

    d.conn = conn
    d.connArgs = c
    d.connUri = uri
    d.SetEncoding()
    d.Logon()
    return err
}

func (d *DtcConnection) Logon() {
    logonRequest := LogonRequest{
        Username: d.connArgs.Username,
        Password: d.connArgs.Password,
        Integer_1: 2,
        TradeMode: 0,
        HeartbeatIntervalInSeconds: 6,
        ClientName: "go-dtc",
        ProtocolVersion: DTCVersion_value["CURRENT_VERSION"],
    }

    msg, err := proto.Marshal( &logonRequest )
    if( err != nil ){
        log.Fatalf("Failed to marshal logonRequest: %v\n", err)
        os.Exit(1)
    }

    status, err := d.SendMessage( msg, DTCMessageType_value["LOGON_REQUEST"] )

    describe(status)
}

func (d *DtcConnection) SendMessage( msg []byte, mTypeId int32) ([]byte, error) {
    length := 4 + len(msg)
    header := make([]byte, 4)
    binary.LittleEndian.PutUint16(header[0:2], uint16(length))
    binary.LittleEndian.PutUint16(header[2:4], uint16(mTypeId))
    message := append(header, msg...)
    d.conn.Write( message )

    status, err := bufio.NewReader(d.conn).ReadBytes(0x00)

    return status,err
}

func (d *DtcConnection) SetEncoding() {
    encodingReq := EncodingRequest{
        ProtocolVersion: DTCVersion_value["CURRENT_VERSION"],
        Encoding: EncodingEnum_PROTOCOL_BUFFERS,
        ProtocolType: "DTC",
    }
    msg := dtc_bin_encoder( encodingReq )
    status, err := d.SendMessage( msg, DTCMessageType_value["ENCODING_REQUEST"] )

    if( err != nil ){
        log.Fatalf("Failed to set protocol encoding: %v\n", err)
        os.Exit(1)
    }

    describe(status)
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
    return nil;
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

