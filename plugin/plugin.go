package main

import (
	"errors"

	"golang.org/x/tools/go/analysis"

	"LinterForLogs/analyzer"
)

// New is an entrypoint for golangci-lint's Go plugin system.
// https://golangci-lint.run/plugins/go-plugin/
func New(conf any) ([]*analysis.Analyzer, error) {
	_ = conf
	if analyzer.Analyzer == nil {
		return nil, errors.New("loglint: analyzer is nil")
	}
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}
