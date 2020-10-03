package dtc

//// #cgo CPPFLAGS: -std=c++11
//// #cgo LDFLAGS: -l cstdint
//// #include <cstdint>
//// #include "DTCProtocol.cpp"
//import "C"
import "log"
import "fmt"

type ConnectArgs struct {
    Host string
    Port int
    HistPort int
    Username string
    Password string
}

func init() {
    log.SetPrefix("[dtc] ")
}

//func Connect( c ConnectArgs ) (dtcConn *dtc.Conn, err error) {
func Connect( c ConnectArgs ) {
    fmt.Printf("%s@%s:%d\n", c.Username, c.Host, c.Port )
}

