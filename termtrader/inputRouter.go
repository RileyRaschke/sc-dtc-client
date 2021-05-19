package termtrader

import (
    //"fmt"
    //"math"
    //"strings"
    //tm "github.com/buger/goterm"
    //"time"
    //"strconv"
    //"encoding/json"
    //"sync"
    //log "github.com/sirupsen/logrus"
)

func (x *TermTraderPlugin) inputRouter(input rune) {
    //log.Debugf("Recieved keypress: %v", input)
    switch( input ){
        case 91: // delete
            return
        case 51: // delete
            return
        case 126: // delete
            return
        case 27: // delete
            fallthrough
        case 127: // backspace
            if len(x.inputBuffer) > 0 {
                x.inputBuffer = x.inputBuffer[:len(x.inputBuffer)-1]
            }
            return
        case '\n':
            // Execute the command
            x.inputBuffer = ""
            return
        //case ':':
         //   x.inputMode = "COMMAND"
        //case ''
        default:
            x.inputBuffer += string(input)
    }
}

