package parser

type NodeType string

const (

	// Basic Node type
	NodeProgram NodeType = "PROGRAM"
	NodePrint NodeType = "PRINT" // out->"hello-world"
	NodeVariableDeclaration NodeType = "VARIABLE_DECLARATION" // var(name="aditya")
	NodeIfStatement NodeType = "IF_STATEMENT"
	NodeBlock NodeType = "BLOCK"
	NodeForLoop NodeType = "FOR_LOOP"
	NodeRange NodeType = "RANGE"

	// Expression Node type
	NodeStringLiteral NodeType = "STRING_LITERAL"
	NodeNumberLiteral NodeType = "NUMBER_LITERAL"
	NodeIdentifier NodeType = "IDENTIFIER"
	NodeBinaryOperation NodeType = "BINARY_OPERATION"

	// Operator Node type
	NodeComparision NodeType = "COMPARISION"
	NodeGreaterThan NodeType = "GREATERTHAN"
	NodeLessThan NodeType = "LESSTHAN"
	NodeNotEquals NodeType = "NOTEQUALS"

	// Condition node type
	NodeCondition NodeType = "CONDITION"
)

type ASTNode struct {
	Type NodeType
	Value interface{}
	Children []*ASTNode
}