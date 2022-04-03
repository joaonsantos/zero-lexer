package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/joaonsantos/interpreter/lexer"
	"github.com/joaonsantos/interpreter/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("%s", prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)
		for {
			tok := lex.NextToken()
			if tok.Type == token.EOF {
				break
			}
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
