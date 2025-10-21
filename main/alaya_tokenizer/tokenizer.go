package alaya_tokenizer

import (
	"errors"
	"strings"

	"github.com/TcgVanguardTroll/Alaya/main/alaya_token"
)

type Tokenizer struct {
	text             string
	position         int
	currentCharacter byte
}

// isLetter checks whether a character is a letter (a-z, A-Z) or underscore.
// Similar to: Character.isLetter(ch) || ch == '_' in Java
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		ch == '_'
}

// isDigit checks whether a character is a digit (0-9).
// Similar to: Character.isDigit(ch) in Java
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// readIdentifier reads a complete identifier or keyword from the source code.
// Starts at the current position and reads all consecutive letter characters.
// Returns a Token with the appropriate type (either a keyword or IDENT).
//
// Example: "name = 5" starting at 'n' will read "name" and return a NAME token
func (t *Tokenizer) readIdentifier() alaya_token.Token {
	// String builder to accumulate the identifier characters
	var tokenText strings.Builder

	// Keep reading while current character is a letter
	// Similar to: while (position < text.length() && isLetter(currentCharacter))
	for t.position < len(t.text) && isLetter(t.currentCharacter) {
		tokenText.WriteByte(t.currentCharacter) // Add CURRENT character first
		t.Advance()                              // Then move to next
	}

	// Look up if the identifier is a keyword or just a variable name
	tokenType := alaya_token.LookupIdent(tokenText.String())

	return alaya_token.Token{TokenType: tokenType, TokenValue: tokenText.String()}
}

// readNumber reads a complete number from the source code.
// Starts at the current position and reads all consecutive digit characters.
// Returns a Token with type INTEGER.
//
// Example: "123 + 456" starting at '1' will read "123" and return an INTEGER token
func (t *Tokenizer) readNumber() alaya_token.Token {
	// String builder to accumulate the digit characters
	var tokenText strings.Builder

	// Keep reading while current character is a digit
	// Similar to: while (position < text.length() && isDigit(currentCharacter))
	for t.position < len(t.text) && isDigit(t.currentCharacter) {
		tokenText.WriteByte(t.currentCharacter) // Add CURRENT character first
		t.Advance()                              // Then move to next
	}

	return alaya_token.Token{TokenType: alaya_token.INTEGER, TokenValue: tokenText.String()}
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
			token = alaya_token.New(alaya_token.ASCOMPARE, string(t.currentCharacter)+string(t.peek()))
			t.Advance()
		} else {
			token = alaya_token.New(alaya_token.AS, t.currentCharacter)
		}
	case '>':
		token = alaya_token.New(alaya_token.GT, t.currentCharacter)
	case '!':
		if t.peek() == '=' {
			token = alaya_token.New(alaya_token.NotAs, string(t.currentCharacter)+string(t.peek()))
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
		} else if isDigit(t.currentCharacter) {
			token = t.readNumber()
			return token
		} else {
			token = alaya_token.New(alaya_token.ILLEGAL, string(t.currentCharacter))
		}
	}
	t.Advance()
	return token
}