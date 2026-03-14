package english

import (
	"log/slog"

	"go.uber.org/zap"
)

func f() {
	// Starts with a lowercase Cyrillic letter, so only the English rule should trigger.
	slog.Info("\u043f\u0440\u0438\u0432\u0435\u0442") // want "log message must be in English"

	logger := zap.NewNop()
	logger.Info("\u043f\u0440\u0438\u0432\u0435\u0442") // want "log message must be in English"

	// No diagnostics.
	slog.Info("hello world")
	logger.Info("hello world")
}
