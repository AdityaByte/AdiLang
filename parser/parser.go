package parser

import (
	"fmt"
	"strconv"

	"github.com/AdityaByte/AdiLang/lexer"
)

type Parser struct {
	Tokens []lexer.Token
	Pos    int
}

func (p *Parser) currentToken() lexer.Token {
	if p.Pos < len(p.Tokens) {
		return p.Tokens[p.Pos]
	}
	return lexer.Token{lexer.IllegalToken, ""}
}

func (p *Parser) nextToken() {
	p.Pos++
}

// for parsing the variable declaration.
func (p *Parser) parseVariableDeclaration() (*ASTNode, error) {
	if p.currentToken().Type != lexer.VarKeyword {
		return nil, fmt.Errorf("Expected 'var' keyword")
	}
	p.nextToken()

	if p.currentToken().Type != lexer.LParen {
		return nil, fmt.Errorf("Expected '(' keyword")
	}
	p.nextToken()

	if p.currentToken().Type != lexer.Identifier {
		return nil, fmt.Errorf("Expected 'identifier'")
	}
	ident := p.currentToken().Value
	p.nextToken()

	if p.currentToken().Type != lexer.AssignOperator {
		return nil, fmt.Errorf("Expected '=' keyword")
	}
	p.nextToken()

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.currentToken().Type != lexer.RParen {
		return nil, fmt.Errorf("Expected ')' closing paranthesis")
	}
	p.nextToken()

	return &ASTNode{
		Type:     NodeVariableDeclaration,
		Value:    ident,
		Children: []*ASTNode{expr},
	}, nil
}

// for parsing the print statement.
func (p *Parser) parsePrintStatement() (*ASTNode, error) {
	if p.currentToken().Type != lexer.OutKeyword {
		return nil, fmt.Errorf("Expected 'out' keyword")
	}
	p.nextToken()

	if p.currentToken().Type != lexer.PrintOperator {
		fmt.Println("Expected '->'")
		return nil, fmt.Errorf("Expected '->' keyword")
	}
	p.nextToken()

	expr, err := p.parseExpression()

	if err != nil {
		return nil, err
	}
	
	// fmt.Println(p.currentToken().Type, p.currentToken().Value)
	if p.currentToken().Type == lexer.PlusOperator {
		p.nextToken()
		anotherExpr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		return &ASTNode{
			Type: NodePrint,
			Value: expr,
			Children: []*ASTNode{
				anotherExpr, // Children at zero index
			},
		}, nil
	}

	return &ASTNode{
		Type:  NodePrint,
		Value: expr,
	}, nil
}

func (p *Parser) parseIfStatement() (*ASTNode, error) {
	if p.currentToken().Type != lexer.IfKeyword {
		return nil, fmt.Errorf("Expected if keyword")
	}
	p.nextToken()

	cond, err := p.parseCondition()
	if err != nil {
		return nil, err
	}

	body, err := p.parseBlock()

	if err != nil {
		return nil, err
	}

	return &ASTNode{
		Type: NodeIfStatement,
		Children: []*ASTNode{
			cond,
			body,
		},
	}, nil
}

func (p *Parser) parseCondition() (*ASTNode, error) {
	left, err := p.parsePrimary()

	if err != nil {
		return nil, err
	}

	// fmt.Println("left value : ", left.Value)

	operator, err := p.parseOperator()
	if err != nil {
		return nil, err
	}
	operatorValue := operator.Value

	// fmt.Println("Operator Value:", operatorValue)

	right, err := p.parsePrimary()

	// fmt.Println("Right Value :", right.Value)

	if err != nil {
		return nil, err
	}

	return &ASTNode{
		Type:  NodeCondition,
		Value: operatorValue,
		Children: []*ASTNode{
			left,
			right,
		},
	}, nil
}

func (p *Parser) parseOperator() (*ASTNode, error) {
	switch p.currentToken().Type {
	case lexer.ComparisionOperator:
		node := &ASTNode{
			Type:  NodeComparision,
			Value: p.currentToken().Value,
		}
		p.nextToken()
		return node, nil
	case lexer.GreaterThanOperator:
		node := &ASTNode{
			Type:  NodeGreaterThan,
			Value: p.currentToken().Value,
		}
		p.nextToken()
		return node, nil
	case lexer.LessThanOperator:
		node := &ASTNode{
			Type:  NodeLessThan,
			Value: p.currentToken().Value,
		}
		p.nextToken()
		return node, nil
	case lexer.NotEqualsOperator:
		node := &ASTNode{
			Type: NodeNotEquals,
			Value: p.currentToken().Value,
		}
		p.nextToken()
		return node, nil
	default:
		return nil, fmt.Errorf("Expected these '==, >, <'")
	}
}

func (p *Parser) parsePrimary() (*ASTNode, error) {
	switch p.currentToken().Type {
	case lexer.NumberLiteral:
		return p.parseNumberLiteral()
	case lexer.Identifier:
		return p.parseIdentifier()
	default:
		return nil, fmt.Errorf("Expected number, identifier")
	}
}

