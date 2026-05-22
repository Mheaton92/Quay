package network

import (
    "github.com/mheaton92/quay/internal/connection"
    tea "github.com/charmbracelet/bubbletea"
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
}

func NewNetwork(conn connection.Connection) *Model {
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
    }
}

func (m *Model) Init() tea.Cmd {
    return nil
}

func (m *Model) Done() bool {
    return false
}
