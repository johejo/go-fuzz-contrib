package main

import (
	"log"

	"github.com/mattn/anko/parser"
)

func main() {
	_, err := parser.ParseSrc("\ue031")
	if err != nil {
		log.Println(err)
	}
}
