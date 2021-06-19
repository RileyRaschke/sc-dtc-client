package securities

import (
    //"fmt"
    //"math"
    //"strings"
    "sync"
    "sort"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    //"google.golang.org/protobuf/proto"
    //"reflect"
    //"google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/encoding/protojson"
    //"google.golang.org/protobuf/types/known/structpb"
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
    VISIBLE_SECURTIES = []string{"F.US.EPU21","F.US.EPM21","F.US.YMU21","F.US.YMM21","F.US.TYAU21","USDX","F.US.NQU21","F.US.CLEN21"}
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
        log.Infof("Known symbol(%v) added to security store when already exists", symbol)
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

    sort.Strings(ss.symbols)
}

func (ss *SecurityStore) AddData(secData MarketDataUpdate) {
    ss.secMapMtx.Lock()
    defer ss.secMapMtx.Unlock()
    // TODO: I shouldn't need to go to string before a map right?
    var mktDataI interface{}
    var msgJson = protojson.Format((secData.Msg).(protoreflect.ProtoMessage))
    err := json.Unmarshal([]byte(msgJson), &mktDataI)
    if err != nil {
        log.Errorf("Failed to unmarshal json data with error: %v", err)
        return
    }
    dmm := mktDataI.(map[string]interface{})

    /*
    //v, err := structpb.NewValue((secData.Msg).(protoreflect.ProtoMessage))
    v, err := structpb.NewValue(secData.Msg)
    if err != nil {
        log.Errorf("Failed to convert secData to structpb value with error: %v", err)
        return
    }
    m := v.GetStructValue()
    dmm := m.AsMap()
    if dmm == nil {
        d, _ := m.MarshalJSON()
        log.Debugf("MarketDataUpdate marshalled to a nil interface with json: %v", string(d) )
        // Market data unavailable?
        return
    }
    if _, ok := dmm["SymbolID"]; !ok {
        d, _ := m.MarshalJSON()
        log.Debugf("SecurityStore recieved addData without SymbolID as: %v", string(d))
        return
    }
    */
    var secId = int32( dmm["SymbolID"].(float64) )
    if sec, ok := ss.secMap[secId]; ok {
        sec.AddData(secData)
    }
}

func (ss *SecurityStore) GetSymbols() []string {
    ss.secMapMtx.Lock()
    defer ss.secMapMtx.Unlock()
    if len(ss.symbols) == 0 {
        return []string{}
    }
    res := make([]string, len(ss.symbols))
    copy( res, ss.symbols )
    return res
}

