package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"gophers.dev/cmds/hclfmt-web/internal/format"
	"gophers.dev/cmds/hclfmt-web/internal/web/static"
	petrify "gophers.dev/cmds/petrify/v5"
)

func Set(router *mux.Router, tool *format.Tool) {
	statics := http.FileServer(&petrify.AssetFS{
		Asset:     static.Asset,
		AssetDir:  static.AssetDir,
		AssetInfo: static.AssetInfo,
		Prefix:    "static",
	})

	// statics
	router.Handle("/hclfmt/static/css/{file}", http.StripPrefix("/hclfmt/static/", statics))

	// api service
	router.Handle("/hclfmt", newHCLFmt(tool)).Methods(http.MethodGet, http.MethodPost)
}
