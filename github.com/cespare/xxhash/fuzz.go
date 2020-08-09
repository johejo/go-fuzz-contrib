package xxhash

import "github.com/cespare/xxhash/v2"

func Fuzz(data []byte) int {
	h := xxhash.New()
	_, err := h.Write(data)
	if err != nil {
		return 0
	}
	return 1
}
