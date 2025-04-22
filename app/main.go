package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	if len(fileContents) > 0 {
		runFileContents(fileContents)
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}

func runFileContents(fileContents []byte) {
	source := string(fileContents)
	lines := strings.Split(source, "\n")
	tokens := make([]Token, 0)
	for i, line := range lines {
		tokens = processLine(line, i+1, tokens)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println("EOF  null") 
}

func processLine(source string, line int, tokens []Token) []Token {
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
		default:
			printScannerError(line, "Unexpected character", string(char))
		}
	}

	return tokens
}

func printScannerError(line int, where string, msg string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", line, where, msg)
	os.Exit(1)
}
