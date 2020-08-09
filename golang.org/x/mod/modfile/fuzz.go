package modfile

import "golang.org/x/mod/modfile"

func Fuzz(data []byte) int {
	_, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return 0
	}
	return 1
}
