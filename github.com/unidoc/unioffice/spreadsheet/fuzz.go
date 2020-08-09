package spreadsheet

import (
	"bytes"

	"github.com/unidoc/unioffice/spreadsheet"
)

func Fuzz(data []byte) int {
	_, err := spreadsheet.Read(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return 0
	}
	return 1
}
