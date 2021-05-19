package repl

// Package repl implements the Read-Eval-Print-Loop or interactive console
// by lexing, parsing and evaluating the input in the interpreter

import (
	"bufio"
	"fmt"
	"io"

	"Primate/lexer"
	"Primate/token"
)

// PROMPT is the REPL prompt displayed for each input
const PROMPT = ">> "

// Start starts the REPL in a continious loop
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
