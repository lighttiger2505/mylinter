package main

import (
	"mylinter"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(mylinter.Analyzer) }
