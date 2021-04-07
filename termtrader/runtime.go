package termtrader

import (
    "fmt"
    "strings"
    "github.com/gookit/color"
    "sort"
    //"strconv"
    //"encoding/json"
    "sync"
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
    securityMapMutex sync.RWMutex
}

func New(sm *map[int32] *securities.Security, mtx sync.RWMutex) *TermTraderPlugin {
    x := &TermTraderPlugin{
        make(chan securities.MarketDataUpdate),
        "",
        sm,
        mtx,
    }
    go x.Run()
    return x
}

func (x *TermTraderPlugin) Run() {
    fmt.Println("Running TermTraderPlugin")
    log.Info(fmt.Sprintf("Running TermTraderPlugin"))
    var mktData securities.MarketDataUpdate
    tm.Clear() // Clear current screen
    for {
        select {
        case mktData = <-x.ReceiveData:
            x.lastMsgJson = protojson.Format((mktData.Msg).(protoreflect.ProtoMessage))
            x.securityMapMutex.Lock()
            x.DrawWatchlist()
            x.securityMapMutex.Unlock()
        }
    }
}

func (x *TermTraderPlugin) screenWrite(screenData *[]string) {
    //tm.Clear() // Clear current screen - Do Manually to prevent flashing
    // By moving cursor to top-left position we ensure that console output
    // will be overwritten each time, instead of adding new.
    tm.MoveCursor(1, 1)
    for _, row := range *screenData {
        tm.Println( row )
    }
    tm.Flush() // Call it every time at the end of rendering
}


func (x *TermTraderPlugin) DrawWatchlist() {
    nameMap := map[string]int32{}
    syms := make([]string, 0, len(*x.securityMap))
    for k, v := range *x.securityMap {
        nameMap[v.Definition.Symbol] = k
        v.AddingDataMutex.Lock()
        syms = append( syms, v.Definition.Symbol )
        v.AddingDataMutex.Unlock()
    }
    sort.Strings(syms)

    rowData := []string{}
    rowData = append(rowData, fmt.Sprintf("Current Time: %v", time.Now().Format(time.RFC1123)))
    rowData = append(rowData, "")
    rowData = append(rowData,
        fmt.Sprintf(" %-15v %10v %10v %10v %9v %9v %10v %10v %10v %10v",
            "Symbol", "Bid", "Ask", "Last", "dChg", "dChg%", "Settle","High","Low","Volume",//"OI",
        ),
    )
    fmtStrColor := " %-24v %10v %10v %18v %18v %18v %10v %10v %10v %10v"
    for _, symKey := range syms {
        sec := (*x.securityMap)[nameMap[symKey]]
        rowData = append(rowData, fmt.Sprintf(fmtStrColor,
                color.FgYellow.Render(sec.GetSymbol()),
                sec.BidString(),
                sec.AskString(),
                color.Bold.Render(sec.LastString()),
                ColorizeChangeString( sec.DchgString() ),
                ColorizeChangeString(
                    fmt.Sprintf("%.2f%%", ((sec.GetLastPrice()-sec.GetSettlementPrice())/sec.GetSettlementPrice())*100),
                ),
                sec.SettlementString(),
                sec.FormatPrice(sec.SessionHighPrice),
                sec.FormatPrice(sec.SessionLowPrice),
                sec.SessionVolume,
                //sec.OpenInterest,
                //time.Unix(int64(sec.SessionSettlementDateTime), 0),
            ))
    }
    rowData = append(rowData, "")
    //rowData = append(rowData, fmt.Sprintf("%+v\n", (*x.securityMap)[nameMap["F.US.MESM21"]].Definition))
    rowData = append(rowData, "")
    x.screenWrite(&rowData)
}

func ColorizeChangeString(v string) string {
    if strings.HasPrefix(v,"-") || strings.HasPrefix(v, "'-") {
        red := color.FgRed.Render
        return red(v)

    } else {
        green := color.FgGreen.Render
        return green(v)

    }
}

