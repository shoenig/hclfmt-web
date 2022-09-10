package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shoenig/hclfmt-web/internal/config"
	"github.com/shoenig/hclfmt-web/internal/web"
)

func initWeb(fs *FmtService) error {
	address := config.Get()
	fs.log.Tracef("setting up web server @ %s", address)

	router := mux.NewRouter()
	web.Set(router, fs.tool)

	go func() {
		if err := (&http.Server{
			Addr:    address,
			Handler: router,
		}).ListenAndServe(); err != nil {
			fs.log.Errorf("failed to listen and serve", err)
		}
	}()

	return nil
}
