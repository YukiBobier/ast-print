package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(in io.Reader, out io.Writer) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", in, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Fprint(out, fset, node, ast.NotNilFilter)
}
