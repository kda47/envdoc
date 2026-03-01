package main

import (
	"github.com/kda47/envdoc/linter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	analyzer := linter.NewAnlyzer(true)
	unitchecker.Main(analyzer)
}
