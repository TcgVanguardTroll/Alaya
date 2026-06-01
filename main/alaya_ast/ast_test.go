package alaya_ast

import (
	"testing"

	token2 "github.com/TcgVanguardTroll/Alaya/main/alaya_token"
)

// TestNumNode tests the Num AST node creation and methods.
func TestNumNode(t *testing.T) {
	token := token2.Token{TokenType: token2.INTEGER, TokenValue: "42"}
	num := &Num{Token: token, Value: 42}

	if num.Value != 42 {
		t.Errorf("Expected value 42, got %d", num.Value)
	}

	if num.TokenType() != string(token2.INTEGER) {
		t.Errorf("Expected token type INTEGER, got %s", num.TokenType())
	}
}
