package web

import (
	"embed"
	"github.com/oxtoacart/bpool"
	"html/template"
	"net/http"
)

var Version string

//go:embed templates
var tmplDir embed.FS

var tmpl *template.Template
var bufpool *bpool.BufferPool

func init() {
	bufpool = bpool.NewBufferPool(48)
	var err error
	tmpl = template.New("pages")
	tmpl = tmpl.Funcs(template.FuncMap{
		"toJS":   toJS,
		"toHTML": toHTML,
	})
	tmpl, err = tmpl.ParseFS(tmplDir, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	buf := bufpool.Get()
	err := tmpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		return err
	}
	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch data.(type) {
	case *Stash:
		w.Header().Set("Content-Security-Policy",
			"default-src 'self' ;"+
				"script-src 'self' 'nonce-"+data.(*Stash).Nonce+"' https://cdnjs.cloudflare.com/ajax/libs/;"+
				"font-src https://fonts.gstatic.com;"+
				"img-src 'self';"+
				"frame-src 'self';"+
				"style-src 'self' 'nonce-"+data.(*Stash).Nonce+"' https://cdnjs.cloudflare.com/ajax/libs/ https://fonts.googleapis.com/;"+
				"connect-src 'self';",
		)
	default:
	}
	w.Header().Set("Version", Version)
	buf.WriteTo(w)
	bufpool.Put(buf)
	return nil
}

func toJS(s string) template.JS {
	return template.JS(s)
}

func toHTML(s string) template.HTML {
	return template.HTML(s)
}
