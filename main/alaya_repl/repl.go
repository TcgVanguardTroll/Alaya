package alaya_repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/TcgVanguardTroll/Alaya/main/alaya_parser"
	"github.com/TcgVanguardTroll/Alaya/main/alaya_tokenizer"
)

const PROMPT = " ** "

func Start(in io.Reader, out io.Writer) {
    defer in.Close()
    defer out.Close()
    scanner := bufio.NewScanner(in)
    t := alaya_tokenizer.New("")
    p := alaya_parser.Parser{Tokenizer: t, CurrentToken: t.GetNextToken()}
    operators := []alaya_token.Type{alaya_token.ASTERISK, alaya_token.SLASH, alaya_token.PLUS, alaya_token.MINUS}
    for {
        fmt.Println(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }
        line := scanner.Text()
        t.text = line
        t.position = 0
        t.currentCharacter = t.text[t.position]
        p.CurrentToken = t.GetNextToken()
        fmt.Printf("%+v\n", p.Expr())
    }
}

func (p *Parser) isMatch(tokenType alaya_token.Type) {
    switch p.CurrentToken.TokenType {
    case tokenType:
        p.CurrentToken = p.Tokenizer.GetNextToken()
    default:
        panic("There was no Match !!")
    }
}

func (p *Parser) term() BinOp {
    binNode, _ := p.factor()
    for contains(p.CurrentToken.TokenType, operators) {
        tok := p.CurrentToken
        switch tok.TokenType {
        case alaya_token.ASTERISK:
			p.isMatch(alaya_token.ASTERISK)
			binNode = NewBinOp(binNode, tok, p.factor())
		case alaya_token.SLASH:
			p.isMatch(alaya_token.SLASH)
			binNode = NewBinOp(binNode, tok, p.factor())
			}
		}
		return binNode
	}

func (p *Parser) Expr() BinOp {
			currentNode := p.term()
			for contains(p.CurrentToken.TokenType, operators) {
			tok := p.CurrentToken
			switch tok.TokenType {
			case alaya_token.PLUS:
			p.isMatch(alaya_token.PLUS)
			currentNode = NewBinOp(currentNode, tok, p.term())
			case alaya_token.MINUS:
			p.isMatch(alaya_token.MINUS)
			currentNode = NewBinOp(currentNode, tok, p.term())
			}
			}
			return currentNode
			}

func (p *Parser) parse() {
	p.CurrentToken = p.Tokenizer.GetNextToken()
	return p.Expr()
	}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
		return
		}
		line := scanner.Text()
		t := alaya_tokenizer.New(line)
		p := alaya_parser.Parser{Tokenizer: t}
		fmt.Println(p.parse())
		}
	}