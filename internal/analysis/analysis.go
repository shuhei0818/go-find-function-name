package analysis

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

type Analyzer struct {
	// Input is filename.
	Input string
	// Line is line number.
	Line int
	// Ouput is destination of the result.
	Output io.Writer
}

func New(name string, line int) *Analyzer {
	return &Analyzer{
		Input:  name,
		Line:   line,
		Output: os.Stdout,
	}
}

// Do analyzes the source code and outputs the function name at the specified line number.
// The source would be written into Output.
func (a *Analyzer) Do() error {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, a.Input, nil, parser.Mode(0))
	if err != nil {
		return err
	}

	var name string
	ast.Inspect(f, func(n ast.Node) bool {
		if ident, ok := n.(*ast.FuncDecl); ok {
			start := fset.Position(ident.Pos()).Line
			end := fset.Position(ident.End()).Line
			if a.withinRange(start, end) {
				name = ident.Name.Name
			}
		}
		return true
	})

	if len(name) != 0 {
		fmt.Fprintf(a.Output, "%s\n", name)
	} else {
		return errors.New("not found function")
	}

	return nil
}

func (a *Analyzer) withinRange(s, e int) bool {
	return s <= a.Line && a.Line <= e
}
