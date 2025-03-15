package parser

type NodeType string

const (

	// Basic Node type
	NodeProgram NodeType = "PROGRAM"
	NodePrint NodeType = "PRINT" // out->"hello-world"
	NodeVariableDeclaration NodeType = "VARIABLE_DECLARATION" // var(name="aditya")
	NodeIfStatement NodeType = "IF_STATEMENT"
	NodeBlock NodeType = "BLOCK"

	// Expression Node type
	NodeStringLiteral NodeType = "STRING_LITERAL"
	NodeNumberLiteral NodeType = "NUMBER_LITERAL"
	NodeIdentifier NodeType = "IDENTIFIER"
	NodeBinaryOperation NodeType = "BINARY_OPERATION"
)

type ASTNode struct {
	Type NodeType
	Value interface{}
	Children []*ASTNode
}