package a

import (
	"log/slog"

	"go.uber.org/zap"
)

func slogCases() {
	slog.Info("Starting server")           // want "log message must start with lowercase"
	slog.Info("запуск сервера")            // want "log message must be in English"
	slog.Info("connection failed!!!")      // want "log message contains special characters"
	slog.Info("password is 123")           // want "log message contains sensitive data: password"
	slog.Info("password " + "is here")     // want "log message contains sensitive data: password"
	slog.Info("password " + getPassword()) // want "log message contains sensitive data: password"
}

func getPassword() string {
	return "123"
}

func zapCases() {
	logger := zap.NewNop()
	logger.Info("Starting server") // want "log message must start with lowercase"
}
