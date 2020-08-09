package licensecheck

import "github.com/google/licensecheck"

func Fuzz(data []byte) int {
	_, ok := licensecheck.Cover(data, licensecheck.Options{})
	if !ok {
		return 0
	}
	return 1
}
