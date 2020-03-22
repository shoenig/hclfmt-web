package main

import (
	"os"

	"gophers.dev/cmds/hclfmt-web/internal/config"
	"gophers.dev/cmds/hclfmt-web/internal/service"
	"gophers.dev/pkgs/loggy"
)

// generate HTML static file(s)
//go:generate go run gophers.dev/cmds/petrify/v5/cmd/petrify -prefix internal/web/ -o internal/web/static/generated.go -pkg static internal/web/static/...

func main() {
	log := loggy.New("main")

	configuration, err := config.Load()
	if err != nil {
		log.Errorf("unable to load config: %v", err)
		os.Exit(1)
	}

	fs, err := service.NewFmtService(configuration)
	if err != nil {
		log.Errorf("unable to launch: %v", err)
		os.Exit(1)
	}

	fs.Start()
}
