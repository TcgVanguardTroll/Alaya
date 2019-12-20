package tokenizer

import (
	"Alaya/main/token"
	"errors"
	"strconv"
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
	currentToken token.Token
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func New(input string) *Tokenizer {
	var newTokenizer = &Tokenizer{text: input}
	newTokenizer.readChar()
	return newTokenizer
}

func (t Tokenizer) Error() error {
	var err = errors.New("error parsing text")
	return err
}

func (t *Tokenizer) readChar() {
	if t.nextPosition >= len(t.text) {
		t.currentToken.TokenValue = string(0)
	} else {
		t.currentToken.TokenValue = string(t.text[t.nextPosition])
	}
	t.position = t.nextPosition
	t.nextPosition += 1
}

func (t *Tokenizer) GetNextToken() token.Token {
	var currentToken token.Token
	var currentByte = t.currentToken.TokenValue
	if currentByte == string(0) {
		currentToken = token.New(token.EOF, currentByte)
	} else if isNumeric(currentByte) {
		currentToken = token.New(token.INTEGER, currentByte)
	} else if "+" == currentByte {
		currentToken = token.New(token.PLUS, currentByte)
	} else if "=" == currentByte {
		currentToken = token.New(token.AS, currentByte)
	} else if "(" == currentByte {
		currentToken = token.New(token.LPAREN, currentByte)
	} else if ")" == currentByte {
		currentToken = token.New(token.RPAREN, currentByte)
	} else if ";" == currentByte {
		currentToken = token.New(token.SEMICOLON, currentByte)
	} else if "," == currentByte {
		currentToken = token.New(token.COMMA, currentByte)
	} else if "}" == currentByte {
		currentToken = token.New(token.RBRACE, currentByte)
	} else if "{" == currentByte {
		currentToken = token.New(token.LBRACE, currentByte)
	}
	t.readChar()
	return currentToken

}
func (t Tokenizer) consume(tokenToEat token.Type) {
	if t.currentToken.TokenType == tokenToEat {
		t.currentToken = t.GetNextToken()
	}
}

func (t Tokenizer) toExpression() int {
	t.currentToken = t.GetNextToken()

	var left = t.currentToken
	t.consume(token.INTEGER)

	_ = t.currentToken
	t.readChar()

	var right = t.currentToken
	t.readChar()

	var rightRes, _ = strconv.Atoi(right.TokenValue)
	var leftRes, _ = strconv.Atoi(left.TokenValue)

	return leftRes + rightRes
}
