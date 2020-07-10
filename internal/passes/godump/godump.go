package godump

import (
	"fmt"
	"os"

	"github.com/knsh14/astree"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var (
	flagMode string
)

func init() {
	Analyzer.Flags.StringVar(&flagMode, "mode", "ast", "ast,ssa")
}

// Analyzer dumps AST or SSA IR.
var Analyzer = &analysis.Analyzer{
	Name: "godump",
	Doc:  "godump dumps AST or SSA IR",
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	switch flagMode {
	case "ast":
		dumpAST(pass)
	case "ssa":
		dumpSSA(pass)
	}
	return nil, nil
}

func dumpAST(pass *analysis.Pass) {
	for i, f := range pass.Files {
		tf := pass.Fset.File(f.Pos())
		fmt.Println(tf.Name())
		astree.File(os.Stdout, pass.Fset, f)
		if i < len(pass.Files)-1 {
			fmt.Println()
		}
	}
}

func dumpSSA(pass *analysis.Pass) {
	srcFuncs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for i, sf := range srcFuncs {
		fmt.Println(sf)
		for _, b := range sf.Blocks {
			fmt.Println("\tBlock", b.Index)
			for _, instr := range b.Instrs {
				fmt.Printf("\t\t%[1]T\t%[1]v\n", instr)
			}
		}
		if i < len(srcFuncs)-1 {
			fmt.Println()
		}
	}
}
