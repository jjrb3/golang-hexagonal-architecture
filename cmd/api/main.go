package main

import (
	"log"

	"github.com/jjrb3/golang-hexagonal-architecture/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
