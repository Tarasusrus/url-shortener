package configs

import (
	"flag"
	"github.com/caarlos0/env"
)

type FlagConfig struct {
	Address  string `env:"SERVER_ADDRESS"`
	BaseURL  string `env:"BASE_URL"`
	LogLevel string `enc:"LOG_LEVEL"`
}

const (
	DefaultAddress         = "localhost:8080"
	DefaultBaseURL         = "http://" + DefaultAddress + "/"
	DefaultLogLevel        = "info"
	AddressFlagDescription = "Флаг -a отвечает за адрес запуска HTTP-сервера (значение может быть таким: localhost:8888)"
	BaseURLFlagDescription = "Флаг -b отвечает за базовый адрес результирующего сокращённого URL (значение: адрес сервера перед коротким URL, например http://localhost:8000/)"
	FlagLogLevel           = "Флаг -l отвечает за уровень логирования (допустимые значения: debug, info, warn, error, dpanic, panic, fatal)"
)

func NewFlagConfig() (*FlagConfig, error) {
	c := new(FlagConfig)
	flag.StringVar(&c.Address, "a", DefaultAddress, AddressFlagDescription)
	flag.StringVar(&c.BaseURL, "b", DefaultBaseURL, BaseURLFlagDescription)
	flag.StringVar(&c.LogLevel, "l", DefaultLogLevel, FlagLogLevel)
	flag.Parse()

	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *FlagConfig) GetAddress() string {
	return c.Address
}

func (c *FlagConfig) GetBaseURL() string {
	return c.BaseURL
}
