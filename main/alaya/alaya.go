package alaya

/*
	Imports needed for Alaya.
*/
import (
	"Alaya/main/alaya_ast"
	"Alaya/main/alaya_parser"
)

/* Alaya Struct that we can use. */
type Alaya struct {
	Parser alaya_parser.Parser
}

func (a *Alaya) visitNum(ast alaya_ast.AST) int {
	return
}
