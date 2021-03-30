package securities

import (
    //log "github.com/sirupsen/logrus"
    //"reflect"
    //"github.com/golang/protobuf/proto"
    //"google.golang.org/protobuf/reflect/protoreflect"
    //"google.golang.org/protobuf/encoding/protojson"
    //"github.com/RileyR387/sc-dtc-client/marketdata"
    "github.com/RileyR387/sc-dtc-client/dtcproto"
)

type Security struct {
    Definition *dtcproto.SecurityDefinitionResponse
    BidDepth map[int] int
    AskDepth map[int] int
    Market float64
}

func (s *Security) AddData() {

}
