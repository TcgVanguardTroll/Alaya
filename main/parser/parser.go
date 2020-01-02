package parser

import (
	"Alaya/main/token"
	"Alaya/main/tokenizer"
)

type Parser struct {
	Tokenizer    tokenizer.Tokenizer
	CurrentToken token.Token
}

func (p *Parser) match(tokenType token.Type) (string, error) {
	if p.CurrentToken.TokenType != tokenType {
		p.CurrentToken = p.Tokenizer.GetNextToken()
		return "Was A Match !", nil

	} else {
		return "Wrong Type of Token !", p.Tokenizer.Error()
	}
}
