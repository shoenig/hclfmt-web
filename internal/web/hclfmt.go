package web

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"gophers.dev/cmds/hclfmt-web/internal/format"
	"gophers.dev/cmds/hclfmt-web/internal/web/static"
	"gophers.dev/pkgs/loggy"
)

type hclFmt struct {
	log  loggy.Logger
	tool *format.Tool
	html *template.Template
}

func newHCLFmt(tool *format.Tool) http.Handler {
	return &hclFmt{
		log:  loggy.New("format"),
		tool: tool,
		html: static.MustParseTemplates(
			"static/html/input.html",
		),
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

	result, err := h.tool.Process(rc)
	if err != nil {
		h.log.Errorf("failed to process body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := response(result)
	_, _ = w.Write([]byte(response))
}

func response(result *format.Result) string {
	if result.Diagnostics.Problematic {
		return result.Diagnostics.Body
	}
	return result.Fmt
}
