package service

import (
	"github.com/shoenig/hclfmt-web/internal/format"
	"github.com/shoenig/loggy"
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
		log:  loggy.New("hcl"),
	})
}

func (fs *FmtService) Start() {
	fs.log.Infof("--- starting ---")
	select {
	// intentionally left blank
	}
}

func initialize(fs *FmtService) (*FmtService, error) {
	return fs, initWeb(fs)
}
