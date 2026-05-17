package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui/theme"
	"github.com/mheaton92/quay/internal/ui/detail"
	"github.com/mheaton92/quay/internal/ui/statusbar"
	"net"
	"strings"
)

func (m Model) View() string {
	if m.width < 70 || m.height < 20 {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f85149")).
			Align(lipgloss.Center).
			Width(m.width).
			Render("\n\nTerminal too small\nMinimum size: 70x20\nCurrent: " +
				fmt.Sprintf("%dx%d", m.width, m.height))
	}
	styles := theme.DefaultStyles()
	nameLen := maxNameLen(m.store.Connections)
	width := nameLen + 14
	separator := strings.Repeat("─", width)
	dot := styles.Dot.Render("● ")

	if m.showForm {
		w := m.width
		v := m.height
		if w == 0 {
			w = 80
		}
		if v == 0 {
			v = 20
		}
		formStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(w-14).
			Height(v-14).
			Padding(1, 2)
		return formStyle.Render(m.form.View())
	}

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

	leftPanel := styles.Panel.Render(output)
	var rightPanel string
	if len(m.store.Connections) > 0 {
		selected := m.store.Connections[m.cursor]
		rightPanel = detail.Render(selected, m.width/2)
	} else {
		rightPanel = styles.Panel.Render("No connections yet — press 'a' to add one")
	}
	mainView := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
	statusBar := statusbar.Render(m.width)
	return lipgloss.JoinVertical(lipgloss.Left, mainView, statusBar)
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
