package config

import (
	"flag"

	hcl2 "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"gophers.dev/pkgs/loggy"
)

func loadConfig(filename string) (Configuration, error) {
	var c Configuration
	var ctx hcl2.EvalContext
	if err := hclsimple.DecodeFile(filename, &ctx, &c); err != nil {
		return c, err
	}
	return c, nil
}

func Load() (Configuration, error) {
	log := loggy.New("config")

	s := flag.String("config", "", "configuration filename")
	flag.Parse()
	if *s == "" {
		log.Infof("using default configuration")
		return Configuration{
			BindAddress: "0.0.0.0",
			BindPort:    9100,
			Service:     "hclfmt",
		}, nil
	}
	log.Infof("loading configuration from %s", *s)
	return loadConfig(*s)
}
