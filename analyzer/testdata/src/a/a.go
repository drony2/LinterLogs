package a

import (
	"log/slog"

	"go.uber.org/zap"
)

func slogCases() {
	slog.Info("Starting server")                                                                 // want "log message must start with lowercase"
	slog.Info("\u0437\u0430\u043f\u0443\u0441\u043a \u0441\u0435\u0440\u0432\u0435\u0440\u0430") // want "log message must be in English"
	slog.Info("connection failed!!!")                                                            // want "log message contains special characters"
	slog.Info("password is 123")                                                                 // want "log message contains sensitive data: password"
	slog.Info("password " + "is here")                                                           // want "log message contains sensitive data: password"
	slog.Info("password " + getPassword())                                                       // want "log message contains sensitive data: password"
}

func getPassword() string {
	return "123"
}

func zapCases() {
	logger := zap.NewNop()
	logger.Info("Starting server") // want "log message must start with lowercase"
}
