package tokenizer

import (
	"Alaya/main/token"
	"errors"
	"unicode"
)

/*
The lexer, also called lexical analyzer or tokenizer,
is a program that breaks down the input source code into
a sequence of lexemes.
*/

//The job of the tokenizer is to read tokens one at a time from
//the input stream and pass the tokens to the parser.
type Tokenizer struct {
	text         string
	position     int
	nextPosition int
	currentToken byte
}

func New(input string) *Tokenizer {
	var newTokenizer = &Tokenizer{
		text: input,
	}
	newTokenizer.readChar()
	return newTokenizer
}

func (t Tokenizer) Error() error {
	var err = errors.New("error parsing text")
	return err
}

func (t Tokenizer) readChar() {
	if t.nextPosition >= len(t.text) {
		t.currentToken = 0
	} else {
		t.currentToken = t.text[t.nextPosition]
	}
	t.position = t.nextPosition
	t.nextPosition += 1
}

func (t Tokenizer) GetNextToken() token.Token {
	var text = t.text
	var currentByte = text[t.position]
	var currentToken token.Token
	if t.position > len(text)-1 {
		return token.New(token.EOF, 0)
	} else if unicode.IsDigit(rune(currentByte)) {
		currentToken = token.New(token.INTEGER, currentByte)
	} else if '+' == currentByte {
		currentToken = token.New(token.PLUS, currentByte)
	} else if '=' == currentByte {
		currentToken = token.New(token.AS, currentByte)
	} else if '(' == currentByte {
		currentToken = token.New(token.LPAREN, currentByte)
	} else if ')' == currentByte {
		currentToken = token.New(token.RPAREN, currentByte)
	} else if ';' == currentByte {
		currentToken = token.New(token.SEMICOLON, currentByte)
	} else if ',' == currentByte {
		currentToken = token.New(token.COMMA, currentByte)
	} else if '}' == currentByte {
		currentToken = token.New(token.RBRACE, currentByte)
	} else if '{' == currentByte {
		currentToken = token.New(token.LBRACE, currentByte)
	}
	t.readChar()
	return currentToken

}
