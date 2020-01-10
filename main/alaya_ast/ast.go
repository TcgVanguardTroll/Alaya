package alaya_ast

import (
	token2 "Alaya/main/alaya_token"
	"strconv"
)

type (
	AST interface {
		AddValue(val string) int
		GetValue(ast AST) *AST
	}
)

type BinOp struct {
	Left  AST
	Op    token2.Token
	Right AST
}

type NumOp struct {
	Token token2.Token
	Value int
}

func NewBinOp(left AST, op token2.Token, right AST) *AST {
	var b AST
	b = BinOp{
		Left:  left,
		Op:    op,
		Right: right,
	}
	return &b
}

func NewNumOp(left token2.Token) *AST {
	var n AST
	val, _ := strconv.Atoi(left.TokenValue)
	n = NumOp{
		Token: left,
		Value: val,
	}
	return &n
}

func (b BinOp) AddValue(rightVal string) int {
	right, _ := strconv.Atoi(rightVal)
	return right
}
func (b BinOp) AddOp(op token2.Token) token2.Token {
	return op
}

func (b BinOp) GetValue(ast AST) *AST {
	return &(ast)
}

func (n NumOp) AddToken(leftVal token2.Token) token2.Token {
	return leftVal
}

func (n NumOp) AddValue(rightVal string) int {
	right, _ := strconv.Atoi(rightVal)
	return right
}
func (n NumOp) GetValue(ast AST) *AST {
	return &(ast)
}
