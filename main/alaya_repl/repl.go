package alaya_repl

import (
	"Alaya/main/alaya_parser"
	"Alaya/main/alaya_tokenizer"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = " ** "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		var t = alaya_tokenizer.New(line)
		var p = alaya_parser.Parser{
			Tokenizer:    t,
			CurrentToken: t.GetNextToken(),
		}
		//for tok := t.GetNextToken(); tok.TokenType != alaya_token.EOF; tok = t.GetNextToken() {
		//	fmt.Printf("%+v\n", tok)
		//}
		//for tok := p.Expr() ; p.CurrentToken.TokenType != alaya_token.EOF ; {
		//	fmt.Printf("%+v\n", tok)
		//}
		fmt.Printf("%+v\n", p.Expr())

	}
}
