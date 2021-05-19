package termtrader
import (
    //"fmt"
    "strings"
    "github.com/gookit/color"
)

func ColorizeChangeString(v string) string {
    if strings.HasPrefix(v,"-") || strings.HasPrefix(v, "'-") {
        red := color.FgRed.Render
        return red(v)

    } else {
        green := color.FgGreen.Render
        return green(v)

    }
}
