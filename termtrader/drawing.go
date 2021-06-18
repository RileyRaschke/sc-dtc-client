package termtrader

import (
    "fmt"
    "time"
    "sort"
    "strconv"
    "github.com/gookit/color"
    tm "github.com/buger/goterm"
    log "github.com/sirupsen/logrus"
    //"github.com/RileyR387/sc-dtc-client/securities"
    //"github.com/RileyR387/sc-dtc-client/accounts"
)

func blankRow() string {
    fstr := "%" + fmt.Sprintf("%d",tm.Width()) + "v"
    return fmt.Sprintf(fstr, " ")
}

func rPad(x string) string {
    return fmt.Sprintf("%s%*s", x, tm.Width()-len(x)," ")
}

func lPad(x string) string {
    return fmt.Sprintf("%*s%s", tm.Width()-len(x), " ", x)
}

func (x *TermTraderPlugin) draw() {
    x.dStart = time.Now()
    rowData := []string{}
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
    positions := x.accountStore.GetPositions()
    if len(positions) > 0 {
        rowData = append(rowData, blankRow())
        rowData = append(rowData, color.Bold.Render(" Positions "))
        fmtStr := " %-15v %10v %10v %10v"
        rowData = append(rowData, fmt.Sprintf(fmtStr, "Symbol","Quantity","AvgPrice","Age"))
        for _, pos := range positions {
            rowData = append(rowData,
                fmt.Sprintf(fmtStr,
                    pos.Symbol,
                    ColorizeChangeString(fmt.Sprintf("%v",pos.Quantity)),
                    pos.AveragePrice,
                    time.Now().Sub(time.Unix(int64(pos.EntryDateTime),0)),
                ),
            )
        }
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
        if sec.IsHidden() {
            continue
        }
        log.Infof("Printing: %v",sec.GetSymbol())
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
func (x *TermTraderPlugin) drawFooter() {
    tm.MoveCursor(1,tm.Height()-6)
    tm.Println( blankRow() )

    tm.Println( fmt.Sprintf("%v", rPad(x.inputBuffer)) )

    tm.Println( blankRow() )
    duration := time.Since(x.dStart)

    third := tm.Width()/3-1
    thirdS := strconv.Itoa(third)
    fmtThree := "%-" + thirdS + "v %-" + thirdS + "v %" + thirdS + "v"
    middle := fmt.Sprintf("Runtime: %v seconds", time.Now().Unix()-x.startTime)
    tm.Println(
        fmt.Sprintf( fmtThree,
            time.Now().Format(time.RFC1123),
            fmt.Sprintf("%*s", -third, fmt.Sprintf("%*s", (third+len(middle))/2, middle)),
            fmt.Sprintf("Draw Time: %.1fms", float64(int64(duration))/1000/1000.0),
        ),
    )
    tm.Println( blankRow() )
}

func (x *TermTraderPlugin) screenWrite(screenData *[]string) {
    tm.MoveCursor(1, 1)
    for _, row := range *screenData {
        tm.Println( row )
    }
    x.drawFooter()
    tm.Flush() // Call it every time at the end of rendering
}
