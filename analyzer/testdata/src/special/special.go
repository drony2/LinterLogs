package special

import (
	"log/slog"

	"go.uber.org/zap"
)

func f() {
	slog.Info("connection failed!!!") // want "log message contains special characters"

	logger := zap.NewNop()
	logger.Info("connection failed!!!") // want "log message contains special characters"

	// No diagnostics.
	slog.Info("connection failed")
	logger.Info("connection failed")
}
