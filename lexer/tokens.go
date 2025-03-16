package lexer

type TokenType string

const (
	// Identifiers and literals 
	Identifier TokenType = "IDENTIFIER"
	NumberLiteral TokenType = "NUMBER"
	StringLiteral TokenType = "STRING"

	// Keywords :
	VarKeyword TokenType = "VARIABLE"
	OutKeyword TokenType = "OUTPUT"
	IfKeyword TokenType = "IF"
	ElseKeyword TokenType = "ELSE"
	ForDudeKeyword TokenType = "FOR_DUDE" // For for loop
	InKeyword TokenType = "IN"
	RangeKeyword TokenType = "RANGE"

	// Operators
	AssignOperator TokenType = "ASSIGN"
	PlusOperator TokenType = "PLUS"
	MinusOperator TokenType = "MINUS"
	PrintOperator TokenType = "PRINTOPERATOR" // ->
	GreaterThanOperator TokenType = "GREATERTHAN" // >
	LessThanOperator TokenType = "LESSTHAN" // <
	ComparisionOperator TokenType = "COMPARISION" // ==

	// Brackets
	LBrace TokenType = "LEFTBRACE"
	RBrace TokenType = "RIGHTBRACE"
	LParen TokenType = "LEFTPARENTHESIS"
	RParen TokenType = "RIGHTPARENTHESIS"

	// Special Case
	IllegalToken TokenType = "ILLEGAL"
)

// Structure of the Token

type Token struct {
	Type TokenType
	Value string
}