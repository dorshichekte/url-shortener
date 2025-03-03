package config

import (
	"errors"
	"flag"
	"strconv"
	"strings"
	"sync"
)

const DefaultAddress = "http://localhost:8080"

type ServerAddress struct {
	Host string
	Port int
}

type Config struct {
	ServerAddress ServerAddress
	BaseURL       string
}

var (
	Instance *Config
	once     sync.Once
)

func (sa *ServerAddress) String() string {
	return sa.Host + ":" + strconv.Itoa(sa.Port)
}

func (sa *ServerAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("need address in a form host:port")
	}

	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}

	sa.Host = hp[0]
	sa.Port = port
	return nil
}

func Create() {
	address := flag.String("a", DefaultAddress, "server address")
	baseURL := flag.String("b", DefaultAddress, "base host URL")

	flag.Parse()

	Instance = &Config{
		BaseURL: *baseURL,
	}

	err := Instance.ServerAddress.Set(*address)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	if Instance == nil {
		once.Do(func() {
			Create()
		})
	}
	return Instance
}
