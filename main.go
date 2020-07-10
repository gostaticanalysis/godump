package main

import (
	"github.com/gostaticanalysis/godump/internal/passes/godump"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(godump.Analyzer) }

