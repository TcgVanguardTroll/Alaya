package tokenizer

import (
	"errors"
)

/*
The lexer, also called lexical analyzer or tokenizer,
is a program that breaks down the input source code into
a sequence of lexemes.
*/

type Tokenizer struct {
	text         string
	position     int
	currentToken string
}

func (t Tokenizer) New(input string) *Tokenizer {
	var newTokenizer = &Tokenizer{
		text: input,
	}
	return newTokenizer
}

func (t Tokenizer) Error() error {
	var err = errors.New("error parsing text")
	return err
}

//func (t Tokenizer) GetNextToken() token.Token {
//
//}
