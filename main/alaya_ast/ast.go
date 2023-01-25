package alaya_ast

import token2 "Alaya/main/alaya_token"

type Node interface {
	TokenType() string
}

type Condition interface {
	Node
	statementNode()
}

type AsStatement struct {
	Token token2.Token
	Name  *Identifier
	Value Expression
}

func (as *AsStatement) statementNode() {}

func (as *AsStatement) TokenValue() string { return as.Token.TokenValue }

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token2.Token
	Value string
}

func (i *Identifier) expressionNode()    {}
func (i *Identifier) TokenValue() string { return i.Token.TokenValue }

type Num struct {
    Token token2.Token
    Value int
}

func (n *Num) expressionNode() {}
func (n *Num) TokenType() string { return n.Token.TokenType }
func (n *Num) TokenValue() string { return n.Token.TokenValue }

type BinOp struct {
	Left  Expression
	Right Expression
	Op    token2.Token
}

func (b *BinOp) expressionNode() {}
func (b *BinOp) TokenType() string { return b.Op.TokenType }
func (b *BinOp) TokenValue() string { return b.Op.TokenValue }

type UnaryOp struct {
	Op   token2.Token
	Expr Expression
}

func (u *UnaryOp) expressionNode() {}
func (u *UnaryOp) TokenType() string { return u.Op.TokenType }
func (u *UnaryOp) TokenValue() string { return u.Op.TokenValue }

type Root struct {
	Statements []Statement
}

func (r *Root) TokenValue() string {
	if len(r.Statements) > 0 {
		return r.Statements[0].TokenType()
	} else {
		return ""
	}
}