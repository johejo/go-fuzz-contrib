// +build gofuzz

package timefmt

import (
	fuzz "github.com/google/gofuzz"
	timefmt "github.com/itchyny/timefmt-go"
)

func Fuzz(data []byte) int {
	type Param struct {
		Source string
		Format string
	}
	var p Param
	fuzz.NewFromGoFuzz(data).Fuzz(&p)
	_, err := timefmt.Parse(p.Source, p.Format)
	if err != nil {
		return 0
	}
	return 1
}
