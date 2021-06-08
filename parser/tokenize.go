package parser

import (
	"strings"
)

func Tokenize(input string) []string {
	tokens := strings.Split(input, "|")
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.Trim(tokens[i], "\t \n")
	}
	return tokens
}
