package strftime

import (
	"github.com/lestrrat-go/strftime"
)

func Fuzz(data []byte) int {
	_, err := strftime.New(string(data))
	if err != nil {
		return 0
	}
	return 1
}
