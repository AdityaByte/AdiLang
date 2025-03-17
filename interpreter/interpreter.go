package interpreter

import (
	"fmt"
	"log"

	"github.com/AdityaByte/AdiLang/parser"
)

func executeStatement(nodes []*parser.ASTNode, env *Environment) error {
	for _, node := range nodes {
		// fmt.Println("node type:", node.Type)
		switch node.Type {
		case parser.NodeVariableDeclaration:
			if err := executeVariableDeclaration(node, env); err != nil {
				return err
			}
		case parser.NodePrint:
			if err := executePrintStatement(node, env); err != nil {
				return err
			}
		case parser.NodeForLoop:
			if err := executeForLoop(node, env); err != nil {
				return err
			}
		case parser.NodeIfStatement:
			if err := executeIfStatement(node, env); err != nil {
				return err
			}
		case parser.NodeBlock:
			if err := executeBlock(node, env); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unknown statement : %v", node.Type)
		}
	}

	return nil
}

func executeBlock(node *parser.ASTNode, parentEnv *Environment) error {

	blockEnv := NewEnvironment(parentEnv)

	for _, stmt := range node.Children {
		if err := executeStatement([]*parser.ASTNode{stmt}, blockEnv); err != nil {
			return err
		}
	}

	return nil
}

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

	if node.Children != nil {
		anotherValue, err := evaluateExpression(node.Children[0], env)
		if err != nil {
			return err
		}

		log.Println(value.(string) + anotherValue.(string))
		return nil
	}

	log.Println(value)
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

func executeForLoop(node *parser.ASTNode, env *Environment) error {

	loopVar := node.Value.(string)
	rangeNode := node.Children[0]
	body := node.Children[1] // NodeBlock

	// fmt.Printf("body: %v and its type %v", body, body.Children)
	// fmt.Printf("body: %T and its type %T", body, body.Children)

	if rangeNode.Type != parser.NodeRange {
		return fmt.Errorf("expected range")
	}

	limit := rangeNode.Value.(int)

	loopEnv := NewEnvironment(env)

	for i := 0; i < limit; i++ {
		loopEnv.Set(loopVar, i)
		// fmt.Println("loopvar value: ", loopEnv)
		if err := executeBlock(body, loopEnv); err != nil {
			return err
		}
	}
	return nil
}

func executeIfStatement(node *parser.ASTNode, env *Environment) error {

	if len(node.Children) < 2 {
		return fmt.Errorf("invalid if statement missing conditon and body")
	}

	cond := node.Children[0]
	body := node.Children[1]

	left, err := evaluateExpression(cond.Children[0], env)

	if err != nil {
		return fmt.Errorf("Error evaluating left hand side ", err)
	}

	right, err := evaluateExpression(cond.Children[1], env)

	if err != nil {
		return fmt.Errorf("Error evaluating right hand side ", err)
	}

	operator, ok := cond.Value.(string) // If the thing is ok it return true otherwise false

	if !ok {
		return fmt.Errorf("Invalid condition operator: %v", operator)
	}

	var result bool

	// fmt.Println("left value:", left)
	// fmt.Println("right value:", right)
	// fmt.Println("Operator value:", operator)
	// fmt.Printf("%v -> %T and %v -> %T and %v -> %T \n", left, left, right, right, operator, operator)

	switch operator {
	case "==":
		result = left == right
	case ">":
		result = left.(int) > right.(int)
	case "<":
		result = left.(int) < right.(int)
	case "!=":
		result = left != right
	default:
		return fmt.Errorf("Unsupported operator: %v", operator)
	}

	if result {
		if err := executeBlock(body, env); err != nil {
			return fmt.Errorf("Error executing in body: %v", err)
		}
	}

	return nil
}

func Interpret(ast []*parser.ASTNode, env *Environment) error {
	return executeStatement(ast, env)
}
