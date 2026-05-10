package main

import (
	"fmt"
	"github.com/mheaton92/quay/internal/config"
	"github.com/mheaton92/quay/internal/connection"
	"log"
)

func main() {
	err := config.EnsureConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	store := connection.Store{}
	err = store.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn := connection.NewConnection("Test", "192.168.1.1", "root", 22)
	err = store.Add(conn)
	if err != nil {
		log.Fatal(err)
	}

	for _, conn := range store.Connections {
		fmt.Println(conn.Name, conn.IP)
	}

}
