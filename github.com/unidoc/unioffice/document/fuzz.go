package document

import (
	"bytes"

	"github.com/unidoc/unioffice/document"
)

func Fuzz(data []byte) int {
	_, err := document.Read(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return 0
	}
	return 1
}
