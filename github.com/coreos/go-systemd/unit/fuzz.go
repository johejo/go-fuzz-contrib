package unit

import (
	"bytes"

	"github.com/coreos/go-systemd/unit"
)

func Fuzz(data []byte) int {
	_, err := unit.Deserialize(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	return 1
}
