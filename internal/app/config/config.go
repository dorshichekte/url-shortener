package config

import (
	"flag"
	"fmt"

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

	flag.Parse()
}

func (c *Config) init() {
	c.initEnv()

	isInstanceEmpty := c.BaseURL == "" || c.ServerAddress == ""
	if isInstanceEmpty {
		c.initFlags()
	}
}
