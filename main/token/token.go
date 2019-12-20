package token

import (
	"fmt"
)

type Type string

type Token struct {
	TokenType  Type
	TokenValue string
	Position   int
}

func (t Token) string() string {
	return fmt.Sprintf("Token(%s,%s)", t.TokenType, t.TokenValue)

}
func New(TokenType Type, TokenValue string) Token {
	var currentToken = Token{
		TokenType,
		TokenValue,
		0,
	}
	return currentToken

}

/*

 Token Types
*/

const (
	ILLEGAL = "ILLEGAL"
	INTEGER = "INTEGER"
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
