package dtc

import (
	"errors"

	"github.com/RileyR387/sc-dtc-client/dtcproto"
	"github.com/RileyR387/sc-dtc-client/securities"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	//"google.golang.org/protobuf/reflect/protoreflect"
	//"google.golang.org/protobuf/reflect/protoreflect"
)

func (d *DtcConnection) _RouteMessage(msg []byte, mTypeId int32) error {
	mTypeStr, ok := dtcproto.DTCMessageType_name[mTypeId]
	if !ok {
		// TODO: What are these messages of type 155 with no body? They come in two's...
		//log.Tracef("Unknown message type id: %v", mTypeId)
		/*
			if msg != nil {
				log.Errorf("Unknown message type has value: %v", protojson.Format(msg))
			}
			return errors.New("Router received unknown message type id: " + fmt.Sprintf("%v", mTypeId))
		*/
	}
	if msg == nil {
		if mTypeStr == "" || mTypeStr == "MESSAGE_TYPE_UNSET" {
			return nil
		}
		log.Errorf("Received %v with empty body", mTypeStr)
		return errors.New("router received null message")
	}
	switch dtcproto.DTCMessageType(mTypeId) {
	case dtcproto.DTCMessageType_MESSAGE_TYPE_UNSET:
		log.Trace("Received MESSAGE_TYPE_UNSET")
		return nil
	// Authentication and connection monitoring
	case dtcproto.DTCMessageType_LOGON_REQUEST:
		return nil // server action
	case dtcproto.DTCMessageType_LOGON_RESPONSE:
		return nil // handled at logon
	case dtcproto.DTCMessageType_HEARTBEAT:
		//log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
		r := &dtcproto.Heartbeat{}
		proto.Unmarshal(msg, r)
		d.heartbeatUpdate <- r
		return nil
	/**
	 * Order entry and modification
	 */
	case dtcproto.DTCMessageType_SUBMIT_NEW_SINGLE_ORDER:
		fallthrough
	case dtcproto.DTCMessageType_SUBMIT_NEW_OCO_ORDER:
		fallthrough
	case dtcproto.DTCMessageType_SUBMIT_FLATTEN_POSITION_ORDER:
		fallthrough
	case dtcproto.DTCMessageType_CANCEL_ORDER:
		fallthrough
	case dtcproto.DTCMessageType_CANCEL_REPLACE_ORDER:
		fallthrough
	// Trading related
	case dtcproto.DTCMessageType_OPEN_ORDERS_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_OPEN_ORDERS_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_ORDER_UPDATE:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_ORDER_FILLS_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_ORDER_FILL_RESPONSE:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_ORDER_FILLS_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_CURRENT_POSITIONS_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_CURRENT_POSITIONS_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_POSITION_UPDATE:
		fallthrough
	/**
	 * Account List/Balance Data
	 **/
	case dtcproto.DTCMessageType_TRADE_ACCOUNTS_REQUEST:
		fallthrough // server action
	case dtcproto.DTCMessageType_TRADE_ACCOUNT_RESPONSE:
		fallthrough
	case dtcproto.DTCMessageType_ACCOUNT_BALANCE_UPDATE:
		fallthrough
	case dtcproto.DTCMessageType_ACCOUNT_BALANCE_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_ACCOUNT_BALANCE_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT:
		fallthrough
	case dtcproto.DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_ACCOUNT_BALANCE_RESPONSE:
		/*
			log.Debugf("Balance or Order Data Received %v(%v)\n%v",
				dtcproto.DTCMessageType_name[mTypeId],
				mTypeId,
				protojson.Format(msg),
			)
			go d.accountStore.AddData(msg, mTypeId)
		*/
		return nil

	case dtcproto.DTCMessageType_LOGOFF:
		log.Warn("Received logoff request from server!")
		d.loggedOn = false
		if !d.connArgs.Reconnect {
			d.Disconnect()
		}
		return nil // server action
	case dtcproto.DTCMessageType_ENCODING_REQUEST:
		return nil // server action
	case dtcproto.DTCMessageType_ENCODING_RESPONSE:
		return nil // handled upon request
	case dtcproto.DTCMessageType_SECURITY_DEFINITION_RESPONSE:
		s := &dtcproto.SecurityDefinitionResponse{}
		/*
			log.Debugf("Security Definition Response Received %v(%v)\n%v",
				dtcproto.DTCMessageType_name[mTypeId],
				mTypeId,
				protojson.Format(msg),
			)
		*/
		proto.Unmarshal(msg, s)
		//d.addSecurity(msg.(*SecurityDefinition))
		//d.addSecurity(*s.(*SecurityDefinition))
		d.addSecurity(s)
		return nil

	/**
	 * Market data
	 **/
	case dtcproto.DTCMessageType_MARKET_DATA_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_SNAPSHOT:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_TRADE_COMPACT:
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
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_VOLUME:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_OPEN_INTEREST:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT:
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
	case dtcproto.DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_FEED_STATUS:
		fallthrough
	case dtcproto.DTCMessageType_MARKET_DATA_FEED_SYMBOL_STATUS:
		fallthrough
	case dtcproto.DTCMessageType_TRADING_SYMBOL_STATUS:
		/** older
		//d.securityMap[ xx ].AddData( msg, mTypeId )
		//d.marketData <- &msg
		**/
		// was in use VVV
		d.marketData <- securities.MarketDataUpdate{msg, mTypeId}
		return nil
	// Symbol discovery and security definitions
	case dtcproto.DTCMessageType_EXCHANGE_LIST_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_EXCHANGE_LIST_RESPONSE:
		fallthrough
	case dtcproto.DTCMessageType_SYMBOLS_FOR_EXCHANGE_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_SYMBOLS_FOR_UNDERLYING_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_SECURITY_DEFINITION_FOR_SYMBOL_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_SYMBOL_SEARCH_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_SECURITY_DEFINITION_REJECT:
		fallthrough
	// Logging
	case dtcproto.DTCMessageType_USER_MESSAGE:
		fallthrough
	case dtcproto.DTCMessageType_GENERAL_LOG_MESSAGE:
		fallthrough
	case dtcproto.DTCMessageType_ALERT_MESSAGE:
		fallthrough
	case dtcproto.DTCMessageType_JOURNAL_ENTRY_ADD:
		fallthrough
	case dtcproto.DTCMessageType_JOURNAL_ENTRIES_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_JOURNAL_ENTRIES_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_JOURNAL_ENTRY_RESPONSE:
		fallthrough
	// Historical price data
	case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_RESPONSE_HEADER:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_RECORD_RESPONSE:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_RESPONSE_TRAILER:
		fallthrough
	// Historical market depth data
	case dtcproto.DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_REQUEST:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_REJECT:
		fallthrough
	case dtcproto.DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE:
		fallthrough
	default:
		mTypeStr := dtcproto.DTCMessageType_name[mTypeId]
		if mTypeStr == "" {
			log.Error("No message type determined!\n")
			describe(msg)
		} else {
			/*
				log.Debugf("Received %v(%v)\n%v",
					dtcproto.DTCMessageType_name[mTypeId],
					mTypeId,
					protojson.Format(msg.(protoreflect.ProtoMessage)),
				)
			*/
		}
	}
	return nil
}
