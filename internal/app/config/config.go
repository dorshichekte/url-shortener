package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

func NewConfig() *Config {
	cfg := &Config{}
	cfg.init()
	return cfg
}

func (c *Config) initEnv() {
	if err := env.Parse(c); err != nil {
		fmt.Println(err)
	}
}

func (c *Config) initFlags() {
	flag.StringVar(&c.ServerAddress, "a", DefaultAddress, "server address")
	flag.StringVar(&c.BaseURL, "b", DefaultAddressWithProtocol, "base host URL")
	flag.StringVar(&c.FileStoragePath, "f", StoragePath, "file storage path")

	flag.Parse()
}

func (c *Config) init() {
	c.initEnv()

	isInstanceEmpty := c.BaseURL == "" || c.ServerAddress == "" || c.FileStoragePath == ""
	if isInstanceEmpty {
		c.initFlags()
	}

	log.Println(c, "config")
}
