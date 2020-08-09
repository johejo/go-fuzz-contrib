package zglob

import "github.com/mattn/go-zglob"

func Fuzz(data []byte) int {
	_, err := zglob.New(string(data))
	if err != nil {
		return 0
	}
	return 1
}
