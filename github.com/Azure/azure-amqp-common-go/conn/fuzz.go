package conn

import "github.com/Azure/azure-amqp-common-go/v3/conn"

func Fuzz(data []byte) int {
	_, err := conn.ParsedConnectionFromStr(string(data))
	if err != nil {
		return 0
	}
	return 1
}
