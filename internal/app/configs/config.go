package configs

import (
	"flag"
	"github.com/caarlos0/env"
)

// FlagConfig Структура содержит данные для настройки сервера
type FlagConfig struct {
	Address string `env:"SERVER_ADDRESS"`
	BaseURL string `env:"BASE_URL"`
}

// Начальная (по умолчанию) конфигурация сервера
const (
	DefaultAddress         = "localhost:8080"
	DefaultBaseURL         = "http://" + DefaultAddress + "/"
	AddressFlagDescription = "Флаг -a отвечает за адрес запуска HTTP-сервера (значение может быть таким: localhost:8888)"
	BaseURLFlagDescription = "Флаг -b отвечает за базовый адрес результирующего сокращённого URL (значение: адрес сервера перед коротким URL, например http://localhost:8000/)"
)

// NewFlagConfig создает новый объект FlagConfig, основываясь на данных из флагов командной строки и переменных окружения.
func NewFlagConfig() (*FlagConfig, error) {
	c := new(FlagConfig)
	flag.StringVar(&c.Address, "a", DefaultAddress, AddressFlagDescription)
	flag.StringVar(&c.BaseURL, "b", DefaultBaseURL, BaseURLFlagDescription)
	flag.Parse()

	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}

// GetAddress возвращает адрес сервера из конфигурации
func (c *FlagConfig) GetAddress() string {
	return c.Address
}

// GetBaseURL возвращает базовый URL из конфигурации
func (c *FlagConfig) GetBaseURL() string {
	return c.BaseURL
}
