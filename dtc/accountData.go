package dtc

import (
    //"fmt"
    "errors"
    log "github.com/sirupsen/logrus"
    "google.golang.org/protobuf/proto"
    //"google.golang.org/protobuf/encoding/protojson"
    "github.com/RileyR387/sc-dtc-client/dtcproto"
)

/**
* Should trigger repsonse handled by DtcConnection._Listen() -> _RouteMessage() but seeing no bytes on wire...
*/
func (d *DtcConnection) LoadAccounts() error {
    tradeActReq := dtcproto.TradeAccountsRequest{
        RequestID: d.nextRequestID(),
    }
    //describe( tradeActReq.ProtoReflect().Descriptor().FullName() )

    msg, err := proto.Marshal( &tradeActReq )
    if err != nil {
        log.Errorf("Failed to marshal TradeAccountsRequest message: %v\n", err)
        return errors.New("Failed to marshal TradeAccountsRequest, check proto version?")
    }
    //msg = append(msg, 0x00) // not spec for protobuf, but tried it anyways...
    //describe( msg )

    log.Debug("Sending TRADE_ACCOUNTS_REQUEST")
    //_, err = d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["TRADE_ACCOUNTS_REQUEST"] ))
    d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["TRADE_ACCOUNTS_REQUEST"] ))
    //fmt.Println( protojson.Format(&tradeActReq) )
    return err
}

/**
* Should trigger repsonse handled by DtcConnection._Listen() -> _RouteMessage() but seeing no bytes on wire...
*/
func (d *DtcConnection) AccountBlanaceRefresh(acctId string) error {
    balReq := dtcproto.AccountBalanceRequest{
        RequestID: d.nextRequestID(),
        TradeAccount: acctId,
    }
    //describe( balReq.ProtoReflect().Descriptor().FullName() )

    msg, err := proto.Marshal( &balReq )
    if err != nil {
        log.Errorf("Failed to marshal ACCOUNT_BALANCE_REQUEST message: %v\n", err)
        return errors.New("Failed to marshal ACCOUNT_BALANCE_REQUEST, check proto version?")
    }
    //describe( msg )

    log.Debug("Sending ACCOUNT_BALANCE_REQUEST")
    //_, err = d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["ACCOUNT_BALANCE_REQUEST"] ))
    d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["ACCOUNT_BALANCE_REQUEST"] ))
    //fmt.Println( protojson.Format(&balReq) )
    return err
}

func (d *DtcConnection) HistoricalFills(tradeAccount string) error {
    fillHistReq := dtcproto.HistoricalOrderFillsRequest{
        RequestID: d.nextRequestID(),
        ServerOrderID: "",
        TradeAccount: tradeAccount,
        NumberOfDays: 3,
    }
    msg, err := proto.Marshal( &fillHistReq )
    if err != nil {
        log.Errorf("Failed to marshal HISTORICAL_ORDER_FILLS_REQUEST message: %v\n", err)
        return errors.New("Failed to marshal HISTORICAL_ORDER_FILLS_REQUEST, check proto version?")
    }
    log.Debug("Sending HISTORICAL_ORDER_FILLS_REQUEST")
    //_, err = d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["HISTORICAL_ORDER_FILLS_REQUEST"] ))
    d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["HISTORICAL_ORDER_FILLS_REQUEST"] ))
    //fmt.Println( protojson.Format(&fillHistReq) )
    return err
}

func (d *DtcConnection) CurrentPositions(acctId string) error {
    balReq := dtcproto.CurrentPositionsRequest {
        RequestID: d.nextRequestID(),
        TradeAccount: acctId,
    }
    //describe( balReq.ProtoReflect().Descriptor().FullName() )
    msg, err := proto.Marshal( &balReq )
    if err != nil {
        log.Errorf("Failed to marshal CURRENT_POSITIONS_REQUEST message: %v\n", err)
        return errors.New("Failed to marshal CURRENT_POSITIONS_REQUEST, check proto version?")
    }
    log.Debug("Sending CURRENT_POSITIONS_REQUEST")
    d.conn.Write( PackMessage( msg, dtcproto.DTCMessageType_value["CURRENT_POSITIONS_REQUEST"] ))
    return err
}
