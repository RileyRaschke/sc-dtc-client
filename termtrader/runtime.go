package termtrader

import (
    "fmt"
    "math"
    "strings"
    "github.com/gookit/color"
    "sort"
    //"strconv"
    //"encoding/json"
    //"sync"
    log "github.com/sirupsen/logrus"
    tm "github.com/buger/goterm"
    "time"
    //"github.com/golang/protobuf/proto"
    //"google.golang.org/protobuf/encoding/protojson"
    //"google.golang.org/protobuf/reflect/protoreflect"
    "github.com/RileyR387/sc-dtc-client/securities"
    "github.com/RileyR387/sc-dtc-client/accounts"
)

const REFRESH_RATE_HZ float64 = 59.98

type TermTraderPlugin struct {
    ReceiveData chan securities.MarketDataUpdate
    //lastMsgJson string
    securityStore *securities.SecurityStore
    accountStore *accounts.AccountStore
    startTime int64
    refreshMicroseconds int
}

func New(ss *securities.SecurityStore, as *accounts.AccountStore) *TermTraderPlugin {
    microsecondsFloat := (1.0/REFRESH_RATE_HZ)*1000*1000
    x := &TermTraderPlugin{
        make(chan securities.MarketDataUpdate),
        ss,
        as,
        time.Now().Unix(),
        int(math.Ceil(microsecondsFloat)),
    }
    go x.Run()
    return x
}

func (x *TermTraderPlugin) Run() {
    fmt.Println("Running TermTraderPlugin")
    log.Info(fmt.Sprintf("Running TermTraderPlugin"))
    //var mktData securities.MarketDataUpdate

    tm.Clear() // Clear current screen
    for {
        select {
        //case mktData = <-x.ReceiveData:
        case <-x.ReceiveData:
            //x.draw()
            continue
        default:
            time.Sleep( time.Duration(x.refreshMicroseconds) * time.Microsecond)
            x.draw()
        }
    }
}

func (x *TermTraderPlugin) draw() {
    rowData := []string{}
    rowData = append(rowData,
        fmt.Sprintf("Current Time: %v\tRuntime: %v",
            time.Now().Format(time.RFC1123),
            time.Now().Unix()-x.startTime,
        ),
    )
    rowData = append(rowData, blankRow())
    rowData = append(rowData, (*x.drawWatchlist())...)
    rowData = append(rowData, (*x.drawAccountInfo())...)
    rowData = append(rowData, blankRow())
    x.screenWrite(&rowData)
}

func (x *TermTraderPlugin) drawAccountInfo() *[]string {
    rowData := []string{}
    rowData = append(rowData, blankRow())
    if x.accountStore.GetCashBalance() > 1 {
        rowData = append(rowData, fmt.Sprintf("         As of: %v", time.Unix(x.accountStore.LastUpdated(),0).Format(time.RFC1123) ) )
        rowData = append(rowData, fmt.Sprintf("  Cash Balance: %.2f", x.accountStore.GetCashBalance() ) )
        rowData = append(rowData, fmt.Sprintf(" Net Liquidity: %.2f", x.accountStore.GetNetBalance() ) )
        rowData = append(rowData, fmt.Sprintf("    Margin Req: %.2f", x.accountStore.GetMarginReq() ) )
    }
    return &rowData
}

func (x *TermTraderPlugin) drawWatchlist() *[]string {
    syms := x.securityStore.GetSymbols()
    sort.Strings(syms)
    //rowData := make([]string, len(syms)+1)
    rowData := []string{}

    rowData = append(rowData,
        fmt.Sprintf(" %-15v %10v %10v %10v %9v %9v %10v %10v %10v %10v",
            "Symbol", "Bid", "Ask", "Last", "dChg", "dChg%", "Settle","High","Low","Volume",//"OI",
        ),
    )
    fmtStrColor := " %-24v %10v %10v %18v %18v %18v %10v %10v %10v %10v"
    for _, symbol := range syms {
        sec := x.securityStore.GetSecurityBySymbol(symbol)
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
                sec.FormatPrice(sec.GetSessionHighPrice()),
                sec.FormatPrice(sec.GetSessionLowPrice()),
                sec.GetSessionVolume(),
                //sec.OpenInterest,
                //time.Unix(int64(sec.SessionSettlementDateTime), 0),
            ))
    }
    return &rowData
}

func blankRow() string {
    return fmt.Sprintf("%120v", " ")
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
