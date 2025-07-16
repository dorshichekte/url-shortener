package config

import (
	"time"
)

type HTTPAdapter struct {
	Server HTTPServer
	Router Router
}

type HTTPServer struct {
	Address           string
	ReadHeaderTimeout time.Duration `env:"SERVER_READ_HEADER_TIMEOUT"`
	ReadTimeout       time.Duration `env:"SERVER_READ_TIMEOUT"`
	ShutdownTimeout   time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT"`
	WriteTimeout      time.Duration `env:"SERVER_WRITE_TIMEOUT"`
}

type Router struct {
	Shutdown time.Duration `env:"ROUTER_SHUTDOWN"`
	Timeout  time.Duration `env:"ROUTER_TIMEOUT"`
}
