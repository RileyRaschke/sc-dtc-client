package web

import (
    "net/http"
    "html/template"
)

func (x *WebServerPlugin) Snapshot(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    //t:= template.Must(template.New("root.tmpl").ParseFiles([]string{"web/root.html"}...))
    t, err := template.ParseFiles("web/root.html")
    if err != nil {
        panic(err)
    }
    x.textCacheMtx.Lock()
    defer x.textCacheMtx.Unlock()
    err = t.Execute(w, map[string]string{
        "Text": x.textView,
    })
    if err != nil {
        panic(err)
    }
}

func (x *WebServerPlugin) Symbols(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    //t:= template.Must(template.New("root.tmpl").ParseFiles([]string{"web/root.html"}...))
    t, err := template.ParseFiles("web/symbols.html")
    if err != nil {
        panic(err)
    }
    err = t.Execute(w, map[string][]string {
        "Symbols": x.securityStore.GetSymbols(),
    })
    if err != nil {
        panic(err)
    }
}
