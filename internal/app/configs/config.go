package configs

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
)

type FlagConfig struct {
	address string `env:"SERVER_ADDRESS"`
	baseURL string `env:"BASE_URL"`
}

func NewFlagConfig() *FlagConfig {
	c := new(FlagConfig)
	flag.StringVar(&c.address,
		"a",
		"localhost:8080",
		"Флаг -a отвечает за адрес запуска HTTP-сервера (значение может быть таким: localhost:8888)")
	flag.StringVar(&c.baseURL,
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

func (c *FlagConfig) Address() string {
	return c.address
}

func (c *FlagConfig) BaseURL() string {
	return c.baseURL
}
