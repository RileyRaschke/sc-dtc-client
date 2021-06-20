package web

import (
	"net/http"
)

func (x *WebServerPlugin) Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] != "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	x.Snapshot(w, r)
}

func (x *WebServerPlugin) Snapshot(w http.ResponseWriter, r *http.Request) {
	stash := x.initStash(r)
	x.textCacheMtx.Lock()
	defer x.textCacheMtx.Unlock()
	stash.Values["Text"] = x.textView
	err := renderTemplate(w, "root.tmpl", stash)
	if err != nil {
		panic(err)
	}
}

func (x *WebServerPlugin) Symbols(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "symbols.tmpl", map[string][]string{
		"Symbols": x.securityStore.GetSymbols(),
	})
	if err != nil {
		panic(err)
	}
}
