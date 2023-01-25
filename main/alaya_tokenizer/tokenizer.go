package alaya_tokenizer

import (
	"Alaya/main/alaya_token"
	"strings"
	"errors"


)

type Tokenizer struct {
	text             string
	position         int
	currentCharacter byte
}

func (t *Tokenizer) peek() byte {
	if t.position+1 >= len(t.text) {
		return 0
	}
	return t.text[t.position+1]
}

func (t *Tokenizer) backup(distance int) error {
	if (t.position - distance) < 0 {
		return errors.New("can not backup to negative position")
	}
	t.position -= distance
	return nil
}

func New(input string) *Tokenizer {
	var newTokenizer = &Tokenizer{text: input, position: 0}
	newTokenizer.currentCharacter = newTokenizer.text[newTokenizer.position]
	return newTokenizer
}

func (t *Tokenizer) skipComment() {
	for t.currentCharacter != '\n' && t.currentCharacter != 0 {
		t.Advance()
	}
}
func (t *Tokenizer) IgnoreWhitespace() {
	for t.currentCharacter == ' ' || t.currentCharacter == '\t' || t.currentCharacter == '\n' || t.currentCharacter == '\r' {
		t.Advance()
	}
}

func (t *Tokenizer) Advance() {
	t.position += 1
	if t.position >= len(t.text) {
		t.currentCharacter = 0
	} else {
		t.currentCharacter = t.text[t.position]
	}
}
func (t *Tokenizer) GetNextToken() alaya_token.Token {
	t.IgnoreWhitespace()
	if t.currentCharacter == '#' {
		t.skipComment()
		return t.GetNextToken()
	}
	
	var token alaya_token.Token
	
	switch t.currentCharacter {
	case '<':
		token = alaya_token.New(alaya_token.LT, t.currentCharacter)
	case '=':
		if t.peek() == '=' {
			token = alaya_token.New(alaya_token.ASCOMPARE, t.currentCharacter+string(t.peek()))
			t.Advance()
		} else {
			token = alaya_token.New(alaya_token.AS, t.currentCharacter)
		}
	case '>':
		token = alaya_token.New(alaya_token.GT, t.currentCharacter)
	case '!':
		if t.peek() == '=' {
			token = alaya_token.New(alaya_token.NotAs, t.currentCharacter+string(t.peek()))
			t.Advance()
		} else {
			token = alaya_token.New(alaya_token.BANG, t.currentCharacter)
		}
	case '*':
		token = alaya_token.New(alaya_token.ASTERISK, t.currentCharacter)
	case '/':
		token = alaya_token.New(alaya_token.SLASH, t.currentCharacter)
	case '[':
		token = alaya_token.New(alaya_token.LBRACK, t.currentCharacter)
	case ']':
		token = alaya_token.New(alaya_token.RBRACK, t.currentCharacter)
	case '(':
		token = alaya_token.New(alaya_token.LPAREN, t.currentCharacter)
	case ')':
		token = alaya_token.New(alaya_token.RPAREN, t.currentCharacter)
	case '{':
		token = alaya_token.New(alaya_token.LBRACE, t.currentCharacter)
	case '}':
		token = alaya_token.New(alaya_token.RBRACE, t.currentCharacter)
	case ',':
		token = alaya_token.New(alaya_token.COMMA, t.currentCharacter)
	case '.':
		token = alaya_token.New(alaya_token.DOT, t.currentCharacter)
	case '+':
		token = alaya_token.New(alaya_token.PLUS, t.currentCharacter)
	case '-':
		token = alaya_token.New(alaya_token.MINUS, t.currentCharacter)
	case ';':
		token = alaya_token.New(alaya_token.SEMICOLON, t.currentCharacter)
	case ':':
		token = alaya_token.New(alaya_token.COLON, t.currentCharacter)
	case 0:
		token = alaya_token.New(alaya_token.EOF, "")
	default:
		if isLetter(t.currentCharacter) {
		token = t.readIdentifier()
		return token
		} 
		else if isDigit(t.currentCharacter) {
		token = t.readNumber()
		return token
		} 
		else {
		token = alaya_token.New(alaya_token.ILLEGAL, string(t.currentCharacter))
		}
	}
	t.Advance()
	return token
}