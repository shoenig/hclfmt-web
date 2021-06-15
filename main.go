package main

import (
	"os"

	"gophers.dev/cmds/hclfmt-web/internal/service"
	"gophers.dev/pkgs/loggy"
)

// generate HTML static file(s)
//go:generate go run gophers.dev/cmds/petrify/v5/cmd/petrify -prefix internal/web/ -o internal/web/static/generated.go -pkg static internal/web/static/...

func main() {
	log := loggy.New("main")

	fs, err := service.NewFmtService()
	if err != nil {
		log.Errorf("unable to launch: %v", err)
		os.Exit(1)
	}

	fs.Start()
}
