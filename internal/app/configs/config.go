package configs

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
)

type FlagConfig struct {
	Address string `env:"SERVER_ADDRESS"`
	BaseURL string `env:"BASE_URL"`
}

func NewFlagConfig() *FlagConfig {
	c := new(FlagConfig)
	flag.StringVar(&c.Address,
		"a",
		"localhost:8080",
		"Флаг -a отвечает за адрес запуска HTTP-сервера (значение может быть таким: localhost:8888)")
	flag.StringVar(&c.BaseURL,
		"b",
		"http://localhost:8080/",
		"Флаг -b отвечает за базовый адрес результирующего сокращённого URL "+
			"(значение: адрес сервера перед коротким URL, например http://localhost:8000/)")
	flag.Parse()
	// Если указана переменная окружения, она будет переопределить значение флага
	err := env.Parse(c)
	if err != nil {
		// обработка ошибок при разборе переменных окружения
		log.Fatal(err)
	}
	return c
}

func (c *FlagConfig) GetAddress() string {
	return c.Address
}

func (c *FlagConfig) GetBaseURL() string {
	return c.BaseURL
}
