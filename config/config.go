package config

import (
	"handystuff/common/log"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Conf config

type config struct {
	AESCrypt struct {
		C    int
		Keys map[string]string
	}
	Tick struct {
		Enable bool
		Addr   string
	}
	App struct {
		Port int
	}
}

func Load(file string) error {
	log.Printlnf("load config from %s ...", file)
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	c := yaml.Unmarshal(bs, &Conf)
	return c
}
