package presentation

import (
	"bytes"

	"github.com/unidoc/unioffice/presentation"
)

func Fuzz(data []byte) int {
	_, err := presentation.Read(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return 0
	}
	return 1
}
