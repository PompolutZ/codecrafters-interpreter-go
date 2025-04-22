package scanner

import (
	"fmt"
	"os"
)

type Scanner struct {
	hasErrors bool
	source string 
	tokens []Token
	current int
	line int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		hasErrors: false,
		tokens: make([]Token, 0),
		current: 0,
		line: 1,
	}
}

// Scan processes the source code and extracts tokens
func (s *Scanner) Scan() {
	s.extractTokens()

	for _, token := range s.tokens {
		fmt.Println(token)
	}

	fmt.Println("EOF  null")
	
	if s.hasErrors {
		os.Exit(65)
	}
}

func (s *Scanner) extractTokens() {
	for !s.isAtEnd() {
		char := s.advance()
		switch char {
		case '(':
			s.tokens = append(s.tokens, NewToken(LEFT_PAREN, string(char), nil))
		case ')':
			s.tokens = append(s.tokens, NewToken(RIGHT_PAREN, string(char), nil))
		case '{':
			s.tokens = append(s.tokens, NewToken(LEFT_BRACE, string(char), nil))
		case '}':
			s.tokens = append(s.tokens, NewToken(RIGHT_BRACE, string(char), nil))
		case ',':
			s.tokens = append(s.tokens, NewToken(COMMA, string(char), nil))
		case '.':
			s.tokens = append(s.tokens, NewToken(DOT, string(char), nil))
		case '-':
			s.tokens = append(s.tokens, NewToken(MINUS, string(char), nil))
		case '+':
			s.tokens = append(s.tokens, NewToken(PLUS, string(char), nil))
		case ';':
			s.tokens = append(s.tokens, NewToken(SEMICOLON, string(char), nil))
		case '*':
			s.tokens = append(s.tokens, NewToken(STAR, string(char), nil))
		case '=':
			if s.match('=') {
				s.tokens = append(s.tokens, NewToken(EQUAL_EQUAL, "==", nil))
			} else {
				s.tokens = append(s.tokens, NewToken(EQUAL, string(char), nil))
			}
		case '!':
			if s.match('=') {
				s.tokens = append(s.tokens, NewToken(BANG_EQUAL, "!=", nil))
			} else {
				s.tokens = append(s.tokens, NewToken(BANG, string(char), nil))
			}

		default:
			printScannerError(s.line, "Unexpected character", string(char))
			s.hasErrors = true
		}
	}
}

func (s *Scanner) advance() byte {
	char := s.source[s.current]
	s.current++
	return char
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func printScannerError(line int, where string, msg string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", line, where, msg)
}
