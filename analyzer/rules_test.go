package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"LinterForLogs/analyzer"
)

func TestAnalyzerRuleLowercase(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.Analyzer, "lowercase")
}

func TestAnalyzerRuleEnglish(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.Analyzer, "english")
}

func TestAnalyzerRuleSpecialCharacters(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.Analyzer, "special")
}

func TestAnalyzerRuleSensitiveData(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.Analyzer, "sensitive")
}
