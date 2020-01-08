package alaya_tokenizer

import (
	_ "Alaya/main/alaya_token"
	token2 "Alaya/main/alaya_token"
	"testing"
)

func TestTokenizer_GetNextToken(t *testing.T) {
	input := `=+(){},;!===`
	var (
		tests = []struct {
			expectedType    token2.Type
			expectedLiteral string
		}{
			{token2.AS, "="}, {token2.PLUS, "+"},
			{token2.LPAREN, "("}, {token2.RPAREN, ")"},
			{token2.LBRACE, "{"}, {token2.RBRACE, "}"},
			{token2.COMMA, ","}, {token2.SEMICOLON, ";"},
			{token2.NotAs, "!="}, {token2.ASCOMPARE, "+="},
			{token2.EOF, ""},
		}
	)
	var l = New(input)
	for i, tt := range tests {
		tok := l.GetNextToken()
		if tok.TokenType != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.TokenType)
		}
	}
}
func TestTokenizer_GetNextToken_Whitespace(t *testing.T) {
	input := `name 9 = cmnd(a,b)`
	var (
		tests = []struct {
			expectedType    token2.Type
			expectedLiteral string
		}{
			{token2.NAME, "name"}, {token2.INTEGER, "9"},
			{token2.AS, "="}, {token2.COMMAND, "cmnd"},
			{token2.LPAREN, "("}, {token2.IDENT, "a"},
			{token2.COMMA, ","}, {token2.IDENT, "b"},
			{token2.RPAREN, ")"}, {token2.EOF, "0"},
		}
	)
	var l = New(input)
	for i, tt := range tests {
		tok := l.GetNextToken()
		if tok.TokenType != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.TokenType)
		}
	}
}

func TestTokenizer_Identifier(t *testing.T) {
	input := `Jordan9`
	tokeinzer := New(input)
	if tokeinzer.readIdentifier() != "Jordan" {
		t.Fatal()
	}
}
