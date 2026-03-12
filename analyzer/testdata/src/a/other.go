package a

type OtherLogger struct{}

func (OtherLogger) Info(msg string) {}

func other() {
	var l OtherLogger
	l.Info("Starting server") // no diagnostics: unsupported logger
}
