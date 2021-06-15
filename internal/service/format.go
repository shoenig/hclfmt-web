package service

import (
	"gophers.dev/cmds/hclfmt-web/internal/format"
	"gophers.dev/pkgs/loggy"
)

const (
	maxReqLen = 1 * 1024 * 1024 // 1 megabyte
)

type FmtService struct {
	tool *format.Tool
	log  loggy.Logger
}

func NewFmtService() (*FmtService, error) {
	return initialize(&FmtService{
		tool: format.NewTool(maxReqLen),
		log:  loggy.New("hclfmt"),
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
