package excelize

import (
	"bytes"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Fuzz(data []byte) int {
	_, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	return 1
}
