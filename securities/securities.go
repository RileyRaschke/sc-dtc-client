package securities

import (
    //"fmt"
    //"math"
    //"strings"
    "sync"
    //log "github.com/sirupsen/logrus"
    //"reflect"
    "github.com/golang/protobuf/proto"
    //"google.golang.org/protobuf/reflect/protoreflect"
    //"google.golang.org/protobuf/encoding/protojson"
    //"github.com/RileyR387/sc-dtc-client/marketdata"
    "github.com/RileyR387/sc-dtc-client/dtcproto"
)

type MarketDataUpdate struct {
    Msg proto.Message
    TypeId int32
}

type Security struct {
    Definition *dtcproto.SecurityDefinitionResponse
    BidDepth map[int] int
    AskDepth map[int] int
    Bid float64
    Ask float64
    Last float64
    SettlementPrice float64
    SessionSettlementDateTime uint32
    TradingSessionDate uint32
    SessionOpenPrice float64
    SessionHighPrice float64
    SessionLowPrice float64
    SessionVolume uint32
    SessionNumTrades uint32
    OpenInterest uint32
    LastData float64
    LastTradeVolume uint32
    LastSide string
    AddingDataMutex sync.Mutex
}

func FromDef(def *dtcproto.SecurityDefinitionResponse) *Security {
    res := &Security{Definition: def}
    return res
}

func (s *Security) BidPrice() float64 {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    return s.Bid
}
func (s *Security) AskPrice() float64 {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    return s.Ask
}
func (s *Security) GetSettlementPrice() float64 {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    return s.SettlementPrice
}
func (s *Security) GetLastPrice() float64 {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    return s.Last
}
func (s *Security) GetSymbol() string {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    if s.Definition == nil {
        return ""
    }
    return s.Definition.Symbol
}
func (s *Security) GetSessionHighPrice() float64 {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    return s.SessionHighPrice
}
func (s *Security) GetSessionLowPrice() float64 {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    return s.SessionLowPrice
}
func (s *Security) GetSessionVolume() uint32 {
    s.AddingDataMutex.Lock()
    defer s.AddingDataMutex.Unlock()
    return s.SessionVolume
}
