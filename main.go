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
	fs, err := service.NewFmtService(config.Configuration{
		WebServer: config.WebServer{
			BindAddress: "0.0.0.0",
			Port:        9800,
		},
	})
	if err != nil {
		loggy.New("main").Errorf("unable to launch: %v", err)
		os.Exit(1)
	}
	fs.Start()
}
