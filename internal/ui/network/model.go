package network

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/charmbracelet/bubbles/textinput"
)

type Tool int

const (
	ToolPortScanner Tool = iota
	ToolWakeOnLAN
	ToolDNSLookup
	ToolTraceroute
	ToolSSLChecker
	ToolSubnetScanner
	ToolBandwidthTest
)

type Model struct {
	conn     connection.Connection
	cursor   int
	tools    []string
	result   string
	err      error
	showHelp bool
	input			textinput.Model
	showInput	bool
	inputLabel string
}

func NewNetwork(conn connection.Connection) *Model {
	ti := textinput.New()
	ti.Placeholder = "22,80,443,8006"
	ti.CharLimit = 256
	ti.Width = 40

	return &Model{
		conn: conn,
		tools: []string{
			"Port Scanner",
			"Wake on LAN",
			"DNS Lookup",
			"Traceroute",
			"SSL Checker",
			"Subnet Scanner",
			"Bandwidth Test",
		},
		input: ti,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Done() bool {
	return false
}
