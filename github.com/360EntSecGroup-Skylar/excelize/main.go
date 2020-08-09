package main

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	b := []byte("PK\x06\x06hips/imagePKmage" +
		"PK\x06\a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00" +
		"PK\x05\x06PK\x06\a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
		"\x01\x00\x00\x00PK\x05\x06PK\x05\x06\x01\u007f\xff\xff\xff\xf7g\x19" +
		"\xd5#\xe8\x03\x00\x00")
	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range zr.File {
		r, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
	}
}
