package ctxtimecheck

import (
	"go/types"
	"iter"

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
	seq := pass.ResultOf[ssainspect.Analyzer].(iter.Seq[*ssainspect.Cursor])

	timenow, _ := analysisutil.ObjectOf(pass, "time", "Now").(*types.Func)
	if timenow == nil {
		// skip
		return nil, nil
	}

	for s := range seq {
		if analysisutil.Called(s.Instr, nil, timenow) {
			pass.Reportf(s.Instr.Pos(), "do not use %s, use ctxtime.Now", timenow.FullName())
		}
	}

	return nil, nil
}
