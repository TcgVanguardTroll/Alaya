package alaya_token

import (
	"fmt"
)

// Type represents a token type in the Alaya language.
type Type string

// Token represents a lexical token with its type and value.
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

// LookupIdent checks if the given identifier is a reserved keyword.
// If the identifier is a keyword, it returns the keyword's Type.
// Otherwise, it returns IDENT for user-defined identifiers.
func LookupIdent(ident string) Type {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// New creates a new Token with the specified type and value.
// The value can be a byte or string; other types are converted using fmt.Sprintf.
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
