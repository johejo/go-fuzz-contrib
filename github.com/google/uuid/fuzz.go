package uuid

import "github.com/google/uuid"

func Fuzz(data []byte) int {
	id, err := uuid.Parse(string(data))
	if err != nil {
		return 0
	}
	_ = id.String()
	return 1
}
