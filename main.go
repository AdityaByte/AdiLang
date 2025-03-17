package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AdityaByte/AdiLang/interpreter"
	"github.com/AdityaByte/AdiLang/lexer"
	"github.com/AdityaByte/AdiLang/parser"
)

func printToken(tokens []lexer.Token) {
	for _, token := range tokens {
		fmt.Printf("%+v \n", token)
	}
}

func printAST(nodes []*parser.ASTNode, indent string) {
    for _, node := range nodes {
        if node == nil {
            fmt.Printf("%s<nil>\n", indent)
            continue
        }
        fmt.Printf("%sType: %v, Value: %v\n", indent, node.Type, node.Value)
        if len(node.Children) > 0 {
            printAST(node.Children, indent+"  ")
        }
    }
}


func main() {

	if len(os.Args) < 2 {
		log.Println("Usage adilang <filename>.adi")
		return
	}

	filename := os.Args[1]

	if !strings.HasSuffix(filename, ".adi") {
		log.Println("File extension must be .adi")
		return
	}

	code, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("Error Reading file", err)
		return
	}

	sourceCode := string(code)

	tokens := lexer.Lexer(sourceCode)

	// printToken(tokens)

	parser := parser.Parser{Tokens: tokens, Pos: 0}

	astNodes, err := parser.Parse()

	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	// for _, astNode := range astNodes {
	// 	fmt.Println("%v", astNode)
	// }

	// printAST(astNodes, "")

	// fmt.Println("**********************************")

	// Creating a new Environment
	env := interpreter.NewEnvironment(nil)

	if err := interpreter.Interpret(astNodes, env); err != nil {
		log.Fatal("Error:", err)
	}
}
