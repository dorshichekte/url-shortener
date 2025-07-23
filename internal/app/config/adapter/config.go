package config

import (
	"os"
	"strconv"
	"time"

	"url-shortener/internal/pkg/constants"
)

// New создает экземпляр HTTPAdapter конфига.
func New(serverAddress string) *HTTPAdapter {
	var config HTTPAdapter

	config.init(serverAddress)

	return &config
}

func (c *HTTPAdapter) init(serverAddress string) {
	c.Server.Address = serverAddress
	c.initEnv()
	c.initDefaultValues()
}

func (c *HTTPAdapter) initEnv() {
	c.Server.ReadHeaderTimeout = parseDurationFromEnv("SERVER_READ_HEADER_TIMEOUT")
	c.Server.ReadTimeout = parseDurationFromEnv("SERVER_READ_TIMEOUT")
	c.Server.ShutdownTimeout = parseDurationFromEnv("SERVER_SHUTDOWN_TIMEOUT")
	c.Server.WriteTimeout = parseDurationFromEnv("SERVER_WRITE_TIMEOUT")

	c.Router.Shutdown = parseDurationFromEnv("ROUTER_SHUTDOWN")
	c.Router.Timeout = parseDurationFromEnv("ROUTER_TIMEOUT")
}

func (c *HTTPAdapter) initDefaultValues() {
	if c.Server.ReadHeaderTimeout == 0 {
		c.Server.ReadHeaderTimeout = constants.DefaultTimeRequest
	}

	if c.Server.ReadTimeout == 0 {
		c.Server.ReadTimeout = constants.DefaultTimeRequest
	}

	if c.Server.ShutdownTimeout == 0 {
		c.Server.ShutdownTimeout = constants.DefaultTimeRequest
	}

	if c.Server.WriteTimeout == 0 {
		c.Server.WriteTimeout = constants.DefaultTimeRequest
	}

	if c.Router.Shutdown == 0 {
		c.Router.Shutdown = constants.DefaultTimeRequest
	}

	if c.Router.Timeout == 0 {
		c.Router.Timeout = constants.DefaultTimeRequest
	}
}

func parseDurationFromEnv(key string) time.Duration {
	envValue := os.Getenv(key)
	if envValue == "" {
		return 0
	}

	if sec, err := strconv.Atoi(envValue); err == nil {
		return time.Duration(sec) * time.Second
	}

	return 0
}
