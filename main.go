package main

import (
	"fmt"
	"github.com/RileyR387/sc-dtc-client/dtc"
	"github.com/pborman/getopt/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	Version string
	quit    chan int
	signals chan os.Signal
	client  dtc.DtcClient
)

func init() {
	quit = make(chan int)
	signals = make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
}

func main() {

	args := dtc.ConnectArgs{
		viper.GetString("dtc.Host"),
		viper.GetString("dtc.Port"),
		viper.GetString("dtc.HistPort"),
		viper.GetString("dtc.Username"),
		viper.GetString("dtc.Password"),
		viper.GetBool("dtc.Reconnect"),
	}

	err := client.Connect(args)
	if err != nil {
		log.Fatalf("Failed to connect with: %v\n", err)
	}

	CatchInterupt()

	ScreenServer()

	os.Exit(0)
}

func ScreenServer() {
	for {
		select {
		case <-quit:
			if client.Connected() {
				client.Terminate()
			}
			return
		default:
			if !client.Terminated() {
				time.Sleep(time.Second)
			} else {
				log.Printf("Client unexpectedly disconnected from server\n")
				// TODO: Attempt reconnect? Client connection will need timeout function..
				time.Sleep(time.Second)
				return
			}
		}
	}
}

func CatchInterupt() {
	log.Infof("%v %v - Started", me, Version)
	go func() {
		i := <-signals
		log.Printf("Received signal: %v, exiting...\n", i)
		quit <- 0
	}()
}

func usage(msg ...string) {
	if len(msg) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", msg[0])
	}
	// strip off the first line of generated usage
	b := &strings.Builder{}
	getopt.PrintUsage(b)
	u := strings.SplitAfterN(b.String(), "\n", 2)
	fmt.Printf(`Usage: %s

OPTIONS
%s
`, me, u[1])

	os.Exit(1)
}

func ShowVersion() {
	fmt.Printf("%v %v\n", me, Version)
	os.Exit(0)
}
