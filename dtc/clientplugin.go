package dtc

import (
	"github.com/RileyR387/sc-dtc-client/securities"
	//    "github.com/RileyR387/sc-dtc-client/accounts"
)

type ClientPlugin interface {
	//New(ss *securities.SecurityStore, as *accounts.AccountStore) *ClientPlugin
	Run()
	Stop()
	ReceiveData(mdu securities.MarketDataUpdate)
}
