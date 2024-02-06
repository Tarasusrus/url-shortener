package configs

import "flag"

type FlagConfig struct {
	address string
	baseURL string
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
	return c
}

func (c *FlagConfig) Address() string {
	return c.address
}

func (c *FlagConfig) BaseURL() string {
	return c.baseURL
}
