package dtc

import (
    log "github.com/sirupsen/logrus"
    //"fmt"
    "errors"
    "reflect"
    "github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/encoding/protojson"
    "github.com/RileyR387/sc-dtc-client/dtcproto"
    "github.com/RileyR387/sc-dtc-client/securities"
)

func (d *DtcConnection) _RouteMessage(msg proto.Message, rtype reflect.Type, mTypeId int32) error {
    if msg == nil {
        mTypeStr := dtcproto.DTCMessageType_name[mTypeId]
        log.Errorf("Received %v with empty body", mTypeStr)
        return errors.New("Router received null message.")
    }
    switch( dtcproto.DTCMessageType(mTypeId) ){
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
        d.heartbeatUpdate <-msg.(*Heartbeat)
        return nil
    // Account list
    case dtcproto.DTCMessageType_TRADE_ACCOUNT_RESPONSE:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_TRADE_ACCOUNTS_REQUEST:
        return nil // server action
    // Account balance
    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_UPDATE:
        log.Debugf("Received %v(%v)\n%v",
            dtcproto.DTCMessageType_name[mTypeId],
            mTypeId,
            protojson.Format(msg.(protoreflect.ProtoMessage)),
        )
        return nil
    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_REQUEST:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_REJECT:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_REJECT:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REQUEST:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REJECT:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_HISTORICAL_ACCOUNT_BALANCE_RESPONSE:
        log.Tracef("Received %v(%v)", dtcproto.DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case dtcproto.DTCMessageType_LOGOFF:
        log.Warn("Received logoff request from server!")
        d.loggedOn = false
        if ! d.connArgs.Reconnect {
            d.Disconnect()
        }
        return nil // server action
    case dtcproto.DTCMessageType_ENCODING_REQUEST:
        return nil // server action
    case dtcproto.DTCMessageType_ENCODING_RESPONSE:
        return nil // handled upon request
    case dtcproto.DTCMessageType_SECURITY_DEFINITION_RESPONSE:
        d.addSecurity( msg.(*SecurityDefinition) )
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
        //d.securityMap[ xx ].AddData( msg, mTypeId )
        //d.marketData <- &msg
        d.marketData <- securities.MarketDataUpdate{ msg, mTypeId }
        return nil
    /**
    * Order entry and modification
    */
    case dtcproto.DTCMessageType_SUBMIT_NEW_SINGLE_ORDER:
        fallthrough
    case dtcproto.DTCMessageType_SUBMIT_NEW_SINGLE_ORDER_INT:
        fallthrough
    case dtcproto.DTCMessageType_SUBMIT_NEW_OCO_ORDER:
        fallthrough
    case dtcproto.DTCMessageType_SUBMIT_NEW_OCO_ORDER_INT:
        fallthrough
    case dtcproto.DTCMessageType_SUBMIT_FLATTEN_POSITION_ORDER:
        fallthrough
    case dtcproto.DTCMessageType_CANCEL_ORDER:
        fallthrough
    case dtcproto.DTCMessageType_CANCEL_REPLACE_ORDER:
        fallthrough
    case dtcproto.DTCMessageType_CANCEL_REPLACE_ORDER_INT:
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
    case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_RECORD_RESPONSE_INT:
        fallthrough
    case dtcproto.DTCMessageType_HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT:
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
            log.Debugf("Received %v(%v)\n%v",
                dtcproto.DTCMessageType_name[mTypeId],
                mTypeId,
                protojson.Format(msg.(protoreflect.ProtoMessage)),
            )
        }
    }
    return nil
}
