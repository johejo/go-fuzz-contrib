package gojq

import "github.com/itchyny/gojq"

func Fuzz(data []byte) int {
	_, err := gojq.Parse(string(data))
	if err != nil {
		return 0
	}
	return 1
}
