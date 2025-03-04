package config

import (
	"flag"
	"sync"

	"github.com/caarlos0/env"
)

const (
	DefaultAddress             = "localhost:8080"
	DefaultAddressWithProtocol = "http://localhost:8080"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
}

var (
	Instance *Config
	once     sync.Once
)

func initEnv(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return err
	}

	return nil
}

func initFlags(cfg *Config) {
	flag.StringVar(&cfg.ServerAddress, "a", DefaultAddress, "server address")
	flag.StringVar(&cfg.BaseURL, "b", DefaultAddressWithProtocol, "base host URL")

	flag.Parse()
}

func Create() {
	Instance = &Config{}

	err := initEnv(Instance)
	if err == nil {
		return
	}

	initFlags(Instance)
}

func Get() *Config {
	if Instance == nil {
		once.Do(func() {
			Create()
		})
	}

	return Instance
}
