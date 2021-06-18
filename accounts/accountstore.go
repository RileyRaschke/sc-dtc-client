package accounts

import (
    "time"
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
type Position = *dtcproto.PositionUpdate
type Order = *dtcproto.OrderUpdate

type Account struct {
    AccountData
    Positions map[string]Position
}

func (a *Account) AddPosition(p Position) {

}

type AccountStore struct {
    acctIds []string
    accounts map[string]Account
    acctUpdateMutex sync.Mutex
    lastUpdated int64
}

func NewAccountStore() *AccountStore {
    return &AccountStore{ []string{}, make(map[string]Account), sync.Mutex{}, time.Now().Unix() }
}

func (as *AccountStore) AddData(msg proto.Message, mTypeId int32) {
    log.Tracef("Adding account data with type %v(%v)", dtcproto.DTCMessageType(mTypeId), mTypeId)
    as.acctUpdateMutex.Lock()
    defer as.acctUpdateMutex.Unlock()

    switch( dtcproto.DTCMessageType(mTypeId) ){

    case dtcproto.DTCMessageType_ORDER_UPDATE:
        order := msg.(Order)
        log.Debugf("Order Update: %v %v %v %v %v %v %v %v",
            order.TradeAccount,
            order.BuySell,
            order.Symbol,
            order.OrderStatus,
            order.Price1,
            order.AverageFillPrice,
            order.OrderQuantity,
            order.FilledQuantity,
        )
        return;
    case dtcproto.DTCMessageType_POSITION_UPDATE:
        position := msg.(Position)
        log.Debugf("Position Update: %v %v %v %v %v",
            position.TradeAccount,
            position.Quantity,
            position.Symbol,
            time.Unix(int64(position.EntryDateTime),0).Format(time.RFC1123),
            position.AveragePrice,
        )
        if _, ok := as.accounts[position.TradeAccount]; !ok {
            // not yet known account...
            return
        }
        as.accounts[position.TradeAccount].Positions[position.Symbol] = position
        return;
    case dtcproto.DTCMessageType_ACCOUNT_BALANCE_UPDATE:
        abu := msg.(*dtcproto.AccountBalanceUpdate)
        if _, ok := as.accounts[abu.TradeAccount]; !ok {
            as.acctIds = append(as.acctIds, abu.TradeAccount)
        }
        //as.accounts[abu.TradeAccount] = abu.(Account)
        as.accounts[abu.TradeAccount] = Account{ abu, make(map[string]Position) }
        log.Infof("Trade account balance: %.2f", as.accounts[abu.TradeAccount].SecuritiesValue)
        as.lastUpdated = time.Now().Unix()
        return;
    default:
        log.Debugf("(accountstore) - Balance or Order Data Received %v(%v)\n%v",
            dtcproto.DTCMessageType_name[mTypeId],
            mTypeId,
            protojson.Format(msg.(protoreflect.ProtoMessage)),
        )
    }
}

func (as *AccountStore) LastUpdated() int64 {
    as.acctUpdateMutex.Lock()
    defer as.acctUpdateMutex.Unlock()
    return as.lastUpdated
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

func (as *AccountStore) GetPositions() []Position {
    as.acctUpdateMutex.Lock()
    defer as.acctUpdateMutex.Unlock()
    res := []Position{}
    for _, id := range as.acctIds {
        for _, val := range as.accounts[id].Positions {
            if val.Quantity != 0 {
                res = append(res, val)
            }
        }
    }
    return res
}

func (as *AccountStore) GetMarginReq() float64 {
    as.acctUpdateMutex.Lock()
    defer as.acctUpdateMutex.Unlock()
    var res float64
    for _, id := range as.acctIds {
        res += as.accounts[id].MarginRequirement
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

