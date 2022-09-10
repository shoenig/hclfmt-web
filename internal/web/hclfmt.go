package web

import (
	"embed"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/shoenig/hclfmt-web/internal/format"
	"gophers.dev/pkgs/loggy"
)

type hclFmt struct {
	log  loggy.Logger
	tool *format.Tool
	html *template.Template
}

func newHCLFmt(fs embed.FS, tool *format.Tool) http.Handler {
	html, err := template.ParseFS(fs, "static/html/*.html")
	if err != nil {
		panic(err)
	}
	return &hclFmt{
		log:  loggy.New("format"),
		tool: tool,
		html: html,
	}
}

func (h *hclFmt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.post(w, r)
	}
}

func (h *hclFmt) get(w http.ResponseWriter, r *http.Request) {
	h.log.Tracef("GET %q from %s", r.URL.Path, r.RemoteAddr)
	w.Header().Set("Content-Type", "text/html")
	if err := h.html.Execute(w, nil); err != nil {
		h.log.Errorf("failed to render html: %v", err)
		return
	}
}

type page struct {
	Path    string
	Statics string
}

func (h *hclFmt) post(w http.ResponseWriter, r *http.Request) {
	h.log.Tracef("POST %q from %s", r.URL.Path, r.RemoteAddr)
	w.Header().Set("Content-Type", "text/plain")

	if err := r.ParseForm(); err != nil {
		h.log.Errorf("failed to process form: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rc := ioutil.NopCloser(
		strings.NewReader(r.PostForm.Get("fmttext")),
	)

	body, err := h.tool.Process(rc)
	if err != nil {
		h.log.Errorf("failed to process body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	markdown := r.PostForm.Get("mdown")
	_, _ = w.Write([]byte(response(markdown, body)))

}

func response(markdown string, result *format.Result) string {
	if result.Diagnostics.Problematic {
		return result.Diagnostics.Body
	}
	return format.Markdown(markdown, result.Fmt)
}
