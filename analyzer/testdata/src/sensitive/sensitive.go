package sensitive

import (
	"log/slog"

	"go.uber.org/zap"
)

func f() {
	slog.Info("password is 123") // want "log message contains sensitive data: password"

	logger := zap.NewNop()
	logger.Info("password is 123") // want "log message contains sensitive data: password"

	// No diagnostics.
	slog.Info("user logged in")
	logger.Info("user logged in")
}
