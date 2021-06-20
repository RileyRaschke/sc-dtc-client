package termtrader

import (
	"fmt"
	//"github.com/gookit/color"
	"bufio"
	"math"
	"os"
	"os/exec"
	"time"
	//"strconv"
	//"encoding/json"
	//"sync"
	tm "github.com/buger/goterm"
	log "github.com/sirupsen/logrus"
	//"google.golang.org/protobuf/proto"
	//"google.golang.org/protobuf/encoding/protojson"
	//"google.golang.org/protobuf/reflect/protoreflect"
	"github.com/RileyR387/sc-dtc-client/accounts"
	"github.com/RileyR387/sc-dtc-client/securities"
)

//const REFRESH_RATE_HZ float64 = 60.0
const REFRESH_RATE_HZ float64 = 30.0

type TermTraderPlugin struct {
	stop                chan int
	securityStore       *securities.SecurityStore
	accountStore        *accounts.AccountStore
	startTime           int64
	refreshMicroseconds int
	dStart              time.Time
	inputBuffer         string
	userInputChan       chan rune
}

func New(ss *securities.SecurityStore, as *accounts.AccountStore) *TermTraderPlugin {
	microsecondsFloat := (1.0 / REFRESH_RATE_HZ) * 1000 * 1000
	x := &TermTraderPlugin{
		make(chan int),
		ss,
		as,
		time.Now().Unix(),
		int(math.Ceil(microsecondsFloat)),
		time.Now(),
		"",
		make(chan rune),
	}
	go x.Run()
	return x
}

func (x *TermTraderPlugin) ReceiveData(mdu securities.MarketDataUpdate) {
}

func (x *TermTraderPlugin) Run() {
	fmt.Println("Running TermTraderPlugin")
	log.Info(fmt.Sprintf("Running TermTraderPlugin"))
	//var mktData securities.MarketDataUpdate

	tm.Clear() // Clear current screen
	x.runInputListener()
	frameCnt := 0
	for {
		select {
		case <-x.stop:
			break
		case stdin, _ := <-x.userInputChan:
			x.inputRouter(stdin)
		default:
			if frameCnt >= int(REFRESH_RATE_HZ) {
				tm.Clear() // Clear current screen
				frameCnt = 0
			}
			time.Sleep((time.Duration(x.refreshMicroseconds) * time.Microsecond) - time.Since(x.dStart))
			x.draw()
			frameCnt++
		}
	}
}

func (x *TermTraderPlugin) runInputListener() {
	//ch := make(chan string)
	go func(ch chan rune) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		reader := bufio.NewReader(os.Stdin)
		for {
			char, _, err := reader.ReadRune()
			if err != nil {
				log.Errorf("Error reading rune from stdin buffer: {}", err)
			}
			ch <- char
		}
	}(x.userInputChan)
}

func (x *TermTraderPlugin) Stop() {
	x.stop <- 1
}
