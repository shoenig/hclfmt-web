package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"gophers.dev/cmds/hclfmt-web/internal/web"
	"gophers.dev/pkgs/loggy"
)

type initer func(*FmtService) error

func initWeb(fs *FmtService) error {
	fs.log.Tracef("setting up web server @ %s", fs.config.Address())

	router := mux.NewRouter()
	web.Set(router, fs.tool)

	server, err := fs.config.Server(router)
	if err != nil {
		return errors.Wrap(err, "unable to create web server")
	}

	go func(h http.Handler) {
		err := server.ListenAndServe()
		loggy.New("listener").Errorf("server stopped serving: %v", err)
		os.Exit(1)
	}(router)

	return nil
}
