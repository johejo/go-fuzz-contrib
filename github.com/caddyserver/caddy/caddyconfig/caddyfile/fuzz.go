package caddyfile

import "github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"

func Fuzz(data []byte) int {
	_, err := caddyfile.Parse("Caddyfile", data)
	if err != nil {
		return 0
	}
	return 1
}
