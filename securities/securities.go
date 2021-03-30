package securities

import (
    log "github.com/sirupsen/logrus"
    //"reflect"
    "github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/encoding/protojson"
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
    Market float64
}

func (s *Security) AddData( md MarketDataUpdate ) {
    switch( dtcproto.DTCMessageType(md.TypeId) ){
    /**
    * Market data
    **/
    case dtcproto.DTCMessageType_MARKET_DATA_REQUEST:
        log.Error("Server requests not supported")
        return
    case dtcproto.DTCMessageType_MARKET_DATA_REJECT:
        //log.Errorf("Got some market data reject: %v", md.Msg.(dtcproto.MarketDataReject))
        log.Error("Got some market data reject: FIXME")
        return
    case dtcproto.DTCMessageType_MARKET_DATA_SNAPSHOT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_SNAPSHOT_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_COMPACT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK:
        log.Trace( protojson.Format((md.Msg).(protoreflect.ProtoMessage)) )
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_COMPACT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_VOLUME:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_OPEN_INTEREST:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_NUM_TRADES:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADING_SESSION_DATE:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_REQUEST:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_REJECT:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP:
        return
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_INT:
        return
    case dtcproto.DTCMessageType_MARKET_DATA_FEED_STATUS:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_FEED_SYMBOL_STATUS:
        fallthrough
    case dtcproto.DTCMessageType_TRADING_SYMBOL_STATUS:
        fallthrough
    default:
        log.Trace( protojson.Format((md.Msg).(protoreflect.ProtoMessage)) )
    }

}
