package scanner

import (
	"fmt"
	"os"
	"strings"
)

type Scanner struct {
	hasErrors bool
	source string 
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		hasErrors: false,
	}
}

func (s *Scanner) Scan() {
	lines := strings.Split(s.source, "\n")
	tokens := make([]Token, 0)
	
	for i, line := range lines {
		tokens = s.processLine(line, i+1, tokens)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println("EOF  null")
	
	if s.hasErrors {
		os.Exit(65)
	}
}

func (s *Scanner) processLine(source string, line int, tokens []Token) []Token {
	for _, char := range source {
		switch char {
		case '(':
			tokens = append(tokens, NewToken(LEFT_PAREN, string(char), nil))
		case ')':
			tokens = append(tokens, NewToken(RIGHT_PAREN, string(char), nil))
		case '{':
			tokens = append(tokens, NewToken(LEFT_BRACE, string(char), nil))
		case '}':
			tokens = append(tokens, NewToken(RIGHT_BRACE, string(char), nil))
		case ',':
			tokens = append(tokens, NewToken(COMMA, string(char), nil))
		case '.':
			tokens = append(tokens, NewToken(DOT, string(char), nil))
		case '-':
			tokens = append(tokens, NewToken(MINUS, string(char), nil))
		case '+':
			tokens = append(tokens, NewToken(PLUS, string(char), nil))
		case ';':
			tokens = append(tokens, NewToken(SEMICOLON, string(char), nil))
		case '*':
			tokens = append(tokens, NewToken(STAR, string(char), nil))
		default:
			printScannerError(line, "Unexpected character", string(char))
			s.hasErrors = true
		}
	}

	return tokens
}

func printScannerError(line int, where string, msg string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", line, where, msg)
}
