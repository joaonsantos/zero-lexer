package main

import (
	"fmt"
	"os"

	"github.com/joaonsantos/zero-lexer/rlpl"
)

func main() {
	fmt.Println("Welcome to the zero rlpl!")
	rlpl.Start(os.Stdin, os.Stdout)
}
