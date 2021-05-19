package repl

// Package repl implements the Read-Eval-Print-Loop or interactive console
// by lexing, parsing and evaluating the input in the interpreter

import (
	"bufio"
	"fmt"
	"io"

	"Primate/eval"
	"Primate/lexer"
	"Primate/parser"
)

// PROMPT is the REPL prompt displayed for each input
const PROMPT = ">> "

// ComputerFace is the REPL's face of shock and horror when you encounter a
// parser error :D
const ComputerFace = `
   _________________  
  |  _____________  | 
  | |             | | 
  | |    0   0    | | 
  | |      -      | | 
  | |    \___/    | | 
  | |____     ____| | 
  |______|\_/|______| 
     _|__|/ \|_|_
    / ********** \   
  /  ************  \ 
 --------------------
`

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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		obj := eval.Eval(program)
		if obj != nil {
			io.WriteString(out, obj.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ComputerFace)
	io.WriteString(out, "Woops! We ran into some primate business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
