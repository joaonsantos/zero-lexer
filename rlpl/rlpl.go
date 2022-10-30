package rlpl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/joaonsantos/zero-lexer/lexer"
	"github.com/joaonsantos/zero-lexer/token"
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
		if idx := strings.Index(line, "//"); idx != -1 {
			line = line[:idx]
		}
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
