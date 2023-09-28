package linters

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var TodoAnalyzer = &analysis.Analyzer{
	Name: "todo",
	Doc:  "finds todos without author",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if comment, ok := n.(*ast.Comment); ok {
				if strings.HasPrefix(comment.Text, "// TODO:") || strings.HasPrefix(comment.Text, "// TODO():") {
					pass.Report(analysis.Diagnostic{
						Pos:            comment.Pos(),
						End:            0,
						Category:       "todo",
						Message:        "TODO comment has no author",
						SuggestedFixes: nil,
					})
				}
			}

			return true
		})
	}

	return nil, nil
}
