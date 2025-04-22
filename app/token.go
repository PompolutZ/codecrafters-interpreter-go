package main

import "fmt"

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
)

// String returns the string representation of the token type
func (t TokenType) String() string {
	return [...]string{
		"LEFT_PAREN",
		"RIGHT_PAREN",
	}[t]
}

type Token struct {
	Type TokenType
	Lexeme string
	Literal interface{}
	Line int
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
		Literal:   literal,
	}
}

func (t Token) String() string {
	literal := t.Literal
	if literal == nil {
		literal = "null"
	}

	return fmt.Sprintf("%v %s %v", t.Type, t.Lexeme, literal)
}