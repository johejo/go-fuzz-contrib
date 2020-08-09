package parser

import "github.com/mattn/anko/parser"

func Fuzz(data []byte) int {
	_, err := parser.ParseSrc(string(data))
	if err != nil {
		return 0
	}
	return 1
}
