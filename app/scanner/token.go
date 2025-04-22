package scanner

import "fmt"

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	STAR
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
)

var lexemes = map[TokenType]string{
	LEFT_PAREN:   "(",
	RIGHT_PAREN:  ")",
	LEFT_BRACE:   "{",
	RIGHT_BRACE:  "}",
	COMMA:        ",",
	DOT:         ".",
	MINUS:       "-",
	PLUS:        "+",
	SEMICOLON:   ";",
	STAR:        "*",
	EQUAL:       "=",
	EQUAL_EQUAL: "==",
	BANG:        "!",
	BANG_EQUAL:  "!=",
	LESS:        "<",
	LESS_EQUAL:  "<=",
	GREATER:     ">",
	GREATER_EQUAL: ">=",
}

func (t TokenType) String() string {
	return [...]string{
		"LEFT_PAREN",
		"RIGHT_PAREN",
		"LEFT_BRACE",
		"RIGHT_BRACE",
		"COMMA",
		"DOT",
		"MINUS",
		"PLUS",
		"SEMICOLON",
		"STAR",
		"EQUAL",
		"EQUAL_EQUAL",
		"BANG",
		"BANG_EQUAL",
		"LESS",
		"LESS_EQUAL",
		"GREATER",
		"GREATER_EQUAL",
	}[t]
}

type Token struct {
	Type TokenType
	Line int
}

func NewToken(tokenType TokenType) Token {
	return Token{
		Type:   tokenType,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%v %s %v", t.Type, lexemes[t.Type], "null")
}