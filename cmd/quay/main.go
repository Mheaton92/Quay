package main

import (
	"github.com/mheaton92/quay/internal/config"
	"github.com/mheaton92/quay/internal/connection"
	"log"
)

func main() {
	err := config.EnsureConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	store, err := connection.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	_ = store

}
