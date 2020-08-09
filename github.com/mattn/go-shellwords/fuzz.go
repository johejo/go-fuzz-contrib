package shellwords

import "github.com/mattn/go-shellwords"

func Fuzz(data []byte) int {
	_, err := shellwords.Parse(string(data))
	if err != nil {
		return 0
	}
	return 1
}
