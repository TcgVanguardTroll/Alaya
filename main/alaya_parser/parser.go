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
		symbolTable  map[string]int
	}
)

func NewParser(tokenizer *alaya_tokenizer.Tokenizer) *Parser {
	return &Parser{Tokenizer: tokenizer, symbolTable: make(map[string]int)}
}

func (p *Parser) isMatch(tokenType alaya_token.Type) {
	if p.CurrentToken.TokenType == tokenType {
		p.CurrentToken = p.Tokenizer.GetNextToken()

	} else {
		panic("There was no Match !!")
	}
}

func (p *Parser) factor() int {
	token := p.CurrentToken
	if token.TokenType == alaya_token.INTEGER {
		p.isMatch(alaya_token.INTEGER)
		val, _ := strconv.Atoi(token.Value)
		return val
	} else if token.TokenType == alaya_token.IDENTIFIER {
		p.isMatch(alaya_token.IDENTIFIER)
		val, _ := p.symbolTable[token.Value]
		return val
	} else {
		p.isMatch(alaya_token.LPAREN)
		result := p.expr()
		p.isMatch(alaya_token.RPAREN)
		return result
	}
}

func (p *Parser) term() int {
	result := p.factor()
	for p.CurrentToken.TokenType == alaya_token.ASTERISK || p.CurrentToken.TokenType == alaya_token.SLASH {
		token := p.CurrentToken
		if token.TokenType == alaya_token.ASTERISK {
			p.isMatch(alaya_token.ASTERISK)
			result *= p.factor()
		} else {
			p.isMatch(alaya_token.SLASH)
			result /= p.factor()
		}
	}
	return result
}

func (p *Parser) expr() int {
	result := p.term()
	for p.CurrentToken.TokenType == alaya_token.PLUS || p.CurrentToken.TokenType == alaya_token.MINUS {
		token := p.CurrentToken
		if token.TokenType == alaya_token.PLUS {
			p.isMatch(alaya_token.PLUS)
			result += p.term()
		} else {
			p.isMatch(alaya_token.MINUS)
			result -= p.term()
		}
	}
	return result
}