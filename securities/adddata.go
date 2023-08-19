package securities

import (
	//"fmt"
	//"math"
	//"strings"
	//"sync"
	log "github.com/sirupsen/logrus"
	//"reflect"
	//"google.golang.org/protobuf/proto"
	"github.com/RileyR387/sc-dtc-client/dtcproto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (s *Security) AddData(md MarketDataUpdate) {
	s.AddingDataMutex.Lock()
	defer s.AddingDataMutex.Unlock()
	/**
	 * Market data
	 **/
	switch dtcproto.DTCMessageType(md.TypeId) {
	case dtcproto.DTCMessageType_MARKET_DATA_REQUEST:
		log.Error("Server requests not supported")
		return
	case dtcproto.DTCMessageType_MARKET_DATA_REJECT:
		//log.Errorf("Got some market data reject: %v", md.Msg.(dtcproto.MarketDataReject))
		log.Error("Got some market data reject: FIXME")
		return
	case dtcproto.DTCMessageType_MARKET_DATA_SNAPSHOT:
		//log.Trace( protojson.Format((md.Msg).(protoreflect.ProtoMessage)) )
		mds := md.Msg.(*dtcproto.MarketDataSnapshot)

		s.SettlementPrice = mds.SessionSettlementPrice
		s.SessionSettlementDateTime = mds.SessionSettlementDateTime
		s.TradingSessionDate = mds.TradingSessionDate
		s.SessionOpenPrice = mds.SessionOpenPrice
		s.SessionHighPrice = mds.SessionHighPrice
		s.SessionLowPrice = mds.SessionLowPrice
		s.SessionVolume = uint32(mds.SessionVolume)
		s.SessionNumTrades = mds.SessionNumTrades
		s.OpenInterest = mds.OpenInterest
		s.Bid = mds.BidPrice
		s.Ask = mds.AskPrice
		s.Last = mds.LastTradePrice
		s.LastTradeVolume = uint32(mds.LastTradeVolume)
		s.SessionVolume += uint32(mds.LastTradeVolume)

		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE:
		//log.Trace( protojson.Format((md.Msg).(protoreflect.ProtoMessage)) )
		mdu := md.Msg.(*dtcproto.MarketDataUpdateTrade)
		s.Last = float64(mdu.Price)
		s.LastTradeVolume = uint32(mdu.Volume)
		s.SessionVolume += uint32(mdu.Volume)
		s.LastData = float64(mdu.DateTime)
		s.LastSide = string(mdu.AtBidOrAsk)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_COMPACT:
		//log.Trace( protojson.Format((md.Msg).(protoreflect.ProtoMessage)) )
		mdu := md.Msg.(*dtcproto.MarketDataUpdateTradeCompact)
		s.Last = float64(mdu.Price)
		s.LastTradeVolume = uint32(mdu.Volume)
		s.SessionVolume += uint32(mdu.Volume)
		s.LastData = float64(mdu.DateTime)
		s.LastSide = string(mdu.AtBidOrAsk)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateLastTradeSnapshot)
		s.Last = mdu.LastTradePrice
		s.LastTradeVolume = uint32(mdu.LastTradeVolume)
		s.SessionVolume += uint32(mdu.LastTradeVolume)
		s.LastData = mdu.LastTradeDateTime
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR:
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2:
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP:
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateBidAsk)
		s.Bid = float64(mdu.BidPrice)
		s.Ask = float64(mdu.AskPrice)
		s.LastData = float64(mdu.DateTime)
		// TODO: (`BidQuantity`, `AskQuantity`)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_COMPACT:
		//log.Trace( protojson.Format((md.Msg).(protoreflect.ProtoMessage)) )
		mdu := md.Msg.(*dtcproto.MarketDataUpdateBidAskCompact)
		s.Bid = float64(mdu.BidPrice)
		s.Ask = float64(mdu.AskPrice)
		s.LastData = float64(mdu.DateTime)
		// TODO: (`BidQuantity`, `AskQuantity`)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateSessionHigh)
		s.SessionHighPrice = mdu.Price
		s.LastData = float64(mdu.TradingSessionDate)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateSessionLow)
		s.SessionLowPrice = mdu.Price
		s.LastData = float64(mdu.TradingSessionDate)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_VOLUME:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateSessionVolume)
		s.SessionVolume = uint32(mdu.Volume)
		s.TradingSessionDate = mdu.TradingSessionDate
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_OPEN_INTEREST:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateOpenInterest)
		s.OpenInterest = uint32(mdu.OpenInterest)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateSessionSettlement)
		s.SettlementPrice = float64(mdu.Price)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_NUM_TRADES:
		mdu := md.Msg.(*dtcproto.MarketDataUpdateSessionNumTrades)
		s.SessionNumTrades = uint32(mdu.NumTrades)
		return
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADING_SESSION_DATE:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DEPTH_REQUEST:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DEPTH_REJECT:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS:
		//log.Trace( protojson.Format((md.Msg).(protoreflect.ProtoMessage)) )
		return
	case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
		return
	case dtcproto.DTCMessageType_MARKET_DATA_FEED_STATUS:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_FEED_SYMBOL_STATUS:
		fallthrough
	case dtcproto.DTCMessageType_TRADING_SYMBOL_STATUS:
		fallthrough
	default:
		log.Trace(protojson.Format((md.Msg).(protoreflect.ProtoMessage)))
	}
}
