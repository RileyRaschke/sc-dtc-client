package main

import (
    log "github.com/sirupsen/logrus"
    "runtime"
    "path"
    "strconv"
)

func initLogger(logLevel string) {
    formatter := &log.TextFormatter{
        FullTimestamp: true,
        TimestampFormat: "2006-01-02 15:04:05.000",
        PadLevelText: true,
        QuoteEmptyFields: true,
    }

    log.SetFormatter( formatter )

    level, _ := log.ParseLevel( logLevel )

    log.SetLevel( level )

    if log.GetLevel() == log.TraceLevel {
        log.SetReportCaller(true)
        formatter.CallerPrettyfier = LogPrettyTrace
    }
}

func LogPrettyTrace(f *runtime.Frame) (function string, file string){
    //function = path.Base(f.Function)
    function = ""
    file = " " + path.Base(f.File) + ":" + strconv.Itoa(f.Line)
    return function, file
}

