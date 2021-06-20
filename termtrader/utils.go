package termtrader

import (
	//"fmt"
	"github.com/gookit/color"
	"strings"
)

func ColorizeChangeString(v string) string {
	if strings.HasPrefix(v, "-") || strings.HasPrefix(v, "'-") {
		red := color.FgRed.Render
		return red(v)

	} else {
		green := color.FgGreen.Render
		return green(v)

	}
}
