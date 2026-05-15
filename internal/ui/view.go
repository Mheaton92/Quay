package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func (m Model) View() string {
	dot := "● "
	var output string
	for i, conn := range m.store.Connections {
		if i == m.cursor {
			output += "▶ "
		} else {
			output += "  "
		}
		output += dot + fmt.Sprintf("%-10s", conn.Name) + " ." + lastOctet(conn.IP) + "\n"
	}
	panelStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#58a6ff")).
		Padding(0, 1)

	return panelStyle.Render(output)

}

func lastOctet(ip string) string {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return ""
	}
	return parts[len(parts)-1]
}
