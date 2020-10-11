package dtc

import (
    log "github.com/sirupsen/logrus"
)

type DtcClient struct {
    dtcConn *DtcConnection
}

type ConnectArgs struct {
    Host string
    Port string
    HistPort string
    Username string
    Password string
}

func init() {
    //log.SetPrefix("[DtcClient] ")
}

func Connect( c ConnectArgs ) (*DtcClient, error){
    dc := &DtcClient{
        dtcConn: &DtcConnection{},
    }
    err := dc.dtcConn.Connect(c)
    return dc, err
}

func (dc *DtcClient) Connect( c ConnectArgs ) (error){
    dc.dtcConn = &DtcConnection{}
    err := dc.dtcConn.Connect(c)
    return err
}

func (c *DtcClient) Connected() (bool){
    if c.dtcConn != nil {
        return c.dtcConn.connected
    }
    return false
}
func (c *DtcClient) Disconnect() {
    if c.dtcConn != nil {
        log.Printf("Client disconnecting...\n")
        c.dtcConn.Disconnect()
    } else {
        log.Printf("No connection found!\n")

    }
}

