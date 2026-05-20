package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/ui/connectionlist"
	"github.com/mheaton92/quay/internal/ui/detail"
	"github.com/mheaton92/quay/internal/ui/keybinds"
	"github.com/mheaton92/quay/internal/ui/statusbar"
	"github.com/mheaton92/quay/internal/ui/theme"
	ovl "github.com/jsdoublel/bubbletea-overlay"
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
	panelHeight := m.height - 5
	leftPanel := connectionlist.Render(m.store.Connections, m.cursor, panelHeight)

	var rightPanel string
	if len(m.store.Connections) > 0 {
		selected := m.store.Connections[m.cursor]
		rightPanel = detail.Render(selected, m.width/2, panelHeight)
	} else {
		rightPanel = styles.Panel.Copy().Height(panelHeight).Render("No connections yet — press 'a' to add one")
	}

	deleteName := ""
	if m.confirmDelete && len(m.store.Connections) > 0 {
		deleteName = m.store.Connections[m.cursor].Name
	}

	mainView := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
	statusBar := statusbar.Render(m.width, m.confirmDelete, deleteName)
	fullView := lipgloss.JoinVertical(lipgloss.Left, mainView, statusBar)

	// Build overlay if any panel is active
	var overlay string
	if m.showHelp {
		overlay = keybinds.Render(m.width, m.keybinds)
	} else if m.showForm {
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(m.width - 14).
			Height(m.height - 14).
			Padding(1, 2).
			Render(m.form.View())
	} else if m.showSCP {
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(60).
			Height(15).
			Padding(1, 2).
			Render(m.scpModel.View())
	} else if m.showKeys {
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(60).
			Height(15).
			Padding(1, 2).
			Render(m.keysModel.View())
	}

		if overlay != "" {
			return ovl.Composite(
				overlay,
				fullView,
				ovl.Center,
				ovl.Center,
				0, 0,
			)
		}
	

	return fullView
}