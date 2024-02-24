package logger

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

// Log будет доступен всему коду как синглтон.

// По умолчанию установлен no-op-логер, который не выводит никаких сообщений.
var Log = zap.NewNop()

// Initialize инициализирует синглтон логера с необходимым уровнем логирования.
func Initialize(level string) error {
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		// подробно опишем ошибку
		return fmt.Errorf("invalid logging level '%s': %v. Valid levels are 'debug', 'info', 'warn', 'error', 'dpanic', 'panic', and 'fatal'", level, err)
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl
	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		return err
	}
	// устанавливаем синглтон
	Log = zl
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
