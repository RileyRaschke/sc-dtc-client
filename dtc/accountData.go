package dtc

import (
    "errors"
    log "github.com/sirupsen/logrus"
    "github.com/golang/protobuf/proto"
)

func (d *DtcConnection) LoadAccounts() error {
    tradeActReq := TradeAccountsRequest{
        RequestID: d.nextRequestID(),
    }

    msg, err := proto.Marshal( &tradeActReq )
    if err != nil {
        log.Errorf("Failed to marshal TradeAccountsRequest message: %v\n", err)
        return errors.New("Failed to marshal TradeAccountsRequest, check proto version?")
    }

    log.Debug("Sending TRADE_ACCOUNTS_REQUEST")
    _, err = d.conn.Write( PackMessage( msg, DTCMessageType_value["TRADE_ACCOUNTS_REQUEST"] ))
    return err
}

func (d *DtcConnection) AccountBlanaceRefresh() error {
    balReq := AccountBalanceRequest{
        RequestID: d.nextRequestID(),
    }

    msg, err := proto.Marshal( &balReq )
    if err != nil {
        log.Errorf("Failed to marshal ACCOUNT_BALANCE_REQUEST message: %v\n", err)
        return errors.New("Failed to marshal ACCOUNT_BALANCE_REQUEST, check proto version?")
    }

    log.Debug("Sending ACCOUNT_BALANCE_REQUEST")
    _, err = d.conn.Write( PackMessage( msg, DTCMessageType_value["ACCOUNT_BALANCE_REQUEST"] ))
    return err
}
