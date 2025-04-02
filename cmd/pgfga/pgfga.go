package main

import (
	"log"

	"github.com/pgvillage-tools/pgfga/internal"
)

func main() {
	internal.Initialize()

	fga, err := internal.NewPgFgaHandler()
	if err != nil {
		log.Fatalf("Error occurred on getting config: %e", err)
	}

	fga.Handle()
}
