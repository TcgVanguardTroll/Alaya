package tokenizer

import (
	_ "Alaya/main/token"
	token2 "Alaya/main/token"
	"testing"
)

func TestTokenizer_GetNextToken(t *testing.T) {
	input := `=+(){},;`
	var tests = []struct {
		expectedType    token2.Type
		expectedLiteral string
	}{
		{token2.AS, "="}, {token2.PLUS, "+"},
		{token2.LPAREN, "("}, {token2.RPAREN, ")"},
		{token2.LBRACE, "{"}, {token2.RBRACE, "}"},
		{token2.COMMA, ","}, {token2.SEMICOLON, ";"},
		{token2.EOF, ""},
	}
	var l = New(input)
	for i, tt := range tests {
		tok := l.GetNextToken()
		if tok.TokenType != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.TokenType)
		}
	}
}

func TestTokenizer_Char(t *testing.T) {
	input := `Jordan`
	tokeinzer := New(input)
	if tokeinzer.Char() != token2.New(token2.NAME, "Jordan") {
		t.Fatal()
	}
}
