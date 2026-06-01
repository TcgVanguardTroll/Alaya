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
// LookupIdent checks if the given identifier is a reserved keyword.
// Similar to: public static Type lookupIdent(String ident) in Java
//
// Parameters:
//   - ident: the identifier string to look up (e.g., "name", "cmnd", "hello")
//
// Returns:
//   - If the identifier is a keyword (exists in Keywords map), returns the keyword's Type
//   - If the identifier is NOT a keyword, returns IDENT (generic identifier type)
//
// Example usage:
//   LookupIdent("name")   -> returns NAME (keyword)
//   LookupIdent("cmnd")   -> returns COMMAND (keyword)
//   LookupIdent("Jordan") -> returns IDENT (not a keyword, just a variable name)
func LookupIdent(ident string) Type {
	// In Go, map lookup returns two values:
	// - tok: the value associated with the key (if found)
	// - ok: boolean indicating if the key exists in the map
	// This is like: if (Keywords.containsKey(ident)) in Java
	if tok, ok := Keywords[ident]; ok {
		// Key was found in the map, return the keyword type
		return tok
	}
	// Key was NOT found, so it's a regular identifier, not a keyword
	return IDENT
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
