package ctxtimecheck

import (
	"go/types"
	"slices"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "ctxtimecheck finds calling time.Now instead of ctxtime.Now"

// Analyzer finds calling time.Now instead of ctxtime.Now.
var Analyzer = &analysis.Analyzer{
	Name: "ctxtimecheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	timenow, _ := analysisutil.ObjectOf(pass, "time", "Now").(*types.Func)
	timedate, _ := analysisutil.ObjectOf(pass, "time", "Date").(*types.Func)

	for _, fn := range allFuncs(pass) {
		for _, b := range fn.Blocks {
			for _, instr := range b.Instrs {
				if timenow != nil && analysisutil.Called(instr, nil, timenow) {
					pass.Reportf(instr.Pos(), "do not use %s, use ctxtime.Now", timenow.FullName())
				}
				if timedate != nil && analysisutil.Called(instr, nil, timedate) {
					pass.Reportf(instr.Pos(), "do not use %s, use ctxtime.Now and its receiver methods to calculate date", timedate.FullName())
				}
			}
		}
	}

	return nil, nil
}

// allFuncs returns source functions plus the synthetic package initializer.
// Package-level var initializers live there but are omitted from SrcFuncs.
func allFuncs(pass *analysis.Pass) []*ssa.Function {
	ssainfo := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	funcs := slices.Clone(ssainfo.SrcFuncs)
	if init := ssainfo.Pkg.Func("init"); init != nil {
		funcs = appendAnons(funcs, init)
	}
	return funcs
}

func appendAnons(funcs []*ssa.Function, fn *ssa.Function) []*ssa.Function {
	funcs = append(funcs, fn)
	for _, anon := range fn.AnonFuncs {
		funcs = appendAnons(funcs, anon)
	}
	return funcs
}
