// Package logger предоставляет утилиты для логирования в приложении.
// Он включает в себя функционал для инициализации логера с различными уровнями логирования
// и middleware для логирования входящих HTTP-запросов.
package logger

import (
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// Log будет доступен всему коду как синглтон.
// По умолчанию установлен no-op-логер, который не выводит никаких сообщений.
var Log = zap.NewNop()

// ErrInvalidLogLevel определяет ошибку для недопустимого уровня логирования.
var ErrInvalidLogLevel = errors.New("invalid logging level")

// Initialize инициализирует синглтон логера с необходимым уровнем логирования.
func Initialize(level string) error {
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		// подробно опишем ошибку
		return fmt.Errorf(
			"invalid logging level '%s': %v. Valid levels are 'debug', 'info', 'warn', 'error', 'dpanic',"+
				" 'panic', and 'fatal'", ErrInvalidLogLevel, level)
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl
	// создаём логер на основе конфигурации
	zlog, err := cfg.Build()

	if err != nil {
		return fmt.Errorf("failed to build logger: %w", err)
	}
	// устанавливаем синглтон
	Log = zlog

	return nil
}

// RequestLogger — middleware-логер для входящих HTTP-запросов.
func RequestLogger(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log.Debug("got incoming HTTP request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)
		h(w, r)
	})
}
