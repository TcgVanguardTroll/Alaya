package alaya_parser

import (
	"Alaya/main/alaya_token"
	"Alaya/main/alaya_tokenizer"
	"strconv"
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
func (p *Parser) factor() (int, error) {
	var res int
	token := p.CurrentToken
	if token.TokenType == alaya_token.INTEGER {
		p.isMatch(alaya_token.INTEGER)
		res, _ = strconv.Atoi(token.TokenValue)
		return res, nil
	} else {
		p.isMatch(alaya_token.LPAREN)
		var res = p.Expr()
		p.isMatch(alaya_token.RPAREN)
		return res, nil
	}
}

func (p *Parser) isMatch(tokenType alaya_token.Type) {
	if p.CurrentToken.TokenType == tokenType {
		p.CurrentToken = p.Tokenizer.GetNextToken()

	} else {
		panic("There was no Match !!")
	}
}

func (p *Parser) term() int {
	var res, _ = p.factor()

	for containsOperator(p.CurrentToken.TokenType, []alaya_token.Type{alaya_token.ASTERISK, alaya_token.SLASH}) {
		tok := p.CurrentToken
		if tok.TokenType == alaya_token.ASTERISK {
			p.isMatch(alaya_token.ASTERISK)
			match, err := p.factor()
			if err == nil {
				res = res * match
			}
		} else {
			tok := p.CurrentToken
			if tok.TokenType == alaya_token.SLASH {
				p.isMatch(alaya_token.SLASH)
				match, err := p.factor()
				if err == nil {
					res = res / match
				}
			}
		}
	}
	return res
}

//TODO(Implement the ability to skip whitespace in expression 3 3 + 3 should equal 36)

func (p *Parser) Expr() int {
	var result = p.term()
	for containsOperator(p.CurrentToken.TokenType, []alaya_token.Type{alaya_token.MINUS, alaya_token.PLUS}) {
		tok := p.CurrentToken
		if tok.TokenType == alaya_token.PLUS {
			p.isMatch(alaya_token.PLUS)
			match := p.term()
			result = result + match
		} else {
			p.isMatch(alaya_token.MINUS)
			match := p.term()
			result = result - match
		}
	}
	return result
}
