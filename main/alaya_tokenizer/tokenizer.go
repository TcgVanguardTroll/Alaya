package alaya_tokenizer

import (
	"Alaya/main/alaya_token"
	"errors"
	"strings"
)

/*
	The lexer, also called lexical analyzer or alaya_tokenizer,
	is a program that breaks down the input source code into
	a sequence of lexemes.
*/

/*
	The job of the alaya_tokenizer is to read tokens one at a time from
	the input stream and pass the tokens to the alaya_parser.
*/
type Tokenizer struct {
	text             string
	position         int
	currentCharacter byte
}

/*
	Checks the next Character within the Tokenizer without
	incrementing the current position. If at the last position
	returns the EOF value of 0.
*/

func (t *Tokenizer) peek() byte {
	if t.position+1 >= len(t.text) {
		return 0
	}
	return t.text[t.position+1]
}

/*
	Decrements the tokenizer's position by the input distance
	lowest position possible is 0.
*/

func (t *Tokenizer) backup(distance int) {

	if (t.position - distance) < 0 {
		t.position = 0
	}
	t.position -= distance
}

/*
	Returns a New instance of a Tokenizer.
	Input takes the string that will be initialized
	and the position will be set to 0 alongside it's
	initial character.
*/

func New(input string) *Tokenizer {
	var newTokenizer = &Tokenizer{text: input, position: 0}
	newTokenizer.currentCharacter = newTokenizer.text[newTokenizer.position]
	return newTokenizer
}

func (t Tokenizer) Error() error {
	var err = errors.New("error parsing text")
	return err
}

func (t *Tokenizer) GetNextToken() alaya_token.Token {
	tokenVal := t.currentCharacter
	for tokenVal != 0 {
		t.IgnoreWhitespace()
		tokenVal = t.currentCharacter
		switch tokenVal {
		case '<':
			t.Advance()
			return alaya_token.New(alaya_token.LT, tokenVal)
		case '=':
			if t.peek() == '=' {
				c := tokenVal
				t.Advance()
				t.Advance()
				return alaya_token.New(alaya_token.ASCOMPARE, (c)+(tokenVal))
			}
			t.Advance()
			return alaya_token.New(alaya_token.AS, tokenVal)
		case '>':
			t.Advance()
			return alaya_token.New(alaya_token.GT, tokenVal)
		case '!':
			if t.peek() == '=' {
				c := t.currentCharacter
				t.Advance()
				v := t.currentCharacter
				t.Advance()
				return alaya_token.Token{
					TokenType:  alaya_token.NotAs,
					TokenValue: string(c) + string(v),
				}
			}
			t.Advance()
			return alaya_token.New(alaya_token.BANG, tokenVal)
		case '*':
			t.Advance()
			return alaya_token.New(alaya_token.ASTERISK, tokenVal)
		case '/':
			t.Advance()
			return alaya_token.New(alaya_token.SLASH, tokenVal)
		case '[':
			t.Advance()
			return alaya_token.New(alaya_token.LBRACK, tokenVal)
		case ']':
			t.Advance()
			return alaya_token.New(alaya_token.RBRACK, tokenVal)
		case ',':
			t.Advance()
			return alaya_token.New(alaya_token.COMMA, tokenVal)
		case '+':
			t.Advance()
			return alaya_token.New(alaya_token.PLUS, tokenVal)
		case '-':
			t.Advance()
			return alaya_token.New(alaya_token.MINUS, tokenVal)
		case '(':
			t.Advance()
			return alaya_token.New(alaya_token.LPAREN, tokenVal)
		case ')':
			t.Advance()
			return alaya_token.New(alaya_token.RPAREN, tokenVal)
		case '{':
			t.Advance()
			return alaya_token.New(alaya_token.LBRACE, tokenVal)
		case '}':
			t.Advance()
			return alaya_token.New(alaya_token.RBRACE, tokenVal)
		case ';':
			t.Advance()
			return alaya_token.New(alaya_token.SEMICOLON, tokenVal)
		default:
			if t.isLetter() {
				value := t.readIdentifier()
				tokType := t.lookupIdentifier(value)
				return alaya_token.Token{
					TokenType:  tokType,
					TokenValue: value,
				}
			} else if t.isDigit() {
				return alaya_token.Token{
					TokenType:  alaya_token.INTEGER,
					TokenValue: t.readNumber(),
				}
			} else {
				t.Advance()
				return alaya_token.New(alaya_token.ILLEGAL, tokenVal)
			}

		}

	}

	return alaya_token.New(alaya_token.EOF, 0)
}

/*
	Ignores any whitespace from within the tokenizer's
	text property by proceeding as usual by going to the
	next position.
*/
func (t *Tokenizer) IgnoreWhitespace() {
	for t.currentCharacter == ' ' || t.currentCharacter == '\t' || t.currentCharacter == '\n' || t.currentCharacter == '\r' {
		t.Advance()
	}
}

/*
	Function that extracts the Tokenizer's current token
	and as long as it as well as the following characters
	are letters returns a string representing it.
*/
func (t *Tokenizer) readIdentifier() string {
	var sb strings.Builder
	for t.isLetter() {
		sb.WriteByte(t.currentCharacter)
		t.Advance()
	}
	return sb.String()
}

/*
	Builds a numerical representation of a String
	containing digits "0-9" and then returns that
	same representation while simultaneously advancing
	the tokenizer's position.
*/
func (t *Tokenizer) readNumber() string {
	var sb strings.Builder
	for t.isDigit() {
		sb.WriteByte(t.currentCharacter)
		t.Advance()
	}
	return sb.String()
}

/*
	Checks if the alaya_tokenizer's current character is a Letter and
	if so returns True else False.
*/

func (t *Tokenizer) isLetter() bool {
	return t.currentCharacter >= 'a' && t.currentCharacter <= 'z' ||
		t.currentCharacter >= 'A' && t.currentCharacter <= 'Z'
}
func (t *Tokenizer) isDigit() bool {
	return '0' <= t.currentCharacter && t.currentCharacter <= '9'
}
func (t *Tokenizer) isPunc() bool {
	return strings.Contains(",;(){}[]", string(t.currentCharacter))
}
func (t *Tokenizer) isKeyword() bool {
	if _, ok := alaya_token.Keywords[string(t.currentCharacter)]; ok {
		return true
	}
	return false
}
func (t *Tokenizer) lookupIdentifier(ident string) alaya_token.Type {
	if tok, ok := alaya_token.Keywords[ident]; ok {
		return tok
	}
	return alaya_token.IDENT
}

/*	Increments the alaya_tokenizer's current position by one
	and checks if the position is at the end, if so then
	the alaya_tokenizer's current character is EOF Value
	Else The current token character is the character at
	the advanced position.
*/
func (t *Tokenizer) Advance() {
	t.position += 1
	if t.position >= len(t.text) {
		t.currentCharacter = 0
	} else {
		t.currentCharacter = t.text[t.position]
	}
}
