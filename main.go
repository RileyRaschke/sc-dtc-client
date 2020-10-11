package main

import (
    "fmt"
    //"log"
    log "github.com/sirupsen/logrus"
    "os"
    "os/signal"
    "time"
    "syscall"
    "path/filepath"
    "strings"
    "github.com/pborman/getopt/v2"
    "github.com/spf13/viper"
    "github.com/RileyR387/sc-dtc-client/dtc"
)

var (
    version = "undefined"
    me = filepath.Base(os.Args[0])
    yamlFile = fmt.Sprintf("%s.yaml", me)
    envPrefix = "DTCCLIENT"
    configSearchPaths = []string {".", "./etc", "$HOME/.dtc-client-go/", "$HOME/etc", "/etc"}
    genConfig = getopt.BoolLong("genconfig", 'x', "Write example config to \"./" + yamlFile + "\"")
    /*
    iLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    wLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    eLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    */

    client *dtc.DtcClient
)

func init() {
    //log.SetPrefix("[main] ")
    viper.SetConfigName(yamlFile)
    viper.SetConfigType("yaml")
    viper.SetEnvPrefix(envPrefix)
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    viper.AutomaticEnv()
    for _, p := range configSearchPaths {
        viper.AddConfigPath(p)
    }
    viper.SetDefault("dtc.Host", "127.0.0.1")
    viper.SetDefault("dtc.Port", "11099")
    viper.SetDefault("dtc.HistPort", "11098")
    viper.SetDefault("dtc.Username", "")
    viper.SetDefault("dtc.Password", "")
}

func main() {
    getopt.SetUsage(func() { usage() })
    getopt.Parse()

    if *genConfig {
        configWrite()
        os.Exit(1)
        return
    }

    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            log.Printf("%v\n", err)
            usage( fmt.Sprintf("\nTry: %s --genconfig\n", me) )
            os.Exit(1)
        } else {
            log.Fatalf("Failed to parse config: %v\n", err)
            os.Exit(0)
        }
    }

    args := dtc.ConnectArgs {
        viper.GetString("dtc.Host"),
        viper.GetString("dtc.Port"),
        viper.GetString("dtc.HistPort"),
        viper.GetString("dtc.Username"),
        viper.GetString("dtc.Password"),
    }

    client, err := dtc.Connect( args )
    if err != nil {
        log.Fatalf("Failed to connect with: %v\n", err)
    }

    quit := make(chan int)
    CatchInterupt(quit)

    for {
        select {
        case <-quit:
            if client.Connected() {
                client.Disconnect()
            }
            return
        default:
            if client.Connected() {
                time.Sleep( 1 * time.Second )
            } else {
                log.Printf("Client unexpectedly disconnected from server\n")
                /*
                * TODO: Attempt reconnect? Client connection will need timeout function..
                */
                time.Sleep( 1 * time.Second )
                return
            }
        }
    }
    os.Exit(0)
}

func CatchInterupt(quit chan int){
    c:= make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        i := <-c
        log.Printf("Received signal: (%v), exiting...\n", i)
        quit <- 0
    }()
}

func configWrite(){
    viper.SafeWriteConfigAs( fmt.Sprintf("./%s", yamlFile) )
    log.Printf("Wrote example config to: \"./%s\", feel free to move to: %v", yamlFile, configSearchPaths[1:])
    log.Println(`

Alternatively, set the following environment variables:

export ` + envPrefix + `_HOST='127.0.0.1'
export ` + envPrefix + `_PORT='11099'
export ` + envPrefix + `_HISTPORT='11098'
export ` + envPrefix + `_USERNAME=''
export ` + envPrefix + `_PASSWORD=''
`)
    os.Exit(0)
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

