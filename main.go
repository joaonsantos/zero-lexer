package main

import (
	"fmt"
	"os"

	"github.com/joaonsantos/interpreter/repl"
)

func main() {
	fmt.Println("Welcome to the Zero REPL")
	repl.Start(os.Stdin, os.Stdout)
}
