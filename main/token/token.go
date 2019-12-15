package token

import "fmt"

type Type string

type Token struct {
	TokenType  Type
	TokenValue string
	linenumber int
	position   int
}

func (t Token) toString() string {
	return fmt.Sprintf("Token(%s,%s)", t.TokenType, t.TokenValue)

}

/*

 Token Types
*/

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//	Keywords

	CAST    = "CAST"
	COMMAND = "COMMAND"

	//	Operators
	AS   = "="
	PLUS = "+"

	//	Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)
