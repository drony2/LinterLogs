package lowercase

import (
	"log/slog"

	"go.uber.org/zap"
)

func f() {
	slog.Info("Starting server") // want "log message must start with lowercase"

	logger := zap.NewNop()
	logger.Info("Starting server") // want "log message must start with lowercase"

	// No diagnostics.
	slog.Info("starting server")
	logger.Info("starting server")
}
