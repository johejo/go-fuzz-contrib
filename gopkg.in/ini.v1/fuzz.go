package ini

import "gopkg.in/ini.v1"

func Fuzz(data []byte) int {
	_, err := ini.Load(data)
	if err != nil {
		return 0
	}
	return 1
}
