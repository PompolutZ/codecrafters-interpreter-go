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
			s.tokens = append(s.tokens, NewToken(LEFT_PAREN))
		case ')':
			s.tokens = append(s.tokens, NewToken(RIGHT_PAREN))
		case '{':
			s.tokens = append(s.tokens, NewToken(LEFT_BRACE))
		case '}':
			s.tokens = append(s.tokens, NewToken(RIGHT_BRACE))
		case ',':
			s.tokens = append(s.tokens, NewToken(COMMA))
		case '.':
			s.tokens = append(s.tokens, NewToken(DOT))
		case '-':
			s.tokens = append(s.tokens, NewToken(MINUS))
		case '+':
			s.tokens = append(s.tokens, NewToken(PLUS))
		case ';':
			s.tokens = append(s.tokens, NewToken(SEMICOLON))
		case '*':
			s.tokens = append(s.tokens, NewToken(STAR))
		case '=':
			if s.match('=') {
				s.tokens = append(s.tokens, NewToken(EQUAL_EQUAL))
			} else {
				s.tokens = append(s.tokens, NewToken(EQUAL))
			}
		case '!':
			if s.match('=') {
				s.tokens = append(s.tokens, NewToken(BANG_EQUAL))
			} else {
				s.tokens = append(s.tokens, NewToken(BANG))
			}
		case '<':
			if s.match('=') {
				s.tokens = append(s.tokens, NewToken(LESS_EQUAL))
			} else {
				s.tokens = append(s.tokens, NewToken(LESS))
			}
		case '>':
			if s.match('=') {
				s.tokens = append(s.tokens, NewToken(GREATER_EQUAL))
			} else {
				s.tokens = append(s.tokens, NewToken(GREATER))
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
