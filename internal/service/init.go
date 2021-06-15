package service

import (
	"net/http"

	"github.com/gorilla/mux"

	"gophers.dev/cmds/hclfmt-web/internal/web"
)

type initer func(*FmtService) error

func initWeb(fs *FmtService) error {
	fs.log.Tracef("setting up web server @ %s/%s", fs.config.BindAddress, fs.config.Service)

	router := mux.NewRouter()
	web.Set(router, fs.tool)

	service, err := fs.config.GetService()
	if err != nil {
		return err
	}

	go func() {
		if lErr := (&http.Server{
			Addr:      fs.config.Address(),
			TLSConfig: service.ServerTLSConfig(),
			Handler:   router,
		}).ListenAndServeTLS("", ""); lErr != nil {
			fs.log.Errorf("failed to listen and serve Connect TLS:", lErr)
		}
	}()

	return nil
}
