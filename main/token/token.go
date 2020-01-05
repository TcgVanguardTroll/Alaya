package token

import (
	"fmt"
)

type Type string

type Token struct {
	TokenType  Type
	TokenValue string
}

func (t Token) string() string {
	return fmt.Sprintf("Token(%s,%s)", t.TokenType, t.TokenValue)

}
func New(TokenType Type, TokenValue string) Token {
	var currentToken = Token{
		TokenType,
		TokenValue,
	}
	return currentToken

}

/*

 Token Types
*/

const (
	ILLEGAL = "ILLEGAL"
	INTEGER = "INTEGER"
	EOF     = "0"

	// Operators

	//	Keywords

	NAME    = "NAME"
	COMMAND = "COMMAND"

	//	Operators
	AS       = "="
	PLUS     = "+"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	//	Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACK    = "["
	RBRACK    = "]"
)
