package ui

import (
	"github.com/mheaton92/quay/internal/connection"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mheaton92/quay/internal/ui/form"
	"github.com/mheaton92/quay/internal/ui/scp"
	uikeys "github.com/mheaton92/quay/internal/ui/keys"
)

type Panel int

const (
	ConnectionListPanel Panel = 0
	DetailPanel Panel = 1
)

type Model struct {
	store connection.Store
	focused Panel
	width int
	height int
	err error
	cursor int
	form *form.Model
	showForm bool
	confirmDelete bool
	scpModel *scp.Model
	showSCP bool
	keysModel *uikeys.Model
	showKeys bool
}

func NewModel(store connection.Store) Model {
	return Model{
		store: 	 store,
		focused: ConnectionListPanel,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}