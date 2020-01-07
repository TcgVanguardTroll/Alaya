package repl

import (
	"Alaya/main/alaya_token"
	"Alaya/main/tokenizer"
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
		var t = tokenizer.New(line)
		for tok := t.GetNextToken(); tok.TokenType != alaya_token.EOF; tok = t.GetNextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
