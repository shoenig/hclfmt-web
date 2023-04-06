package main

import (
	"os"

	"github.com/shoenig/hclfmt-web/internal/service"
	"github.com/shoenig/loggy"
)

func main() {
	log := loggy.New("main")

	fs, err := service.NewFmtService()
	if err != nil {
		log.Errorf("unable to launch: %v", err)
		os.Exit(1)
	}

	fs.Start()
}
