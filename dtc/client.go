package dtc

import (
    "os"
    "log"
    "fmt"
    "net"
    "io"
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
    describe( logonRequest.ProtoReflect().Descriptor().FullName() )

    msg, err := proto.Marshal( &logonRequest )
    if( err != nil ){
        log.Fatalf("Failed to marshal logonRequest: %v\n", err)
        os.Exit(1)
    }

    log.Printf("Sending logon request...")
    d.conn.Write( PackMessage( msg, DTCMessageType_value["LOGON_REQUEST"] ))

    log.Printf("Unpacking logon response.")
    resp := d._GetMessage()

    logonResponse := LogonResponse{}
    if err := proto.Unmarshal(resp, &logonResponse); err != nil {
        log.Fatalln("Failed to parse logonResponse:", err)
    }
    describe( logonResponse.String() )
}

func (d *DtcConnection) SetEncoding() {
    encodingReq := EncodingRequest{
        ProtocolVersion: DTCVersion_value["CURRENT_VERSION"],
        Encoding: EncodingEnum_PROTOCOL_BUFFERS,
        ProtocolType: "DTC",
    }
    msg := dtc_bin_encoder( encodingReq )
    d.conn.Write( PackMessage(msg, DTCMessageType_value["ENCODING_REQUEST"] ))
    resp := d._GetMessage()
    describe(resp)
}

func PackMessage(msg []byte, mTypeId int32) ([]byte){
    length := 4 + len(msg)
    header := make([]byte, 4)
    binary.LittleEndian.PutUint16(header[0:2], uint16(length))
    binary.LittleEndian.PutUint16(header[2:4], uint16(mTypeId))
    message := append(header, msg...)
    return message
}

func (d *DtcConnection) _GetMessage() ([]byte) {
    r := bufio.NewReader(d.conn)

    length, mTypeId := ParseHeaderBytes(r)

    resp := make([]byte, length)
    _, err := io.ReadFull(r, resp)

    if err != nil {
        log.Printf("Message didn't fill buffer of %d bytes with error: %v\n", length, err)
        return nil
    }

    log.Printf("Received %v with byte length %v", DTCMessageType_name[int32(mTypeId)], length )
    return resp
}

func ParseHeaderBytes(r io.Reader) (uint16, uint16){
    hBuf := make([]byte, 4)
    io.ReadFull(r, hBuf)
    mLength := binary.LittleEndian.Uint16(hBuf[0:2])
    mTypeId := binary.LittleEndian.Uint16(hBuf[2:4])
    return mLength-4, mTypeId
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

