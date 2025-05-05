package config

import (
	"flag"

	"github.com/caarlos0/env"
	"go.uber.org/zap"

	"url-shortener/internal/app/constants"
)

func NewConfig(logger *zap.Logger) *Config {
	cfg := &Config{
		Logger: logger,
		App:    &AppConfig{},
	}
	cfg.init()
	return cfg
}

func (c *Config) initEnv() {
	if err := env.Parse(c); err != nil {
		c.Logger.Error(constants.ErrParsingConfig.Error(), zap.Error(err))
	}
}

func (c *Config) initFlags() {
	flag.StringVar(&c.App.ServerAddress, "a", c.App.ServerAddress, "server address")
	flag.StringVar(&c.App.BaseURL, "b", c.App.BaseURL, "base host URL")
	flag.StringVar(&c.App.DatabaseDSN, "d", c.App.DatabaseDSN, "Connect DB string")
	flag.StringVar(&c.App.FileStoragePath, "f", c.App.FileStoragePath, "file storage path")

	flag.Parse()
}

func (c *Config) initDefaultValue() {
	if c.App.ServerAddress == "" {
		c.App.ServerAddress = DefaultAddress
	}

	if c.App.BaseURL == "" {
		c.App.BaseURL = DefaultAddressWithProtocol
	}

	if c.App.FileStoragePath == "" {
		c.App.FileStoragePath = DefaultStoragePath
	}
}

func (c *Config) init() {
	c.initEnv()
	c.initFlags()
	c.initDefaultValue()
}
