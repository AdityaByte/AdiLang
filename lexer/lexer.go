package lexer

import (
	"strings"
	"unicode"
)

func Lexer(input string) []Token {
	var tokens []Token
	var currentToken strings.Builder
	chars := []rune(input)
	i := 0
	length := len(chars)

	for i < length {
		char := chars[i]

		// Skipping the spaces.
		if unicode.IsSpace(char) {
			i++
			continue
		}

		// Handling comments
		if char == '/' && i+1 < length && chars[i+1] == '/' {
			for i < length && chars[i] != '\n' {
				i++
			}
			continue
		} else if char == '%' {
			i++
			for i < length && chars[i] != '%' {
				i++
			}
			i++ // Skipping the last %
			continue
		}

		// Handling strings.
		if char == '"' {
			currentToken.Reset()
			i++
			for i < length && chars[i] != '"' {
				currentToken.WriteRune(chars[i])
				i++
			}
			tokens = append(tokens, Token{StringLiteral, currentToken.String()})
			// Resetting the current token
			currentToken.Reset()
			i++ // skipping closing "
			continue
		}

		// Handling multicharacters
		if char == '-' && i+1 < length && chars[i+1] == '>' {
			tokens = append(tokens, Token{PrintOperator, "->"})
			i += 2
			continue
		}

		// Handling multicharacters -> ==
		if char == '=' && i+1 < length && chars[i+1] == '='{
			tokens = append(tokens, Token{ComparisionOperator, "=="})
			i += 2
			continue
		}

		// Handling multicharacters -> !=
		if char == '!' && i+1 < length && chars[i+1] == '=' {
			tokens = append(tokens, Token{NotEqualsOperator, "!="})
			i += 2
			continue
		}

		// Handling single character tokens
		if isDelimiter(char) {
			if currentToken.Len() > 0 {
				tokens = append(tokens, classifyToken(currentToken.String()))
				currentToken.Reset()
			}

			switch char {
			case '=':
				tokens = append(tokens, Token{AssignOperator, "="})
			case '(':
				tokens = append(tokens, Token{LParen, "("})
			case ')':
				tokens = append(tokens, Token{RParen, ")"})
			case '{':
				tokens = append(tokens, Token{LBrace, "{"})
			case '}':
				tokens = append(tokens, Token{RBrace, "}"})
			case '>':
				tokens = append(tokens, Token{GreaterThanOperator, ">"})
			case '<':
				tokens = append(tokens, Token{LessThanOperator, "<"})
			case '+':
				tokens = append(tokens, Token{PlusOperator, "+"})
			}
			i++
			continue
		}

		// Handling identifiers/keywords/numbers
		currentToken.WriteRune(char)
		i++

		// checking the next character is the part of the same token
		if i < length && !isDelimiterOrSpace(chars[i]) {
			continue
		}

		// classify and reset the current token
		tokens = append(tokens, classifyToken(currentToken.String()))
		currentToken.Reset()
	}
	return tokens
}

func isDelimiter(char rune) bool {
	switch char {
	case '=', '(', ')', '{', '}', '-', '>', '<', '+': // Added < in this
		return true
	default:
		return false
	}
}

func isDelimiterOrSpace(char rune) bool {
	return unicode.IsSpace(char) || isDelimiter(char)
}

func classifyToken(input string) Token {
	switch input {
	case "var":
		return Token{VarKeyword, input}
	case "out":
		return Token{OutKeyword, input}
	case "ifdude":
		return Token{IfKeyword, input}
	case "else":
		return Token{ElseKeyword, input}
	case "fordude":
		return Token{ForDudeKeyword, input}
	case "in":
		return Token{InKeyword, input}
	case "range":
		return Token{RangeKeyword, input}
	}

	if isNumber(input) {
		return Token{NumberLiteral, input}
	}

	return Token{Identifier, input}
}

func isNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
