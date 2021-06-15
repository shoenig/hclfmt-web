package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"gophers.dev/cmds/hclfmt-web/internal/config"

	"gophers.dev/cmds/hclfmt-web/internal/web"
)

type initer func(*FmtService) error

func initWeb(fs *FmtService) error {
	fs.log.Tracef("setting up web server @ %s", config.Address())

	router := mux.NewRouter()
	web.Set(router, fs.tool)

	service, err := config.GetService()
	if err != nil {
		return err
	}

	go func() {
		if lErr := (&http.Server{
			Addr:      config.Address(),
			TLSConfig: service.ServerTLSConfig(),
			Handler:   router,
		}).ListenAndServeTLS("", ""); lErr != nil {
			fs.log.Errorf("failed to listen and serve Connect TLS:", lErr)
		}
	}()

	return nil
}
