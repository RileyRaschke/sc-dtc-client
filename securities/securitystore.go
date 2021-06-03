package securities

import (
    //"fmt"
    //"math"
    //"strings"
    "sync"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    //"github.com/golang/protobuf/proto"
    //"reflect"
    //"github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/encoding/protojson"
    //"github.com/RileyR387/sc-dtc-client/marketdata"
    //"github.com/RileyR387/sc-dtc-client/dtcproto"
)

type SecurityStore struct {
    secMap map[int32]*Security
    secMapMtx sync.Mutex
    symbolToIDMap map[string]int32
    symbols []string
}

var (
    VISIBLE_SECURTIES = []string{"F.US.EPM21","F.US.YMM21","F.US.TYAU21","USDX","F.US.NQU21","F.US.CLEN21"}
)

func NewSecurityStore() *SecurityStore {
    return &SecurityStore{ make(map[int32]*Security), sync.Mutex{}, make(map[string]int32), []string{} }
}

func (ss *SecurityStore) GetSecurityBySymbol(symb string) *Security {
    ss.secMapMtx.Lock()
    defer ss.secMapMtx.Unlock()
    if secID, ok := ss.symbolToIDMap[symb]; ok {
        return ss.secMap[secID]
    }
    return nil
}

func (ss *SecurityStore) AddSecurity(sec *Security) {
    ss.secMapMtx.Lock()
    defer ss.secMapMtx.Unlock()
    symbol := sec.GetSymbol()
    if _, ok := ss.symbolToIDMap[symbol]; ok {
        // Security has been added already.. reconnect caller?
        return
    }
    ss.symbolToIDMap[symbol] = sec.Definition.RequestID
    sec.Hide()
    for _, v := range VISIBLE_SECURTIES {
        if symbol == v {
            sec.Show()
        }
    }

    ss.secMap[sec.Definition.RequestID] = sec

    ss.symbols = append(ss.symbols, symbol)
}

func (ss *SecurityStore) AddData(secData MarketDataUpdate) {
    ss.secMapMtx.Lock()
    defer ss.secMapMtx.Unlock()
    // TODO: I shouldn't need to go to string before a map right?
    var mktDataI interface{}
    var lastMsgJson = protojson.Format((secData.Msg).(protoreflect.ProtoMessage))
    err := json.Unmarshal([]byte(lastMsgJson), &mktDataI)
    if err != nil {
        log.Errorf("Failed to unmarshal json data with error: %v", err)
        return
    }
    dmm := mktDataI.(map[string]interface{})
    var secId = int32( dmm["SymbolID"].(float64) )
    if sec, ok := ss.secMap[secId]; ok {
        sec.AddData(secData)
    }
}

func (ss *SecurityStore) GetSymbols() []string {
    ss.secMapMtx.Lock()
    defer ss.secMapMtx.Unlock()
    return ss.symbols
}

