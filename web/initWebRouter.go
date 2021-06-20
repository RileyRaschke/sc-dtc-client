package web

import (
	"embed"
	"net/http"
)

//go:embed static
var staticContent embed.FS

func (x *WebServerPlugin) initWebRouter() {

	http.Handle("/static/", http.StripPrefix("", http.FileServer(http.FS(staticContent))))

	http.HandleFunc("/", x.Snapshot)
	http.HandleFunc("/test", x.TestBed)
	http.HandleFunc("/symbols", x.Symbols)
}
