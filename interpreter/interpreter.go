package interpreter

import (
	"fmt"

	"github.com/AdityaByte/AdiLang/parser"
)

func executeVariableDeclaration(node *parser.ASTNode, env *Environment) error {
	if node.Type != parser.NodeVariableDeclaration {
		return fmt.Errorf("expected variable declaration")
	}

	name := node.Value.(string)

	value, err := evaluateExpression(node.Children[0], env)

	if err != nil {
		return err
	}

	env.Set(name, value)
	return nil
}

func executePrintStatement(node *parser.ASTNode, env *Environment) error {
	value, err := evaluateExpression(node.Value.(*parser.ASTNode), env)
	if err != nil {
		return err
	}
	fmt.Println(value)
	return nil
}

func evaluateExpression(node *parser.ASTNode, env *Environment) (interface{}, error) {
	switch node.Type {
	case parser.NodeStringLiteral:
		return node.Value, nil
	case parser.NodeNumberLiteral:
		return node.Value, nil
	case parser.NodeIdentifier:
		return env.Get(node.Value.(string))
	default:
		return nil, fmt.Errorf("unsupported expression type: %s", node.Type)
	}
}

func Interpret(ast []*parser.ASTNode, env *Environment) error {
	for _,node := range ast {
		switch node.Type {
		case parser.NodeVariableDeclaration:
			if err := executeVariableDeclaration(node, env); err != nil {
				return err
			}
		case parser.NodePrint:
			if err := executePrintStatement(node, env); err != nil {
				return err
			} 
		default:
			return fmt.Errorf("unknown node type: %s", node.Type)
		}
	}
	return nil
}