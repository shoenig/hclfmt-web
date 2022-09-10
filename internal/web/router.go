package web

import (
	"embed"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shoenig/hclfmt-web/internal/format"
)

//go:embed static/*
var fs embed.FS

func Set(router *mux.Router, tool *format.Tool) {
	// handle statics
	statics := http.FileServer(http.FS(fs))
	router.Handle("/static/css/{file}", http.StripPrefix("/hclfmt/static/", statics))

	// handle health check
	router.Handle("/health", newHealth()).Methods(http.MethodGet)

	// handle api service
	router.Handle("/", newHCLFmt(fs, tool)).Methods(http.MethodGet, http.MethodPost)
}
