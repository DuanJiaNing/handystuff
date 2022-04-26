package config

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
