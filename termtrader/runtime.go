package termtrader

import (
    "fmt"
    //"strconv"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    tm "github.com/buger/goterm"
    "time"
    //"github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/reflect/protoreflect"
    "github.com/RileyR387/sc-dtc-client/securities"
)

type TermTraderPlugin struct {
    ReceiveData chan securities.MarketDataUpdate
    lastMsgJson string
    securityMap *map[int32] *securities.Security
}

func New(sm *map[int32] *securities.Security) *TermTraderPlugin {
    x := &TermTraderPlugin{
        make(chan securities.MarketDataUpdate),
        "",
        sm,
    }
    go x.Run()
    return x
}

func (x *TermTraderPlugin) Run() {
    fmt.Println("Running TermTraderPlugin")
    log.Info(fmt.Sprintf("Running TermTraderPlugin"))
    var mktData securities.MarketDataUpdate
    for {
        select {
        case mktData = <-x.ReceiveData:
            var mktDataI interface{}
            // TODO: I shouldn't need to go to string before a map right?
            x.lastMsgJson = protojson.Format((mktData.Msg).(protoreflect.ProtoMessage))
            err := json.Unmarshal([]byte(x.lastMsgJson), &mktDataI)
            dmm := mktDataI.(map[string]interface{})

            if symID := int32( dmm["SymbolID"].(float64) ); err == nil {
                symbolDesc := (*x.securityMap)[symID].Definition.Symbol
                log.Tracef("Update for: %v", symbolDesc)
                (*x.securityMap)[symID].AddData(mktData)
            }
        }
    }
}

func (x *TermTraderPlugin) Draw() {
    // By moving cursor to top-left position we ensure that console output
    // will be overwritten each time, instead of adding new.
    tm.Clear() // Clear current screen
    for {
        tm.MoveCursor(1, 1)

        tm.Println("Current Time:", time.Now().Format(time.RFC1123))

        tm.Flush() // Call it every time at the end of rendering

        time.Sleep(time.Second)
    }
}

