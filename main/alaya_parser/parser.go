package alaya_parser

import (
	"Alaya/main/alaya_token"
	"Alaya/main/alaya_tokenizer"
)

type (
	Parser struct {
		Tokenizer    *alaya_tokenizer.Tokenizer
		CurrentToken alaya_token.Token
		NextToken    alaya_token.Token
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

func (p *Parser) isMatch(tokenType alaya_token.Type) {
	if p.CurrentToken.TokenType == tokenType {
		p.CurrentToken = p.Tokenizer.GetNextToken()

	} else {
		panic("There was no Match !!")
	}
}

func (p *Parser) addChild(child AST) {
	p.Children = append(p.Children, child)

}

//	Return an INTEGER token value.
//	factor : INTEGER
//
//
func (p *Parser) factor() (BinOp, NumOp) {
	token := p.CurrentToken
	if token.TokenType == alaya_token.INTEGER {
		p.isMatch(alaya_token.INTEGER)
		currentNode := NewNumOp(token)
		return BinOp{}, currentNode
	} else {
		p.isMatch(alaya_token.LPAREN)
		currentNode := p.Expr()
		p.isMatch(alaya_token.RPAREN)
		return currentNode, NumOp{}
	}
}

func (p *Parser) term() BinOp {
	binNode, _ := p.factor()
	for containsOperator(p.CurrentToken.TokenType, []alaya_token.Type{alaya_token.ASTERISK, alaya_token.SLASH}) {
		tok := p.CurrentToken
		if tok.TokenType == alaya_token.ASTERISK {
			p.isMatch(alaya_token.ASTERISK)
		} else {
			if tok.TokenType == alaya_token.SLASH {
				p.isMatch(alaya_token.SLASH)
			}
		}
		currentNode = NewBinOp(binNode, tok, p.factor())
	}
	p.addChild(currentNode)
	return currentNode
}

//TODO(Implement the ability to skip whitespace in expression 3 3 + 3 should equal 36)

func (p *Parser) Expr() BinOp {
	var currentNode AST
	currentNode = p.term()
	for containsOperator(p.CurrentToken.TokenType, []alaya_token.Type{alaya_token.MINUS, alaya_token.PLUS}) {
		tok := p.CurrentToken
		if tok.TokenType == alaya_token.PLUS {
			p.isMatch(alaya_token.PLUS)
		} else {
			p.isMatch(alaya_token.MINUS)
		}
		currentNode = NewBinOp(currentNode, tok, p.term())
	}
	p.addChild(currentNode)
	return currentNode
}

func (p *Parser) parse() {

}
