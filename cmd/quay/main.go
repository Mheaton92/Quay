package main

import (
	"github.com/mheaton92/quay/internal/config"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui"
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
	
	m := ui.NewModel(*store)
	_ = m

}
