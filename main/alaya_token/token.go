package alaya_token

import (
	"fmt"
)

type Type string

type Token struct {
	TokenType  Type
	TokenValue string
}

const (
	ILLEGAL = "ILLEGAL"
	INTEGER = "INTEGER"
	EOF     = "0"

	//	Keywords
	IDENT   = "IDENT"
	COMMAND = "COMMAND"
	NAME    = "NAME"
	TRUE    = "TRUE"
	FALSE   = "FALSE"
	IS      = "IS"
	ELSE    = "ELSE"
	RETURN  = "RETURN"
	//	Operators
	ASCOMPARE = "=="
	AS        = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	NotAs     = "!="
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	//	Punctuation
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACK    = "["
	RBRACK    = "]"
	DOT       = "."
)

var Keywords = map[string]Type{
	"cmnd":   COMMAND,
	"name":   NAME,
	"true":   TRUE,
	"false":  FALSE,
	"is":     IS,
	"else":   ELSE,
	"return": RETURN,
}

func (t Token) string() string {
	return fmt.Sprintf("Token(%s,%s)", t.TokenType, t.TokenValue)

}
func New(TokenType Type, TokenValue interface{}) Token {
	var value string
	switch v := TokenValue.(type) {
	case byte:
		value = string(v)
	case string:
		value = v
	default:
		value = fmt.Sprintf("%v", v)
	}

	return Token{
		TokenType:  TokenType,
		TokenValue: value,
	}
}
