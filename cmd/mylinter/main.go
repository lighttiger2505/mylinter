package main

import (
	"mylinter/internal/pkg/checkofficeid"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkofficeid.Analyzer) }
