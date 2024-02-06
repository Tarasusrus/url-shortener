package configs

import "flag"

type FlagConfig struct {
	address string
	baseUrl string
}

func NewFlagConfig() *FlagConfig {
	c := new(FlagConfig)
	flag.StringVar(&c.address,
		"a",
		"localhost:8080",
		"Флаг -a отвечает за адрес запуска HTTP-сервера (значение может быть таким: localhost:8888)")
	flag.StringVar(&c.baseUrl,
		"b",
		"http://localhost:8080/qsd54gFg",
		"Флаг -b отвечает за базовый адрес результирующего сокращённого URL "+
			"(значение: адрес сервера перед коротким URL, например http://localhost:8000/)")
	flag.Parse()
	return c
}

func (c *FlagConfig) Address() string {
	return c.address
}

func (c *FlagConfig) BaseURL() string {
	return c.baseUrl
}
