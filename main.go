package main

import (
	"fmt"
	"os"

	"github.com/AdityaByte/AdiLang/interpreter"
	"github.com/AdityaByte/AdiLang/lexer"
	"github.com/AdityaByte/AdiLang/parser"
)

func printToken(tokens []lexer.Token) {
	for _, token := range tokens {
		fmt.Printf("%+v \n", token)
	}
}


func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage adilang <filename>.adi")
		return
	}

	filename := os.Args[1]

	code, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error occured", err)
		return
	}

	sourceCode := string(code)

	tokens := lexer.Lexer(sourceCode)

	printToken(tokens)

	parser := parser.Parser{Tokens: tokens, Pos: 0}

	astNodes := parser.Parse()

	for _, astNode := range astNodes {
		fmt.Println("%v", astNode)
	}

	fmt.Println("**********************************")

	// Creating a new Environment
	env := interpreter.NewEnvironment()

	if err := interpreter.Interpret(astNodes, env); err != nil {
		fmt.Println("Error:", err)
	}
}
