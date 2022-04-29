package main

import (
	"fmt"
	"handystuff/config"
	"log"

	"github.com/spf13/pflag"
)

func main() {
	configFile := pflag.StringP("config", "c", "config.dev.yaml", "specific config file.")
	pflag.Parse()

	conf, err := config.Load(*configFile)
	if err != nil {
		panic(err)
	}

	ginEngine, err := setup(conf)
	if err != nil {
		panic(err)
	}
	log.Fatalln(ginEngine.Run(fmt.Sprintf(":%d", conf.App.Port)))
}
