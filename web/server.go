package web

import (
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/RileyR387/sc-dtc-client/accounts"
	"github.com/RileyR387/sc-dtc-client/securities"
	log "github.com/sirupsen/logrus"
)

type WebServerPlugin struct {
	stop                chan int
	securityStore       *securities.SecurityStore
	accountStore        *accounts.AccountStore
	startTime           int64
	refreshMicroseconds int
	dStart              time.Time
	inputBuffer         string
	userInputChan       chan rune
	textCacheMtx        sync.Mutex
	textView            string
	mux                 *http.ServeMux
}

const REFRESH_RATE_HZ float64 = 4.0

func New(ss *securities.SecurityStore, as *accounts.AccountStore) *WebServerPlugin {
	microsecondsFloat := (1.0 / REFRESH_RATE_HZ) * 1000 * 1000
	x := &WebServerPlugin{
		make(chan int),
		ss,
		as,
		time.Now().Unix(),
		int(math.Ceil(microsecondsFloat)),
		time.Now(),
		"",
		make(chan rune),
		sync.Mutex{},
		"",
		http.NewServeMux(),
	}
	x.initWebRouter()
	go x.Run()
	return x
}

func (x *WebServerPlugin) ReceiveData(mdu securities.MarketDataUpdate) {
}

type View struct {
	Text string
}

func (x *WebServerPlugin) Run() {
	fmt.Println("Running WebServerPlugin")
	log.Info(fmt.Sprintf("Running WebServerPlugin"))
	go x.listen()

	//var mktData securities.MarketDataUpdate
	frameCnt := 0
	for {
		select {
		case <-x.stop:
			break
		//case stdin, _ := <-x.userInputChan:
		//    x.inputRouter(stdin)
		default:
			if frameCnt >= int(REFRESH_RATE_HZ) {
				frameCnt = 0
			}
			time.Sleep((time.Duration(x.refreshMicroseconds) * time.Microsecond) - time.Since(x.dStart))
			x.cacheText()
			frameCnt++
		}
	}
}
func (x *WebServerPlugin) listen() {
	err := http.ListenAndServe(":8081", x.mux)
	if err != nil {
		panic(err)
	}
}
func (x *WebServerPlugin) Stop() {
	x.stop <- 1
}
