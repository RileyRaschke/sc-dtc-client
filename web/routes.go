package web

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates
var htdocs embed.FS

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

func (x *WebServerPlugin) TestBed(w http.ResponseWriter, r *http.Request) {
	stash := x.initStash(r)
	//t:= template.Must(template.New("root.tmpl").ParseFiles([]string{"web/root.html"}...))
	//t, err := template.ParseFS(htdocs, "*.html")
	t, err := template.ParseFS(htdocs, "templates/root.tmpl")
	if err != nil {
		panic(err)
	}
	x.textCacheMtx.Lock()
	defer x.textCacheMtx.Unlock()
	stash.Values["Text"] = x.textView
	err = t.Execute(w, stash)

	if err != nil {
		panic(err)
	}
}
