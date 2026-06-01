package alaya

/*
	Imports needed for Alaya.
*/
import (
	"github.com/TcgVanguardTroll/Alaya/main/alaya_ast"
	"github.com/TcgVanguardTroll/Alaya/main/alaya_parser"
	"github.com/TcgVanguardTroll/Alaya/main/alaya_token"
)

/* Alaya Struct that we can use. */
type Alaya struct {
	Parser alaya_parser.Parser
}

/**
* visitNum is a method of the Alaya struct that recursively traverses an Abstract Syntax Tree (AST) and evaluates the numerical value represented by the tree.
*
* The method takes in a node of the AST and uses a switch statement to check the type of the node. If the node is of type "alaya_ast.Num", it extracts the numerical value from the node using the "Value" field and returns it.
* If the node is of type "alaya_ast.BinOp", it checks the type of the operator using the "Op.Type" field. Depending on the operator type, it performs the corresponding arithmetic operation using the "Left" and "Right" fields of the node and the results of recursively calling the "visitNum" method on these fields.
* If the node is of type "alaya_ast.UnaryOp", it checks the type of the operator using the "Op.Type" field. Depending on the operator type, it performs the corresponding arithmetic operation using the "Expr" field of the node and the result of recursively calling the "visitNum" method on this field.
* If the node is of any other type, the function returns -1 and handle other types of nodes or raises an error.
*
* @param {alaya_ast.AST} ast - A node of the AST representing a numerical value or an arithmetic expression.
* @return {int} - The numerical value represented by the passed in AST node.
 */
func (a *Alaya) visitNum(ast alaya_ast.AST) int {
	switch n := ast.(type) {
	case *alaya_ast.Num:
		return n.Value
	case *alaya_ast.BinOp:
		switch n.Op.TokenType {
		case alaya_token.PLUS:
			return a.visitNum(n.Left) + a.visitNum(n.Right)
		case alaya_token.MINUS:
			return a.visitNum(n.Left) - a.visitNum(n.Right)
		case alaya_token.ASTERISK:
			return a.visitNum(n.Left) * a.visitNum(n.Right)
		case alaya_token.SLASH:
			return a.visitNum(n.Left) / a.visitNum(n.Right)
		}
	case *alaya_ast.UnaryOp:
		switch n.Op.TokenType {
		case alaya_token.PLUS:
			return a.visitNum(n.Expr)
		case alaya_token.MINUS:
			return -1 * a.visitNum(n.Expr)
		}
	}
	// Handle other types of nodes or unrecognized operators.
	return -1
}
