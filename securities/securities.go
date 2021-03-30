package securities

import (
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
    Market float64
}

func (s *Security) AddData(message proto.Message, mTypeId int32 ) {
    switch( dtcproto.DTCMessageType(mTypeId) ){
    /**
    * Market data
    **/
    case dtcproto.DTCMessageType_MARKET_DATA_REQUEST:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_REJECT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_SNAPSHOT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_SNAPSHOT_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_COMPACT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_COMPACT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_VOLUME:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_OPEN_INTEREST:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_NUM_TRADES:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADING_SESSION_DATE:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_REQUEST:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_REJECT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_INT:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_FEED_STATUS:
        fallthrough
    case dtcproto.DTCMessageType_MARKET_DATA_FEED_SYMBOL_STATUS:
        fallthrough
    case dtcproto.DTCMessageType_TRADING_SYMBOL_STATUS:
        fallthrough
    default:

    }

}
