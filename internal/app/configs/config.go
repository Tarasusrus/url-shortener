// Package configs предоставляет функционал для конфигурации приложения через флаги командной строки
// и переменные окружения.
// Основной структурой пакета является FlagConfig, которая определяет параметры конфигурации,
// такие как адрес сервера, базовый URL для сокращённых ссылок и уровень логирования.
//
// Флаги командной строки и переменные окружения предоставляют гибкие способы для настройки поведения приложения
// без необходимости изменения кода, что особенно полезно в различных средах развертывания.
// Пакет использует стороннюю библиотеку github.com/caarlos0/env для упрощения работы с переменными окружения.
//
// Пример использования включает создание нового экземпляра FlagConfig с использованием функции NewFlagConfig,
// которая считывает значения флагов и переменных окружения, предоставляя доступ к настроенным значениям
// через методы GetAddress и GetBaseURL.
package configs

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env"
)

// FlagConfig структуру предполагается заполнять из переменных окружения.
type FlagConfig struct {
	Address  string `env:"SERVER_ADDRESS"`
	BaseURL  string `env:"BASE_URL"`
	LogLevel string `env:"LOG_LEVEL"`
	FilePath string `env:"FILE_STORAGE_PATH"`
}

const (
	DefaultAddress         = "localhost:8080"
	DefaultBaseURL         = "http://" + DefaultAddress + "/"
	DefaultLogLevel        = "info"
	DefaultFilePath        = "short-url-db.json"
	AddressFlagDescription = "Флаг -a отвечает за адрес запуска HTTP-сервера " +
		"(значение может быть таким: localhost:8888)"
	BaseURLFlagDescription = "Флаг -b отвечает за базовый адрес результирующего сокращённого URL " +
		"(значение: адрес сервера перед коротким URL, например http://localhost:8000/)"
	FlagLogLevel = "Флаг -l отвечает за уровень логирования " +
		"(допустимые значения: debug, info, warn, error, dpanic, panic, fatal)"
	FilePathFlagDescription = "file storage path"
)

// NewFlagConfig парсинг флагов.
func NewFlagConfig() (*FlagConfig, error) {
	conf := new(FlagConfig)
	flag.StringVar(&conf.Address, "a", DefaultAddress, AddressFlagDescription)
	flag.StringVar(&conf.BaseURL, "b", DefaultBaseURL, BaseURLFlagDescription)
	flag.StringVar(&conf.LogLevel, "l", DefaultLogLevel, FlagLogLevel)
	flag.StringVar(&conf.FilePath, "f", DefaultFilePath, FilePathFlagDescription)
	flag.Parse()

	if err := env.Parse(conf); err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}

	return conf, nil
}

// GetAddress возвращает Address.
func (c *FlagConfig) GetAddress() string {
	return c.Address
}

// GetBaseURL возвращает BaseURL.
func (c *FlagConfig) GetBaseURL() string {
	return c.BaseURL
}
func (c *FlagConfig) GetFilePath() string {
	return c.FilePath
}
