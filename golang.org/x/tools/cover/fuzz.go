package cover

import (
	"io/ioutil"

	"golang.org/x/tools/cover"
)

func Fuzz(data []byte) int {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return 0
	}
	f.Close()

	ioutil.WriteFile(f.Name(), data, 0644)

	_, err = cover.ParseProfiles(f.Name())
	if err != nil {
		return 0
	}
	return 1
}
