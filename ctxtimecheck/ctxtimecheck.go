package ctxtimecheck

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/ssainspect"
	"golang.org/x/tools/go/analysis"
)

const doc = "ctxtimecheck finds calling time.Now instead of ctxtime.Now"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "ctxtimecheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		ssainspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	in := pass.ResultOf[ssainspect.Analyzer].(*ssainspect.Inspector)

	timenow, _ := analysisutil.ObjectOf(pass, "time", "Now").(*types.Func)
	if timenow == nil {
		// skip
		return nil, nil
	}

	for in.Next() {
		c := in.Cursor()
		if analysisutil.Called(c.Instr, nil, timenow) {
			pass.Reportf(c.Instr.Pos(), "do not use %s, use ctxtime.Now", timenow.FullName())
		}
	}

	return nil, nil
}
