package main

import (
	"fmt"
	"os"
	"strings"
)

func scanFileContents(fileContents []byte) {
	source := string(fileContents)
	lines := strings.Split(source, "\n")
	tokens := make([]Token, 0)
	lineHadError := false
	
	for i, line := range lines {
		tokens, lineHadError = processLine(line, i+1, tokens)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println("EOF  null")
	
	if lineHadError {
		os.Exit(65)
	}
}

func processLine(source string, line int, tokens []Token) ([]Token, bool) {
	hadError := false
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
			hadError = true
		}
	}

	return tokens, hadError
}

func printScannerError(line int, where string, msg string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", line, where, msg)
}
