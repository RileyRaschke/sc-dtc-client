package dtc

import (
    log "github.com/sirupsen/logrus"
    //"reflect"
    //"github.com/golang/protobuf/proto"
    //"google.golang.org/protobuf/reflect/protoreflect"
    //"google.golang.org/protobuf/encoding/protojson"
)

func (d *DtcConnection) addSecurity(def *SecurityDefinitionResponse) {
    log.Infof("Added security %v from exchange %v as %v", def.ExchangeSymbol, def.Exchange, def.Symbol)
    d.securityMap[def.RequestID] = def
}

