package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/connection"
	"strings"
	"net"
)

func (m Model) View() string {
    styles := DefaultStyles()
    nameLen := maxNameLen(m.store.Connections)
    width := nameLen + 14
    separator := strings.Repeat("─", width)
    dot := styles.Dot.Render("● ")

    var output string
    output += styles.Header.
        Width(width).
        Align(lipgloss.Center).
        Render(fmt.Sprintf("CONNECTIONS %d", len(m.store.Connections))) + "\n"
    output += separator + "\n"

    for i, conn := range m.store.Connections {
        if i == m.cursor {
            output += "▶ "
        } else {
            output += "  "
        }
        var display string
        if net.ParseIP(conn.Host) != nil {
            display = "." + lastOctet(conn.Host)
        } else {
            display = conn.Host
        }
        output += dot + fmt.Sprintf("%-10s", conn.Name) + " " + display + "\n"
    }

    return styles.Panel.Render(output)
}

func lastOctet(ip string) string {
    parts := strings.Split(ip, ".")
    if len(parts) < 2 {
        return ip
    }
    return parts[len(parts)-1]
}

func maxNameLen(connections []connection.Connection) int {
	max := 0
	for _, conn := range connections {
		if len(conn.Name) > max {
			max = len(conn.Name)
		}
	}
	return max
}
