package main

import (
	"github.com/newmo-oss/ctxtime/ctxtimecheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ctxtimecheck.Analyzer) }
