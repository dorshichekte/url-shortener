package config

import (
	"flag"
	"os"

	customerror "url-shortener/internal/pkg/error"
)

func New() (*Env, error) {
	config := &Env{}

	err := config.init()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Env) initEnv() {
	c.ServerAddress = os.Getenv("SERVER_ADDRESS")
	c.DatabaseDSN = os.Getenv("DATABASE_DSN")
	c.FileStoragePath = os.Getenv("FILE_STORAGE_PATH")
	c.BaseURL = os.Getenv("BASE_URL")
	c.AccessSecretKey = os.Getenv("ACCESS_SECRET_KEY")
}

func (c *Env) initFlags() {
	flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "Server address")
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "Connect DB string")
	flag.StringVar(&c.FileStoragePath, "f", c.FileStoragePath, "File storage path")
	flag.StringVar(&c.BaseURL, "b", c.BaseURL, "Base host URL")
	flag.StringVar(&c.AccessSecretKey, "ac", c.AccessSecretKey, "Access secret key")

	flag.Parse()
}

func (c *Env) initDefaultValue() {
	if c.ServerAddress == "" {
		c.ServerAddress = defaultAddress
	}

	if c.BaseURL == "" {
		c.BaseURL = defaultAddressWithProtocol
	}

	if c.FileStoragePath == "" {
		c.FileStoragePath = defaultStoragePath
	}

	if c.AccessSecretKey == "" {
		c.AccessSecretKey = defaultAccessSecret
	}
}

func (c *Env) valid() error {
	var emptyVariables []string

	if c.ServerAddress == "" {
		emptyVariables = append(emptyVariables, "SERVER_ADDRESS")
	}

	if c.DatabaseDSN == "" {
		emptyVariables = append(emptyVariables, "DATABASE_DSN")
	}

	if c.FileStoragePath == "" {
		emptyVariables = append(emptyVariables, "FILE_STORAGE_PATH")
	}

	if c.BaseURL == "" {
		emptyVariables = append(emptyVariables, "BASE_URL")
	}

	if len(emptyVariables) != 0 {
		return customerror.NewWithData(errEnvEmptyVariables, emptyVariables)
	}

	return nil
}

func (c *Env) init() (err error) {
	c.initEnv()
	c.initFlags()
	c.initDefaultValue()

	err = c.valid()

	return err
}
