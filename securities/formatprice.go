package securities

import (
    "fmt"
    "math"
    "strings"
    log "github.com/sirupsen/logrus"
    //"github.com/RileyR387/sc-dtc-client/dtcproto"
)

func (s *Security) FormatPrice(p float64) string {
    dfmt := s.Definition.PriceDisplayFormat
    dfmtString := dfmt.String()
    dfmtVal := int(dfmt)
    var fmtStr string
    //log.Debugf("Format value is %v with string val: %v", dfmtVal, dfmtString)
    if dfmtString == "PRICE_DISPLAY_FORMAT_UNSET" || dfmtString == "PRICE_DISPLAY_FORMAT_DECIMAL_0" {
        fmtStr = "%.f"
    }else if  strings.HasPrefix(dfmtString, "PRICE_DISPLAY_FORMAT_DECIMAL") {
        fmtStr = "%." + fmt.Sprint(dfmtVal) + "f"
    }else if strings.HasPrefix(dfmtString, "PRICE_DISPLAY_FORMAT_DENOMINATOR") {
        //log.Debugf("Format value for %v is %v with string val: %v", s.Definition.Symbol, dfmtVal, dfmtString)
        intVal, fraction := math.Modf(float64(p))
        return fmt.Sprintf("%v'%05.2f", intVal, fraction*32)
    } else {
        log.Warnf("Unknown price display format: %v", dfmtString)
        fmtStr = "%.6f"
    }
    return fmt.Sprintf(fmtStr, p)
}

func (s *Security) String() string {
    return fmt.Sprintf("%15v %10v %10v %v", s.Definition.Symbol, s.BidString(), s.AskString(), s.Definition)
}
func (s *Security) BidString() string {
    return s.FormatPrice(s.Bid)
}
func (s *Security) AskString() string {
    return s.FormatPrice(s.Ask)
}
func (s *Security) LastString() string {
    return s.FormatPrice(s.Last)
}
func (s *Security) DchgString() string {
    return s.FormatPrice(s.Last-s.SettlementPrice)
}

func (s *Security) SettlementString() string {
    return s.FormatPrice(s.SettlementPrice)
}
