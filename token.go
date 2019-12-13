package main

type TokenType string

type Token struct {
	tokenType  TokenType
	tokenValue string
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
