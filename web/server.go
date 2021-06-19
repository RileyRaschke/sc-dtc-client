package web

import (
    "fmt"
    "net/http"
    "time"
    "math"
    "sync"
    "strings"
    "strconv"
    log "github.com/sirupsen/logrus"
    "github.com/RileyR387/sc-dtc-client/securities"
    "github.com/RileyR387/sc-dtc-client/accounts"
)

type WebServerPlugin struct {
    stop chan int
    securityStore *securities.SecurityStore
    accountStore *accounts.AccountStore
    startTime int64
    refreshMicroseconds int
    dStart time.Time
    inputBuffer string
    userInputChan chan rune
    textCacheMtx sync.Mutex
    textView string
}

const REFRESH_RATE_HZ float64 = 4.0

func New(ss *securities.SecurityStore, as *accounts.AccountStore) *WebServerPlugin {
    microsecondsFloat := (1.0/REFRESH_RATE_HZ)*1000*1000
    x := &WebServerPlugin{
        make(chan int),
        ss,
        as,
        time.Now().Unix(),
        int(math.Ceil(microsecondsFloat)),
        time.Now(),
        "",
        make(chan rune),
        sync.Mutex{},
        "",
    }
    go x.Run()
    return x
}

func (x *WebServerPlugin) ReceiveData( mdu securities.MarketDataUpdate) {
}

type View struct {
    Text string
}

func (x *WebServerPlugin) Run() {
    fmt.Println("Running WebServerPlugin")
    log.Info(fmt.Sprintf("Running WebServerPlugin"))
    http.HandleFunc("/", x.Snapshot)
    http.HandleFunc("/symbols", x.Symbols)
    go http.ListenAndServe(":8081",nil)

    //var mktData securities.MarketDataUpdate
    frameCnt := 0
    for {
        select {
        case <-x.stop:
            break
        //case stdin, _ := <-x.userInputChan:
        //    x.inputRouter(stdin)
        default:
            if frameCnt >= int(REFRESH_RATE_HZ) {
                frameCnt = 0
            }
            time.Sleep( (time.Duration(x.refreshMicroseconds) * time.Microsecond) - time.Since(x.dStart))
            x.cache()
            frameCnt++
        }
    }
}
func (x *WebServerPlugin) Stop() {
    x.stop <- 1
}

func (x *WebServerPlugin) cache() {
    x.dStart = time.Now()
    rowData := []string{}
    rowData = append(rowData, "")
    rowData = append(rowData, (*x.drawWatchlist())...)
    rowData = append(rowData, (*x.drawAccountInfo())...)
    rowData = append(rowData, "")
    rowData = append(rowData, (*x.drawFooter())...)
    x.textCacheMtx.Lock()
    defer x.textCacheMtx.Unlock()
    x.textView = strings.Join(rowData,"\n")
}

func rPad(x string) string {
    return fmt.Sprintf("%s%*s", x, 120-len(x)," ")
}

func lPad(x string) string {
    return fmt.Sprintf("%*s%s", 120-len(x), " ", x)
}

func (x *WebServerPlugin) drawAccountInfo() *[]string {
    rowData := []string{}
    rowData = append(rowData, "")
    if x.accountStore.GetCashBalance() > 1 {
        rowData = append(rowData, fmt.Sprintf("         As of: %v", time.Unix(x.accountStore.LastUpdated(),0).Format(time.RFC1123) ) )
        rowData = append(rowData, fmt.Sprintf("  Cash Balance: %.2f", x.accountStore.GetCashBalance() ) )
        rowData = append(rowData, fmt.Sprintf(" Net Liquidity: %.2f", x.accountStore.GetNetBalance() ) )
        rowData = append(rowData, fmt.Sprintf("    Margin Req: %.2f", x.accountStore.GetMarginReq() ) )
    }
    positions := x.accountStore.GetPositions()
    if len(positions) > 0 {
        rowData = append(rowData, "")
        rowData = append(rowData, " Positions ")
        fmtStr := " %-15v %10v %10v %10v"
        rowData = append(rowData, fmt.Sprintf(fmtStr, "Symbol","Quantity","AvgPrice","Age"))
        for _, pos := range positions {
            rowData = append(rowData,
                fmt.Sprintf(fmtStr,
                    pos.Symbol,
                    fmt.Sprintf("%v",pos.Quantity),
                    pos.AveragePrice,
                    time.Now().Sub(time.Unix(int64(pos.EntryDateTime),0)),
                ),
            )
        }
    }
    return &rowData
}

func (x *WebServerPlugin) drawWatchlist() *[]string {
    syms := x.securityStore.GetSymbols()
    //rowData := make([]string, len(syms)+1)
    rowData := []string{}

    fmtStr := " %-13v %10v %10v %10v %10v %10v %10v %10v %10v %10v"
    rowData = append(rowData,
        //fmt.Sprintf(" %-15v %10v %10v %10v %9v %9v %10v %10v %10v %10v",
        fmt.Sprintf(fmtStr,
            "Symbol", "Bid", "Ask", "Last", "dChg", "dChg%", "Settle","High","Low","Volume",//"OI",
        ),
    )
    for _, symbol := range syms {
        sec := x.securityStore.GetSecurityBySymbol(symbol)
        //if sec.IsHidden() {
        //    continue
        //}
        rowData = append(rowData, fmt.Sprintf(fmtStr,
                sec.GetSymbol(),
                sec.BidString(),
                sec.AskString(),
                sec.LastString(),
                sec.DchgString(),
                fmt.Sprintf("%.2f%%", ((sec.GetLastPrice()-sec.GetSettlementPrice())/sec.GetSettlementPrice())*100),
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
func (x *WebServerPlugin) drawFooter() *[]string {
    rowData := []string{}
    //tm.MoveCursor(1,tm.Height()-6)
    rowData = append(rowData, "")
    duration := time.Since(x.dStart)
    third := 120/3-1
    thirdS := strconv.Itoa(third)
    fmtThree := "%-" + thirdS + "v %-" + thirdS + "v %" + thirdS + "v"
    middle := fmt.Sprintf("Runtime: %v seconds", time.Now().Unix()-x.startTime)
    rowData = append(rowData,
        fmt.Sprintf( fmtThree,
            time.Now().Format(time.RFC1123),
            fmt.Sprintf("%*s", -third, fmt.Sprintf("%*s", (third+len(middle))/2, middle)),
            fmt.Sprintf("Draw Time: %.1fms", float64(int64(duration))/1000/1000.0),
        ),
    )
    return &rowData
}

