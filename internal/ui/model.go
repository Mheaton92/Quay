package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mheaton92/quay/internal/config"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/monitor"
	"github.com/mheaton92/quay/internal/ui/form"
	uikeys "github.com/mheaton92/quay/internal/ui/keys"
	"github.com/mheaton92/quay/internal/ui/network"
	"github.com/mheaton92/quay/internal/ui/scp"
)

type Panel int

const (
	ConnectionListPanel Panel = 0
	DetailPanel         Panel = 1
)

type Model struct {
	store         connection.Store
	focused       Panel
	width         int
	height        int
	err           error
	cursor        int
	form          *form.Model
	showForm      bool
	formActive    bool
	confirmDelete bool
	scpModel      *scp.Model
	showSCP       bool
	scpActive     bool
	keysModel     *uikeys.Model
	showKeys      bool
	keysActive    bool
	keybinds      config.Keybinds
	showHelp      bool
	monitor       *monitor.Monitor
	pinnedHosts   []string
	networkModel  *network.Model
	showNetwork   bool
	networkActive bool
	activePanel   string
}

func NewModel(store connection.Store) Model {
	m := monitor.NewMonitor(false) // false = use exec ping
	m.Start()

	for _, conn := range store.Connections {
		m.Add(conn.Host)
	}

	// Load persistent pins
	pins, _ := config.LoadPins()

	return Model{
		store:       store,
		focused:     ConnectionListPanel,
		keybinds:    config.DefaultKeybinds(),
		monitor:     m,
		pinnedHosts: pins,
	}
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m *Model) togglePin(host string) {
	for i, h := range m.pinnedHosts {
		if h == host {
			m.pinnedHosts = append(m.pinnedHosts[:i], m.pinnedHosts[i+1:]...)
			return
		}
	}
	m.pinnedHosts = append(m.pinnedHosts, host)
}

func (m *Model) togglePersistentPin(host string) {
	m.togglePin(host)
	config.SavePins(m.pinnedHosts)
}
