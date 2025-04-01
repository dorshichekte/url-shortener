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
	flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "server address")
	flag.StringVar(&c.BaseURL, "b", c.BaseURL, "base host URL")
	flag.StringVar(&c.FileStoragePath, "f", c.FileStoragePath, "file storage path")

	flag.Parse()
}

func (c *Config) initDefault() {
	if c.ServerAddress == "" {
		c.ServerAddress = DefaultAddress
	}

	if c.BaseURL == "" {
		c.BaseURL = DefaultAddressWithProtocol
	}

	if c.FileStoragePath == "" {
		c.FileStoragePath = DefaultStoragePath
	}
}

func (c *Config) init() {
	c.initEnv()
	c.initFlags()
	c.initDefault()
}
