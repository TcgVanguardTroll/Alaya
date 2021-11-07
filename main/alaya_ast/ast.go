package alaya_ast

import token2 "Alaya/main/alaya_token"

type Node interface {
	TokenType() string
}

type Statement interface {
	Node
	StatementNode()
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
	ExpressionNode()
}

type Identifier struct {
	Token token2.Token
	Value string
}

func (i *Identifier) expressionNode()    {}
func (i *Identifier) TokenValue() string { return i.Token.TokenValue }

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
