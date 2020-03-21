package service

import (
	"gophers.dev/cmds/hclfmt-web/internal/config"
	"gophers.dev/cmds/hclfmt-web/internal/format"
	"gophers.dev/pkgs/loggy"
)

type FmtService struct {
	tool   *format.Tool
	config config.Configuration
	log    loggy.Logger
}

func NewFmtService(config config.Configuration) (*FmtService, error) {
	return initialize(&FmtService{
		config: config,
		tool:   format.NewTool(config.WebServer.MaxRequestLen),
		log:    loggy.New("fmt-service"),
	})
}

func (fs *FmtService) Start() {
	fs.log.Infof("--- starting ---")
	select {
	// intentionally left blank
	}
}

func initialize(fs *FmtService) (*FmtService, error) {
	for _, f := range []initer{
		initWeb,
	} {
		if err := f(fs); err != nil {
			return nil, err
		}
	}
	return fs, nil
}
