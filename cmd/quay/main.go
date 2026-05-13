package main

import (
	"github.com/mheaton92/quay/internal/config"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui"
	"log"
	tea "github.com/charmbracelet/bubbletea"
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
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}
