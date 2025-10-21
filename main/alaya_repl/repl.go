package alaya_repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/TcgVanguardTroll/Alaya/main/alaya_parser"
	"github.com/TcgVanguardTroll/Alaya/main/alaya_token"
	"github.com/TcgVanguardTroll/Alaya/main/alaya_tokenizer"
)

const PROMPT = ">> "

// Start begins the REPL (Read-Eval-Print Loop) for the Alaya interpreter.
// It reads expressions from the input, parses and evaluates them, then prints results.
// Similar to: public static void start(Reader in, Writer out) in Java
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		// Read input
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "" {
			continue
		}

		// Tokenize and parse
		tokenizer := alaya_tokenizer.New(line)
		parser := alaya_parser.NewParser(tokenizer)

		// Initialize parser with first token
		parser.CurrentToken = parser.Tokenizer.GetNextToken()

		// Evaluate the expression and print result
		result := parser.Expr()
		fmt.Fprintf(out, "%d\n", result)

		// Check if there are unexpected tokens remaining
		if parser.CurrentToken.TokenType != alaya_token.EOF {
			fmt.Fprintf(out, "Warning: unexpected token '%s' after expression\n",
				parser.CurrentToken.TokenValue)
		}
	}
}
