package dtc

import (
    "errors"
    log "github.com/sirupsen/logrus"
    "github.com/golang/protobuf/proto"
    "fmt"
    "google.golang.org/protobuf/encoding/protojson"
)

/**
* Should trigger repsonse handled by DtcConnection._Listen() -> _RouteMessage() but seeing no bytes on wire...
*/
func (d *DtcConnection) LoadAccounts() error {
    tradeActReq := TradeAccountsRequest{
        RequestID: d.nextRequestID(),
    }
    describe( tradeActReq.ProtoReflect().Descriptor().FullName() )
    fmt.Println( protojson.Format(&tradeActReq) )

    msg, err := proto.Marshal( &tradeActReq )
    if err != nil {
        log.Errorf("Failed to marshal TradeAccountsRequest message: %v\n", err)
        return errors.New("Failed to marshal TradeAccountsRequest, check proto version?")
    }
    //msg = append(msg, 0x00) // not spec for protobuf, but tried it anyways...
    describe( msg )

    log.Debug("Sending TRADE_ACCOUNTS_REQUEST")
    _, err = d.conn.Write( PackMessage( msg, DTCMessageType_value["TRADE_ACCOUNTS_REQUEST"] ))
    return err
}

/**
* Should trigger repsonse handled by DtcConnection._Listen() -> _RouteMessage() but seeing no bytes on wire...
*/
func (d *DtcConnection) AccountBlanaceRefresh() error {
    balReq := AccountBalanceRequest{
        RequestID: d.nextRequestID(),
    }
    describe( balReq.ProtoReflect().Descriptor().FullName() )
    fmt.Println( protojson.Format(&balReq) )

    msg, err := proto.Marshal( &balReq )
    if err != nil {
        log.Errorf("Failed to marshal ACCOUNT_BALANCE_REQUEST message: %v\n", err)
        return errors.New("Failed to marshal ACCOUNT_BALANCE_REQUEST, check proto version?")
    }
    describe( msg )

    log.Debug("Sending ACCOUNT_BALANCE_REQUEST")
    _, err = d.conn.Write( PackMessage( msg, DTCMessageType_value["ACCOUNT_BALANCE_REQUEST"] ))
    return err
}

