package main

import (
	"log"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func main() {
	_, err := caddyfile.Parse("Caddyfile", []byte("import /proc/self/p*"))
	log.Println(err)
}
