package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"LinterForLogs/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
