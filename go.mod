module LinterForLogs

go 1.22.0

require (
	github.com/golangci/plugin-module-register v0.1.1
	go.uber.org/zap v1.0.0
	golang.org/x/tools v0.30.0
)

require (
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
)

replace go.uber.org/zap => ./stubs/go.uber.org/zap
