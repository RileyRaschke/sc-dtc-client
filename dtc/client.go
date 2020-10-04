package dtc

//// #cgo CPPFLAGS: -std=c++11
//// #cgo LDFLAGS: -l cstdint
//// #include <cstdint>
//// #include "DTCProtocol.cpp"
//import "C"
import (
    "os"
    "log"
    "fmt"
    "net"
    "bufio"
)

type ConnectArgs struct {
    Host string
    Port string
    HistPort string
    Username string
    Password string
}

var (
    conn net.Conn
)

func init() {
    log.SetPrefix("[dtc] ")
}

//func Connect( c ConnectArgs ) (dtcConn *dtc.Conn, err error) {
func Connect( c ConnectArgs ) {
    fmt.Printf("%s@%s:%d\n", c.Username, c.Host, c.Port )
    uri := net.JoinHostPort(c.Host, c.Port)
    conn, err := net.Dial("tcp", uri)
    if err != nil {
        log.Fatalf("Failed to connect to DTC server: %v\n", err)
        os.Exit(1)
    }
    logonRequest := LogonRequest{
        Username: c.Username,
        Password: c.Password,
        Integer_1: 2,
        HeartbeatIntervalInSeconds: 6,
        ClientName: "sc-dtc-client-go",
    }
    fmt.Fprintf(conn, logonRequest.String() )

    status, err := bufio.NewReader(conn).ReadString('\n')

    fmt.Println(status)
}

