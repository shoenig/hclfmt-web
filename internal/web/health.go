package web

import (
	"io"
	"net/http"

	"gophers.dev/pkgs/loggy"
)

type health struct {
	log loggy.Logger
}

func newHealth() http.Handler {
	return &health{
		log: loggy.New("health"),
	}
}

func (h *health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Tracef("GET %q from %q", r.URL.Path, r.RemoteAddr)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = io.WriteString(w, "ok")
}