func (p *Parser) parseForLoop() (*ASTNode, error) {
	if p.currentToken().Type != lexer.ForDudeKeyword {
		return nil, fmt.Errorf("Expected 'fordude' keyword")
	}
	p.nextToken()

	if p.currentToken().Type != lexer.Identifier {
		return nil, fmt.Errorf("Expected identifier")
	}
	loopVar := p.currentToken().Value
	p.nextToken()

	if p.currentToken().Type != lexer.InKeyword {
		return nil, fmt.Errorf("Expected 'in' keyword")
	}
	p.nextToken()

	if p.currentToken().Type != lexer.RangeKeyword {
		return nil, fmt.Errorf("Expected 'range' keyword")
	}
	p.nextToken()

	if p.currentToken().Type != lexer.LParen {
		return nil, fmt.Errorf("Expected '(' keyword")
	}
	p.nextToken()

	if p.currentToken().Type != lexer.NumberLiteral {
		return nil, fmt.Errorf("Expected number literal")
	}
	limit, err := strconv.Atoi(p.currentToken().Value)

	if err != nil {
		return nil, fmt.Errorf("invalid number: %s", p.currentToken().Value)
	}
	p.nextToken()

	if p.currentToken().Type != lexer.RParen {
		return nil, fmt.Errorf("Expected ')' keyword")
	}
	p.nextToken()

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &ASTNode{
		Type:  NodeForLoop,
		Value: loopVar,
		Children: []*ASTNode{
			{
				Type:  NodeRange,
				Value: limit,
			},
			body,
		},
	}, nil

}

func (p *Parser) parseBlock() (*ASTNode, error) {
	// fmt.Println("current token in block:", p.currentToken().Value)
	if p.currentToken().Type != lexer.LBrace {
		return nil, fmt.Errorf("Expected 'if' keyword")
	}
	p.nextToken()

	var statements []*ASTNode

	for p.currentToken().Type != lexer.RBrace {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}

	if p.currentToken().Type != lexer.RBrace {
		return nil, fmt.Errorf("expected'}'")
	}
	p.nextToken()

	return &ASTNode{
		Type:     NodeBlock,
		Children: statements,
	}, nil
}

func (p *Parser) parseStatement() (*ASTNode, error) {
	switch p.currentToken().Type {
	case lexer.OutKeyword:
		return p.parsePrintStatement()
	case lexer.VarKeyword:
		return p.parseVariableDeclaration()
	case lexer.ForDudeKeyword:
		return p.parseForLoop()
	case lexer.IfKeyword:
		return p.parseIfStatement()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.currentToken())
	}
}

func (p *Parser) parseExpression() (*ASTNode, error) {
	switch p.currentToken().Type {
	case lexer.StringLiteral:
		return p.parseStringLiteral()
	case lexer.NumberLiteral:
		return p.parseNumberLiteral()
	case lexer.Identifier:
		return p.parseIdentifier()
	default:
		return nil, fmt.Errorf("Expected Expression (string, number or identifier)")
	}
}

func (p *Parser) parseStringLiteral() (*ASTNode, error) {
	node := &ASTNode{
		Type:  NodeStringLiteral,
		Value: p.currentToken().Value,
	}
	p.nextToken()
	return node, nil
}

func (p *Parser) parseNumberLiteral() (*ASTNode, error) {
	value, err := strconv.Atoi(p.currentToken().Value)
	if err != nil {
		return nil, fmt.Errorf("Invalid number: %s", p.currentToken().Value)
	}

	node := &ASTNode{
		Type:  NodeNumberLiteral,
		Value: value,
	}

	p.nextToken()
	return node, nil
}

func (p *Parser) parseIdentifier() (*ASTNode, error) {
	node := &ASTNode{
		Type:  NodeIdentifier,
		Value: p.currentToken().Value,
	}
	p.nextToken()
	return node, nil
}

// Main function which parse out the things.
func (p *Parser) Parse() ([]*ASTNode, error) {
	var Nodes []*ASTNode

	for p.Pos < len(p.Tokens) {
		token := p.currentToken()

		switch token.Type {
		case lexer.OutKeyword:
			astNode, err := p.parsePrintStatement()
			if err != nil {
				return nil, err
			}
			Nodes = append(Nodes, astNode)
		case lexer.VarKeyword:
			astNode, err := p.parseVariableDeclaration()
			if err != nil {
				return nil, err
			}
			Nodes = append(Nodes, astNode)
		case lexer.ForDudeKeyword:
			astNode, err := p.parseForLoop()
			if err != nil {
				return nil, err
			}
			Nodes = append(Nodes, astNode)
		case lexer.IfKeyword:
			astNode, err := p.parseIfStatement()
			if err != nil {
				return nil, err
			}
			Nodes = append(Nodes, astNode)
		case lexer.LBrace:
			astNode, err := p.parseBlock()
			if err != nil {
				return nil, err
			}
			Nodes = append(Nodes, astNode)
		default:
			p.nextToken()
		}
	}
	return Nodes, nil
}
