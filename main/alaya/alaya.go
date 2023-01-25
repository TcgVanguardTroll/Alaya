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
    switch n := ast.(type) {
    case *alaya_ast.Num:
        return n.Value
    case *alaya_ast.BinOp:
        switch n.Op.Type {
        case alaya_parser.PLUS:
            return a.visitNum(n.Left) + a.visitNum(n.Right)
        case alaya_parser.MINUS:
            return a.visitNum(n.Left) - a.visitNum(n.Right)
        case alaya_parser.MUL:
            return a.visitNum(n.Left) * a.visitNum(n.Right)
        case alaya_parser.DIV:
            return a.visitNum(n.Left) / a.visitNum(n.Right)
        }
    case *alaya_ast.UnaryOp:
        switch n.Op.Type {
        case alaya_parser.PLUS:
            return a.visitNum(n.Expr)
        case alaya_parser.MINUS:
            return -1 * a.visitNum(n.Expr)
        }
    default:
        // Handle other types of nodes or raise an error.
        return -1
    }
}
