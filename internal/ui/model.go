package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui/form"
	uikeys "github.com/mheaton92/quay/internal/ui/keys"
	"github.com/mheaton92/quay/internal/ui/scp"
	"github.com/mheaton92/quay/internal/config"
	"github.com/mheaton92/quay/internal/monitor"
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
	confirmDelete bool
	scpModel      *scp.Model
	showSCP       bool
	keysModel     *uikeys.Model
	showKeys      bool
	keybinds      config.Keybinds
	showHelp       bool
	monitor			*monitor.Monitor
}

func NewModel(store connection.Store) Model {
	m := monitor.NewMonitor(false) // false = use exec ping
	m.Start()

	for _, conn := range store.Connections {
		m.Add(conn.Host)
	}

	return Model{
		store:   store,
		focused: ConnectionListPanel,
		keybinds: config.DefaultKeybinds(),
		monitor: m,
	}
}

func (m Model) Init() tea.Cmd {
	return tick()
}
