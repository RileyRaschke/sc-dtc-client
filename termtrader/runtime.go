package termtrader

import (
    "fmt"
    "sort"
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
                //symbolDesc := (*x.securityMap)[symID].Definition.Symbol
                //log.Tracef("Update for: %v", symbolDesc)
                (*x.securityMap)[symID].AddData(mktData)
                x.Draw()
            }
        }
    }
}

func (x *TermTraderPlugin) Draw() {
    nameMap := map[string]int32{}
    syms := make([]string, 0, len(*x.securityMap))
    for k, v := range *x.securityMap {
        nameMap[v.Definition.Symbol] = k
        syms = append( syms, v.Definition.Symbol )
    }
    sort.Strings(syms)

    // By moving cursor to top-left position we ensure that console output
    // will be overwritten each time, instead of adding new.
    tm.Clear() // Clear current screen
    //for {
    tm.MoveCursor(1, 1)

    tm.Println("Current Time:", time.Now().Format(time.RFC1123))
    tm.Println("")

    fmtStr := " %-15v %10v %10v"
    tm.Println( fmt.Sprintf(fmtStr, "Symbol", "Bid", "Ask") )
    for _, symKey := range syms {
        sec := (*x.securityMap)[nameMap[symKey]]
        tm.Println( fmt.Sprintf(fmtStr, sec.Definition.Symbol, sec.Bid, sec.Ask) )
        //tm.Println( fmt.Sprintf( "%v", .String() ) )
    }


    tm.Flush() // Call it every time at the end of rendering

    //}
}

