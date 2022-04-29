package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	AESCrypt AESCryptConfig
	Tick     TickConfig
	App      AppConfig
}

type AppConfig struct {
	Port int
}

type TickConfig struct {
	Enable bool
	Addr   string
}

type AESCryptConfig struct {
	Keys map[string]string
}

func Load(file string) (*Conf, error) {
	fmt.Println("load config file:", file)
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	conf := &Conf{}
	err = yaml.Unmarshal(bs, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
