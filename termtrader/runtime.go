package termtrader

import (
    "fmt"
	tm "github.com/buger/goterm"
	"time"
    "github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/reflect/protoreflect"
)

type TermTraderPlugin struct {
    ReceiveData chan *proto.Message
    lastMsgJson string
}

func New() *TermTraderPlugin {
    x := &TermTraderPlugin{
        make(chan *proto.Message),
        "",
    }
    go x.Run()
    return x
}

func (x *TermTraderPlugin) Run() {
    fmt.Println("Running TermTraderPlugin")
    var msg *proto.Message
    for {
        select {
        case msg = <-x.ReceiveData:
            //fmt.Println( protojson.Format((*msg).(protoreflect.ProtoMessage)))
            x.lastMsgJson = protojson.Format((*msg).(protoreflect.ProtoMessage))
        }
    }
}

func (x *TermTraderPlugin) Draw() {
    // By moving cursor to top-left position we ensure that console output
    // will be overwritten each time, instead of adding new.
	tm.Clear() // Clear current screen
    for {
        tm.MoveCursor(1, 1)

        tm.Println("Current Time:", time.Now().Format(time.RFC1123))

        tm.Flush() // Call it every time at the end of rendering

        time.Sleep(time.Second)
    }
}

