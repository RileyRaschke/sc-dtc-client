package dtc

import (
    log "github.com/sirupsen/logrus"
    "fmt"
    "errors"
    "reflect"
    "github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/encoding/protojson"
)

func (d *DtcConnection) _RouteMessage(msg proto.Message, rtype reflect.Type, mTypeId int32) error {
    if msg == nil {
        mTypeStr := DTCMessageType_name[mTypeId]
        log.Errorf("Received %v with empty body", mTypeStr)
        return errors.New("Router received null message.")
    }
    switch( DTCMessageType(mTypeId) ){
    case DTCMessageType_MESSAGE_TYPE_UNSET:
        log.Trace("Received MESSAGE_TYPE_UNSET")
        return nil
    // Authentication and connection monitoring
    case DTCMessageType_LOGON_REQUEST:
        return nil // server action
    case DTCMessageType_LOGON_RESPONSE:
        return nil // handled at logon
    case DTCMessageType_HEARTBEAT:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        d.heartbeatUpdate <-msg.(*Heartbeat)
        return nil
    // Account list
    case DTCMessageType_TRADE_ACCOUNT_RESPONSE:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_TRADE_ACCOUNTS_REQUEST:
        return nil // server action
    // Account balance
    case DTCMessageType_ACCOUNT_BALANCE_UPDATE:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        fmt.Println( protojson.Format(msg.(protoreflect.ProtoMessage)) )
        return nil
    case DTCMessageType_ACCOUNT_BALANCE_REQUEST:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_ACCOUNT_BALANCE_REJECT:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_REJECT:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REQUEST:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REJECT:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_HISTORICAL_ACCOUNT_BALANCE_RESPONSE:
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        return nil
    case DTCMessageType_LOGOFF:
        log.Warn("Received logoff request from server!")
        d.loggedOn = false
        if ! d.connArgs.Reconnect {
            d.Disconnect()
        }
        return nil // server action
    case DTCMessageType_ENCODING_REQUEST:
        return nil // server action
    case DTCMessageType_ENCODING_RESPONSE:
        return nil // handled upon request
    case DTCMessageType_SECURITY_DEFINITION_RESPONSE:
        //symbolAdd <-msg.(SecurityDefinitionResponse)
        d.addSecurity( msg.(*SecurityDefinition) )
        return nil
    // Market data
    case DTCMessageType_MARKET_DATA_REQUEST:
        fallthrough
    case DTCMessageType_MARKET_DATA_REJECT:
        fallthrough
    case DTCMessageType_MARKET_DATA_SNAPSHOT:
        fallthrough
    case DTCMessageType_MARKET_DATA_SNAPSHOT_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_TRADE:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_TRADE_COMPACT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_TRADE_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_BID_ASK:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_COMPACT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_VOLUME:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_OPEN_INTEREST:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_SESSION_NUM_TRADES:
        fallthrough
    case DTCMessageType_MARKET_DATA_UPDATE_TRADING_SESSION_DATE:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_REQUEST:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_REJECT:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_INT:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP:
        fallthrough
    case DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_INT:
        fallthrough
    case DTCMessageType_MARKET_DATA_FEED_STATUS:
        fallthrough
    case DTCMessageType_MARKET_DATA_FEED_SYMBOL_STATUS:
        //fallthrough
        return nil
    case DTCMessageType_TRADING_SYMBOL_STATUS:
        fallthrough
    // Order entry and modification
    case DTCMessageType_SUBMIT_NEW_SINGLE_ORDER:
        fallthrough
    case DTCMessageType_SUBMIT_NEW_SINGLE_ORDER_INT:
        fallthrough
    case DTCMessageType_SUBMIT_NEW_OCO_ORDER:
        fallthrough
    case DTCMessageType_SUBMIT_NEW_OCO_ORDER_INT:
        fallthrough
    case DTCMessageType_SUBMIT_FLATTEN_POSITION_ORDER:
        fallthrough
    case DTCMessageType_CANCEL_ORDER:
        fallthrough
    case DTCMessageType_CANCEL_REPLACE_ORDER:
        fallthrough
    case DTCMessageType_CANCEL_REPLACE_ORDER_INT:
        fallthrough
    // Trading related
    case DTCMessageType_OPEN_ORDERS_REQUEST:
        fallthrough
    case DTCMessageType_OPEN_ORDERS_REJECT:
        fallthrough
    case DTCMessageType_ORDER_UPDATE:
        fallthrough
    case DTCMessageType_HISTORICAL_ORDER_FILLS_REQUEST:
        fallthrough
    case DTCMessageType_HISTORICAL_ORDER_FILL_RESPONSE:
        fallthrough
    case DTCMessageType_HISTORICAL_ORDER_FILLS_REJECT:
        fallthrough
    case DTCMessageType_CURRENT_POSITIONS_REQUEST:
        fallthrough
    case DTCMessageType_CURRENT_POSITIONS_REJECT:
        fallthrough
    case DTCMessageType_POSITION_UPDATE:
        fallthrough
    // Symbol discovery and security definitions
    case DTCMessageType_EXCHANGE_LIST_REQUEST:
        fallthrough
    case DTCMessageType_EXCHANGE_LIST_RESPONSE:
        fallthrough
    case DTCMessageType_SYMBOLS_FOR_EXCHANGE_REQUEST:
        fallthrough
    case DTCMessageType_UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST:
        fallthrough
    case DTCMessageType_SYMBOLS_FOR_UNDERLYING_REQUEST:
        fallthrough
    case DTCMessageType_SECURITY_DEFINITION_FOR_SYMBOL_REQUEST:
        fallthrough
    case DTCMessageType_SYMBOL_SEARCH_REQUEST:
        fallthrough
    case DTCMessageType_SECURITY_DEFINITION_REJECT:
        fallthrough
    // Logging
    case DTCMessageType_USER_MESSAGE:
        fallthrough
    case DTCMessageType_GENERAL_LOG_MESSAGE:
        fallthrough
    case DTCMessageType_ALERT_MESSAGE:
        fallthrough
    case DTCMessageType_JOURNAL_ENTRY_ADD:
        fallthrough
    case DTCMessageType_JOURNAL_ENTRIES_REQUEST:
        fallthrough
    case DTCMessageType_JOURNAL_ENTRIES_REJECT:
        fallthrough
    case DTCMessageType_JOURNAL_ENTRY_RESPONSE:
        fallthrough
    // Historical price data
    case DTCMessageType_HISTORICAL_PRICE_DATA_REQUEST:
        fallthrough
    case DTCMessageType_HISTORICAL_PRICE_DATA_RESPONSE_HEADER:
        fallthrough
    case DTCMessageType_HISTORICAL_PRICE_DATA_REJECT:
        fallthrough
    case DTCMessageType_HISTORICAL_PRICE_DATA_RECORD_RESPONSE:
        fallthrough
    case DTCMessageType_HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE:
        fallthrough
    case DTCMessageType_HISTORICAL_PRICE_DATA_RECORD_RESPONSE_INT:
        fallthrough
    case DTCMessageType_HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT:
        fallthrough
    case DTCMessageType_HISTORICAL_PRICE_DATA_RESPONSE_TRAILER:
        fallthrough
    // Historical market depth data
    case DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_REQUEST:
        fallthrough
    case DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER:
        fallthrough
    case DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_REJECT:
        fallthrough
    case DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE:
        fallthrough
    default:
        mTypeStr := DTCMessageType_name[mTypeId]
        if mTypeStr == "" {
            log.Debug("No message type determined!\n")
            describe(msg)
        }
        //fmt.Println(msg.String())
        log.Tracef("Received %v(%v)", DTCMessageType_name[mTypeId], mTypeId)
        fmt.Println( protojson.Format(msg.(protoreflect.ProtoMessage)) )
    }
    return nil
}
