package accounts

import (
    "sync"
    log "github.com/sirupsen/logrus"
    "github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/reflect/protoreflect"
    "github.com/RileyR387/sc-dtc-client/dtcproto"
    //"errors"
    //"encoding/json"
)

type AccountData = *dtcproto.AccountBalanceUpdate

type Account struct {
    AccountData
}

type AccountStore struct {
    acctIds []string
    accounts map[string]Account
    acctUpdateMutex sync.Mutex
}

func NewAccountStore() *AccountStore {
    return &AccountStore{ []string{}, make(map[string]Account), sync.Mutex{} }
}

func (as *AccountStore) AddData(msg proto.Message, mTypeId int32) {
    log.Tracef("Adding account data with type %v(%v)", dtcproto.DTCMessageType(mTypeId), mTypeId)
    as.acctUpdateMutex.Lock()
    defer as.acctUpdateMutex.Unlock()

    switch( dtcproto.DTCMessageType(mTypeId) ){

    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_UPDATE:
        abu := msg.(*dtcproto.AccountBalanceUpdate)
        if _, ok := as.accounts[abu.TradeAccount]; !ok {
            as.acctIds = append(as.acctIds, abu.TradeAccount)
        }
        //as.accounts[abu.TradeAccount] = abu.(Account)
        as.accounts[abu.TradeAccount] = Account{ abu }
        log.Infof("Trade account balance: %.2f", as.accounts[abu.TradeAccount].SecuritiesValue)
        return;

    default:
        log.Debugf("(accountstore) - Balance or Order Data Received %v(%v)\n%v",
            dtcproto.DTCMessageType_name[mTypeId],
            mTypeId,
            protojson.Format(msg.(protoreflect.ProtoMessage)),
        )
    }
}

func (as *AccountStore) GetNetBalance() float64 {
    as.acctUpdateMutex.Lock()
    defer as.acctUpdateMutex.Unlock()
    var res float64
    for _, id := range as.acctIds {
        res += as.accounts[id].SecuritiesValue
    }
    return res
}

func (as *AccountStore) GetCashBalance() float64 {
    as.acctUpdateMutex.Lock()
    defer as.acctUpdateMutex.Unlock()
    var res float64
    for _, id := range as.acctIds {
        res += as.accounts[id].CashBalance
    }
    return res
}

//func (as *AccountStore) GetAccounts() []string {}

