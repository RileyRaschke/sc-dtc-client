package web

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed static
var staticContent embed.FS

//go:embed static/favicon.ico
var favicon embed.FS

func (x *WebServerPlugin) initWebRouter() {
	x.mux.Handle("/static/", http.FileServer(http.FS(staticContent)))
	//x.mux.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(http.FS(favicon))))
	x.mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})
	x.mux.HandleFunc("/", x.Root)
	x.mux.HandleFunc("/symbols", x.Symbols)
}
