package main

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
)

func main() {
	configFile := pflag.StringP("config", "c", "config.dev.yaml", "specific config file.")
	pflag.Parse()

	conf, err := loadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	ginEngine, err := setup(conf)
	if err != nil {
		panic(err)
	}
	log.Fatalln(ginEngine.Run(fmt.Sprintf(":%d", conf.App.Port)))
}
