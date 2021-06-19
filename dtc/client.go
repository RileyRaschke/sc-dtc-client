package dtc

import (
    log "github.com/sirupsen/logrus"
    "google.golang.org/protobuf/proto"
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
    Reconnect bool
}

func init() {
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
func (c *DtcClient) Terminated() (bool){
    return c.dtcConn.terminated
}
func (c *DtcClient) Terminate() {
    c.dtcConn.connArgs.Reconnect = false
    c.Disconnect()
}
func (c *DtcClient) Disconnect() {
    if c.dtcConn != nil {
        log.Info("Client disconnecting...\n")
        c.dtcConn.Disconnect()
    } else {
        log.Warn("No connection found!\n")
    }
}

func (c *DtcClient) ListAccounts() (x []proto.Message){
    return x
}

func (c *DtcClient) CurrentPositions() (x []proto.Message){
    return x
}

func (c *DtcClient) ListOpenOrders() (x []proto.Message){
    return x
}

func (c *DtcClient) ListHistoricalFills() (x []proto.Message){
    return x
}

func (c *DtcClient) ListHistoricalOrders() (x []proto.Message){
    return x
}

func (c *DtcClient) AddSymbol(symbol string) (x error){
    return x
}

func (c *DtcClient) RemoveSymbol(symbol string) (x error){
    return x
}

func (c *DtcClient) GetHistoricalData() (x []proto.Message){
    return x
}

func (c *DtcClient) ListSymobls() (x []string){
    return x
}

func (c *DtcClient) NextTick() (x []proto.Message){
    return x
}

func (c *DtcClient) NextUpdate() (x []proto.Message){
    return x
}

func (c *DtcClient) NextPositionUpdate() (x []proto.Message){
    return x
}

