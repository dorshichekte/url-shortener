package config

import (
	"flag"
	"os"
)

var (
	serverAddress   = flag.String("a", DefaultAddress, "server address")
	baseURL         = flag.String("b", DefaultAddressWithProtocol, "base host URL")
	fileStoragePath = flag.String("f", StoragePath, "file storage path")
)

func NewConfig() *Config {
	cfg := &Config{}
	cfg.init()
	return cfg
}

func (c *Config) initEnv() {
	c.ServerAddress = os.Getenv("SERVER_ADDRESS")
	c.BaseURL = os.Getenv("BASE_URL")
	c.FileStoragePath = os.Getenv("FILE_STORAGE_PATH")
}

func (c *Config) initFlags() {
	if c.ServerAddress == "" {
		c.ServerAddress = *serverAddress
	}

	if c.BaseURL == "" {
		c.BaseURL = *baseURL
	}

	if c.FileStoragePath == "" {
		c.FileStoragePath = *fileStoragePath
	}
}

func (c *Config) init() {
	c.initEnv()
	c.initFlags()
}
