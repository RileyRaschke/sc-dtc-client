package main

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "os"
    "os/signal"
    "time"
    "syscall"
    "strings"
    "github.com/pborman/getopt/v2"
    "github.com/spf13/viper"
    "github.com/RileyR387/sc-dtc-client/dtc"
)

var (
    version = "undefined"
    quit chan int
    signals chan os.Signal
    client dtc.DtcClient
)

func init() {
    quit    = make(chan int)
    signals = make(chan os.Signal)
    signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
}

func main() {

    args := dtc.ConnectArgs {
        viper.GetString("dtc.Host"),
        viper.GetString("dtc.Port"),
        viper.GetString("dtc.HistPort"),
        viper.GetString("dtc.Username"),
        viper.GetString("dtc.Password"),
        viper.GetBool("dtc.Reconnect"),
    }

    err := client.Connect( args )
    if err != nil {
        log.Fatalf("Failed to connect with: %v\n", err)
    }

    CatchInterupt()

    for {
        select {
        case <-quit:
            if client.Connected() {
                client.Terminate()
            }
            return
        default:
            if !client.Terminated() {
                time.Sleep( time.Second )
            } else {
                log.Printf("Client unexpectedly disconnected from server\n")
                // TODO: Attempt reconnect? Client connection will need timeout function..
                time.Sleep( time.Second )
                return
            }
        }
    }

    os.Exit(0)
}

func CatchInterupt(){
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

Activity log is written to STDERR.

OPTIONS
%s
`, me, u[1])

    os.Exit(1)
}
