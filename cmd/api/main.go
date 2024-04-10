package main

import (
	"log"

	"github.com/gmd3v/gohex/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Start(); err != nil {
		log.Fatal(err)
	}
}
