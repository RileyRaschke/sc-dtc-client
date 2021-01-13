package dtc

import (
    log "github.com/sirupsen/logrus"
    //"reflect"
    //"github.com/golang/protobuf/proto"
    //"google.golang.org/protobuf/reflect/protoreflect"
    //"google.golang.org/protobuf/encoding/protojson"

    //"github.com/RileyR387/sc-dtc-client/marketdata"
)

type Security struct {
    Definition *SecurityDefinition
    BidDepth map[int] int
    AskDepth map[int] int
    Market float64
}

func (d *DtcConnection) addSecurity(def *SecurityDefinitionResponse) {
    log.Infof("Added security %v from exchange %v as %v", def.ExchangeSymbol, def.Exchange, def.Symbol)
    d.securityMap[def.RequestID] = &Security{Definition: def}
}

func (s *Security) AddData() {

}
