package alaya_ast

import (
	token2 "Alaya/main/alaya_token"
	"strconv"
)

type AST interface {
	AddValue(val string) int
}

type BinOp struct {
	AST
	Left  AST
	Op    token2.Token
	Right AST
}

type NumOp struct {
	AST
	Token token2.Token
	Value int
}

func NewBinOp(left AST, op token2.Token, right AST) AST {
	b := BinOp{}
	return &BinOp{
		Left:  left,
		Op:    b.AddOp(op),
		Right: right,
	}
}

func (b *BinOp) AddValue(rightVal string) int {
	right, _ := strconv.Atoi(rightVal)
	return right
}

func (b *BinOp) AddOp(op token2.Token) token2.Token {
	return op
}

func NewNumOp(left token2.Token, right string) AST {
	n := NumOp{}
	return &NumOp{
		Token: n.AddToken(left),
		Value: n.AddValue(right),
	}
}

func (n *NumOp) AddToken(leftVal token2.Token) token2.Token {
	return leftVal
}

func (n *NumOp) AddValue(rightVal string) int {
	right, _ := strconv.Atoi(rightVal)
	return right
}
