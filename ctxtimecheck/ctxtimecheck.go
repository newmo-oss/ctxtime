package ctxtimecheck

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/ssainspect"
	"golang.org/x/tools/go/analysis"
)

const doc = "ctxtimecheck finds calling time.Now instead of ctxtime.Now"

// Analyzer finds calling time.Now instead of ctxtime.Now.
var Analyzer = &analysis.Analyzer{
	Name: "ctxtimecheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		ssainspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	timenow, _ := analysisutil.ObjectOf(pass, "time", "Now").(*types.Func)
	checknow(pass, timenow)

	timedate := analysisutil.ObjectOf(pass, "time", "Date").(*types.Func)
	checkdate(pass, timedate)

	return nil, nil
}

func checknow(pass *analysis.Pass, timenow *types.Func) {
	if timenow == nil {
		// skip
		return
	}

	for s := range pass.ResultOf[ssainspect.Analyzer].(*ssainspect.Inspector).All() {
		if analysisutil.Called(s.Instr, nil, timenow) {
			pass.Reportf(s.Instr.Pos(), "do not use %s, use ctxtime.Now", timenow.FullName())
		}
	}
}

func checkdate(pass *analysis.Pass, timedate *types.Func) {
	if timedate == nil {
		// skip
		return
	}

	for s := range pass.ResultOf[ssainspect.Analyzer].(*ssainspect.Inspector).All() {
		if analysisutil.Called(s.Instr, nil, timedate) {
			pass.Reportf(s.Instr.Pos(), "do not use %s, use ctxtime.Now and its receiver methods to calculate date", timedate.FullName())
		}
	}
}
