package parser

import (
	"Alaya/main/alaya_token"
	"Alaya/main/tokenizer"
)

type Parser struct {
	Tokenizer tokenizer.Tokenizer
	//SymTable
	CurrentToken alaya_token.Token
	//ICode
}

func (p *Parser) isMatch(tokenType alaya_token.Type) bool {
	if p.CurrentToken.TokenType == tokenType {
		p.CurrentToken = p.Tokenizer.GetNextToken()
		return true

	} else {
		return false
	}
}
