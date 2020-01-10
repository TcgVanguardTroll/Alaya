package alaya_parser

import (
	. "Alaya/main/alaya_ast"
	"Alaya/main/alaya_token"
	"Alaya/main/alaya_tokenizer"
)

type (
	Parser struct {
		Tokenizer    *alaya_tokenizer.Tokenizer
		CurrentToken alaya_token.Token
	}
)

var (
	_ = map[string]int{
		"(": 0,
		")": 0,
		"*": 1,
		"/": 1,
		"+": 2,
		"-": 2,
	}
)

func NewParser(tokenizer *alaya_tokenizer.Tokenizer) *Parser {
	return &Parser{Tokenizer: tokenizer}
}

func containsOperator(inputType alaya_token.Type, typeArr []alaya_token.Type) bool {
	for _, b := range typeArr {
		if b == inputType {
			return true
		}
	}
	return false
}

//	Return an INTEGER token value.
//	factor : INTEGER
//
//
func (p *Parser) factor() AST {
	token := p.CurrentToken
	if token.TokenType == alaya_token.INTEGER {
		p.isMatch(alaya_token.INTEGER)
		node := NewNumOp(token)
		return *node
	} else {
		p.isMatch(alaya_token.LPAREN)
		node := p.Expr()
		p.isMatch(alaya_token.RPAREN)
		return node
	}
}

func (p *Parser) isMatch(tokenType alaya_token.Type) {
	if p.CurrentToken.TokenType == tokenType {
		p.CurrentToken = p.Tokenizer.GetNextToken()

	} else {
		panic("There was no Match !!")
	}
}

func (p *Parser) term() AST {
	var node = p.factor()
	for containsOperator(p.CurrentToken.TokenType, []alaya_token.Type{alaya_token.ASTERISK, alaya_token.SLASH}) {
		tok := p.CurrentToken
		if tok.TokenType == alaya_token.ASTERISK {
			p.isMatch(alaya_token.ASTERISK)
		} else {
			tok := p.CurrentToken
			if tok.TokenType == alaya_token.SLASH {
				p.isMatch(alaya_token.SLASH)
			}
		}
		node = *NewBinOp(node, tok, p.factor())
	}
	return node
}

//TODO(Implement the ability to skip whitespace in expression 3 3 + 3 should equal 36)

func (p *Parser) Expr() AST {
	var node AST
	node = p.term()
	for containsOperator(p.CurrentToken.TokenType, []alaya_token.Type{alaya_token.MINUS, alaya_token.PLUS}) {
		tok := p.CurrentToken
		if tok.TokenType == alaya_token.PLUS {
			p.isMatch(alaya_token.PLUS)
		} else {
			p.isMatch(alaya_token.MINUS)
		}
		node = *NewBinOp(node, tok, p.term())
	}
	return node
}

func (p *Parser) parse() {

}
