package ui

import "github.com/mheaton92/quay/internal/connection"

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
}

func NewModel(store connection.Store) Model {
	return Model{
		store: 	 store,
		focused: ConnectionListPanel,
	}
}