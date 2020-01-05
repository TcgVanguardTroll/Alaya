package tokenizer

import (
	"Alaya/main/token"
	"errors"
	"strconv"
	"strings"
)

/*
The lexer, also called lexical analyzer or tokenizer,
is a program that breaks down the input source code into
a sequence of lexemes.
*/

//The job of the tokenizer is to read tokens one at a time from
//the input stream and pass the tokens to the parser.
type Tokenizer struct {
	text             string
	position         int
	currentCharacter byte
}

func _(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func New(input string) *Tokenizer {
	var newTokenizer = &Tokenizer{text: input, position: 0}
	newTokenizer.currentCharacter = newTokenizer.text[newTokenizer.position]
	return newTokenizer
}

func (t Tokenizer) Error() error {
	var err = errors.New("error parsing text")
	return err
}

func (t *Tokenizer) GetNextToken() token.Token {
	tokenVal := t.currentCharacter
	for tokenVal != 0 {
		switch tokenVal {
		case ' ', '\n', '\t', '\r':
			t.IgnoreWhitespace()
			continue
		case '<':
			t.Advance()
			return token.New(token.LT, string(tokenVal))
		case '>':
			t.Advance()
			return token.New(token.GT, string(tokenVal))
		case '!':
			t.Advance()
			return token.New(token.BANG, string(tokenVal))
		case '=':
			t.Advance()
			return token.New(token.AS, string(tokenVal))
		case '*':
			t.Advance()
			return token.New(token.ASTERISK, string(tokenVal))
		case '/':
			t.Advance()
			return token.New(token.SLASH, string(tokenVal))
		case '[':
			t.Advance()
			return token.New(token.LBRACK, string(tokenVal))
		case ']':
			t.Advance()
			return token.New(token.RBRACK, string(tokenVal))

		case ',':
			t.Advance()
			return token.New(token.COMMA, string(tokenVal))
		case '+':
			t.Advance()
			return token.New(token.PLUS, string(tokenVal))
		case '(':
			t.Advance()
			return token.New(token.LPAREN, string(tokenVal))
		case ')':
			t.Advance()
			return token.New(token.RPAREN, string(tokenVal))
		case '{':
			t.Advance()
			return token.New(token.LBRACE, string(tokenVal))
		case '}':
			t.Advance()
			return token.New(token.RBRACE, string(tokenVal))
		case ';':
			t.Advance()
			return token.New(token.SEMICOLON, string(tokenVal))
		default:
			if t.isLetter() {
				t.Advance()
				return t.Char()
			} else if t.isDigit() {
				t.Advance()
				return token.New(token.INTEGER, string(tokenVal))
			} else {
				t.Advance()
				return token.New(token.ILLEGAL, string(tokenVal))
			}

		}

	}

	return token.New(token.EOF, "0")
}

func (t *Tokenizer) IgnoreWhitespace() {
	for t.currentCharacter == ' ' || t.currentCharacter == '\t' ||
		t.currentCharacter == '\n' || t.currentCharacter == '\r' {
		t.Advance()
	}
}

/*
Function that returns a token for each character.
*/

func (t *Tokenizer) Char() token.Token {
	var sb strings.Builder
	for t.isLetter() {
		sb.WriteByte(t.currentCharacter)
		t.Advance()
	}
	return token.New(token.NAME, sb.String())
}

func (t *Tokenizer) isLetter() bool {
	return t.currentCharacter >= 'a' && t.currentCharacter <= 'z' ||
		t.currentCharacter >= 'A' && t.currentCharacter <= 'Z'
}
func (t *Tokenizer) isDigit() bool {
	return '0' <= t.currentCharacter && t.currentCharacter <= '9'
}

/*
Reads through current tokens Text and returns
the Number numerical value as a String.
*/
//func (t *Tokenizer) readNumber() string {
//	// Determines whether the read number will be a
//	// Whole number or a float !
//	isDouble := false
//
//	for t.currentCharacter != '0' {
//		a:= func() bool{
//			if t.currentCharacter == '.'{
//				if isDouble {
//					return  false
//				}
//				isDouble = true
//				return true
//			}
//			return t.isDigit()
//		}
//	}
//
//	return token.New(token.INTEGER,a)
//}

//
//func (t *Tokenizer) readWhile() {
//	isDouble := false
//	number :=
//
//}

func (t *Tokenizer) Advance() {
	t.position += 1
	if t.position >= len(t.text) {
		t.currentCharacter = 0
	} else {
		t.currentCharacter = t.text[t.position]
	}
}
