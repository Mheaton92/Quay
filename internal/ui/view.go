package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	ovl "github.com/jsdoublel/bubbletea-overlay"
	"github.com/mheaton92/quay/internal/monitor"
	"github.com/mheaton92/quay/internal/ui/connectionlist"
	"github.com/mheaton92/quay/internal/ui/detail"
	"github.com/mheaton92/quay/internal/ui/keybinds"
	"github.com/mheaton92/quay/internal/ui/statusbar"
	"github.com/mheaton92/quay/internal/ui/theme"
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
	netBarHeight := 5
	netBarHeight += len(m.pinnedHosts) * 5

	styles := theme.DefaultStyles()
	panelHeight := m.height - 5 - netBarHeight
	if panelHeight < 5 {
		panelHeight = 5
	}
	leftPanel := connectionlist.Render(m.store.Connections, m.cursor, panelHeight)

	var rightPanel string
	if len(m.store.Connections) > 0 {
		selected := m.store.Connections[m.cursor]
		if m.width >= 100 {
			// Stacked layout — connection+homelab stacked, active panel wider
			infoWidth := (m.width - lipgloss.Width(leftPanel)) / 3
			activeWidth := (m.width - lipgloss.Width(leftPanel)) - infoWidth - 4

			connectionPanel := detail.RenderConnection(selected, infoWidth, panelHeight/2)
			homelabPanel := detail.RenderHomelab(selected, infoWidth, panelHeight/2)
			stackedInfo := lipgloss.JoinVertical(lipgloss.Left, connectionPanel, homelabPanel)

			activeContent := m.activePanel
			if m.networkModel != nil {
				if m.networkActive {
					activeContent = m.networkModel.View()
				} else {
					activeContent = lipgloss.NewStyle().
						Foreground(lipgloss.Color("#484f58")).
						Render(m.networkModel.View())
				}
			}
			if m.formActive {
				activeContent = m.form.View()
			}
			if m.scpActive {
				activeContent = m.scpModel.View()
			}
			if m.keysActive {
				activeContent = m.keysModel.View()
			}
			if activeContent == "" {
				activeContent = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#484f58")).
					Render("no active tool")
			}
			activeBox := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#58a6ff")).
				Width(activeWidth).
				Height(panelHeight).
				Padding(0, 1).
				Render(activeContent)

			rightPanel = lipgloss.JoinHorizontal(lipgloss.Top, stackedInfo, activeBox)
		} else {
			rightPanel = detail.Render(selected, m.width-lipgloss.Width(leftPanel), panelHeight)
		}
	} else {
		rightPanel = styles.Panel.Copy().Height(panelHeight).Render("No connections yet - press 'a' to add one")
	}

	deleteName := ""
	if m.confirmDelete && len(m.store.Connections) > 0 {
		deleteName = m.store.Connections[m.cursor].Name
	}

	mainView := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
	statusBar := statusbar.Render(m.width, m.confirmDelete, deleteName)
	var liveHost string
	var liveStats *monitor.HostStats
	if len(m.store.Connections) > 0 {
		selected := m.store.Connections[m.cursor]
		alreadyPinned := false
		for _, h := range m.pinnedHosts {
			if h == selected.Host {
				alreadyPinned = true
				break
			}
		}
		if !alreadyPinned {
			liveHost = selected.Host
			liveStats = m.monitor.Stats(selected.Host)
		}
	}

	pinnedStats := make(map[string]*monitor.HostStats)
	for _, host := range m.pinnedHosts {
		pinnedStats[host] = m.monitor.Stats(host)
	}

	netBar := statusbar.RenderNetBar(m.pinnedHosts, pinnedStats, liveHost, liveStats, m.width)

	fullView := lipgloss.JoinVertical(
		lipgloss.Left,
		mainView,
		netBar,
		statusBar,
	)

	// Build overlay if any panel is active
	var overlay string
	if m.showHelp {
		overlay = keybinds.Render(m.width, m.keybinds)
	} else if m.showForm {
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(m.width-14).
			Height(m.height-14).
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
	} else if m.showNetwork {
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(60).
			Height(20).
			Padding(1, 2).
			Render(m.networkModel.View())
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
