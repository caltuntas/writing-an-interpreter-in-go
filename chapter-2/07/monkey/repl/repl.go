package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		fmt.Println("scanned text")
		if !scanned {
			return
		}

		line := scanner.Text()
		//line := "let x = 1 * 2;"
		fmt.Println(line)
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		fmt.Println("program parsing finished")
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n") 
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors { 
		io.WriteString(out, "\t"+msg+"\n") 
	} 
}
