package main

import (
	"github.com/mheaton92/quay/internal/config"
	"log"
)

func main() {
	err := config.EnsureConfigDir()
	if err != nil {
		log.Fatal(err)
	}

}
