package main

import (
	tea "github.com/charmbracelet/bubbletea"
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
	p := tea.NewProgram(&m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	/* Test connections
	   store.Add(connection.NewConnection("Proxmox", "192.168.4.90", "root", 22))
	   store.Add(connection.NewConnection("TrueNAS", "192.168.4.91", "truenas_admin", 22))
	   store.Add(connection.NewConnection("AdGuard", "192.168.4.92", "root", 22))
	*/

}
