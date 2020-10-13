package dtc

import (
    "errors"
    log "github.com/sirupsen/logrus"
    "github.com/golang/protobuf/proto"
    "fmt"
    "os"
    "google.golang.org/protobuf/encoding/protojson"
)
/**
* This belongs in connection.go but temporarily here as I debug the lack of responses to other requests
*/
func (d *DtcConnection) _Logon() error {
    logonRequest := LogonRequest{
        Username: d.connArgs.Username,
        Password: d.connArgs.Password,
        Integer_1: 2,
        //TradeMode: TradeModeEnum_TRADE_MODE_UNSET,
        TradeMode: TradeModeEnum_TRADE_MODE_LIVE,
        //TradeMode: TradeModeEnum_TRADE_MODE_SIMULATED,
        HeartbeatIntervalInSeconds: DTC_CLIENT_HEARTBEAT_SECONDS+1,
        ClientName: "go-dtc",
        ProtocolVersion: DTCVersion_value["CURRENT_VERSION"],
    }
    //describe( logonRequest.ProtoReflect().Descriptor().FullName() )
    fmt.Println( protojson.Format(&logonRequest) )

    msg, err := proto.Marshal( &logonRequest )
    if err != nil {
        log.Fatalf("Failed to marshal LogonRequest message: %v\n", err)
        os.Exit(1)
    }

    log.Debug("Sending LOGON_REQUEST")
    d.conn.Write( PackMessage( msg, DTCMessageType_value["LOGON_REQUEST"] ))

    resp, mTypeId := d._GetMessage()

    logonResponse := LogonResponse{}
    if err := proto.Unmarshal(resp, &logonResponse); err != nil {
        log.Fatalln("Failed to parse LogonResponse:", err)
    }
    if logonResponse.Result != LogonStatusEnum_LOGON_SUCCESS {
        /*
        log.WithFields(log.Fields{
            "result": logonResponse.Result,
            "desc": logonResponse.ResultText,
        }).Fatal("Logon Failed")
        */
        log.Fatalf("Logon Failed with result %v and text %v", logonResponse.Result, logonResponse.ResultText)
        return errors.New("Logon Failure")
    }
    log.Debugf("Received %v result: %v", DTCMessageType_name[mTypeId], logonResponse.ResultText)
    fmt.Println( protojson.Format(&logonResponse) )
    return nil
}

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

