package main

import (
	"flag"
	"io"
	"os"

	interpreter "github.com/tejasdeepakmasne/bfit/Interpreter"
	"github.com/tejasdeepakmasne/bfit/lexer"
)

func main() {
	flag.Parse()
	filename := flag.Args()
	file, err := os.Open(filename[0])
	if err != nil {
		panic(err)
	}

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	tokens := lexer.GenerateTokens(input)
	interpreter.Interpret(tokens)
}
