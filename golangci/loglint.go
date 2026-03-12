package golangci

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"LinterForLogs/analyzer"
)

func init() {
	register.Plugin("loglint", New)
}

type plugin struct{}

// New is an entrypoint for golangci-lint's module plugin system.
// https://golangci-lint.run/plugins/module-plugins/
func New(settings any) (register.LinterPlugin, error) {
	_ = settings
	return plugin{}, nil
}

func (plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (plugin) GetLoadMode() string {
	// analyzer.isSupportedLogger uses type information.
	return register.LoadModeTypesInfo
}
