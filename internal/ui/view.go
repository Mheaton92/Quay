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

func (m Model) anyToolActive() bool {
	return m.confirmImport || m.networkActive || m.formActive || m.scpActive || m.keysActive
}

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
	leftWidth := lipgloss.Width(leftPanel)

	var rightPanel string

	if m.width >= 100 {
		totalRight := m.width - leftWidth

		if m.anyToolActive() {
			// tool panel
			middleInner := (totalRight - 4) * 2 / 5
			var middlePanel string
			if len(m.store.Connections) > 0 {
				selected := m.store.Connections[m.cursor]
				middlePanel = detail.Render(selected, middleInner, panelHeight)
			} else {
				middlePanel = styles.Panel.Copy().Width(middleInner).Height(panelHeight).Render("")
			}
			toolInner := totalRight - lipgloss.Width(middlePanel) - 2

			var toolContent string
			switch {
			case m.confirmImport:
				var lines string
				for _, c := range m.pendingImport {
					lines += fmt.Sprintf("  %-20s %s\n", c.Name, c.Host)
				}
				toolContent = "IMPORT FROM ~/.ssh/config\n\n" + lines + "\n[y] import  [esc] cancel"
			case m.networkActive && m.networkModel != nil:
				toolContent = m.networkModel.View()
			case m.formActive && m.form != nil:
				toolContent = m.form.View()
			case m.scpActive && m.scpModel != nil:
				toolContent = m.scpModel.View()
			case m.keysActive && m.keysModel != nil:
				toolContent = m.keysModel.View()
			}

			toolBox := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#58a6ff")).
				Width(toolInner).
				Height(panelHeight).
				Padding(0, 1).
				Render(toolContent)

			rightPanel = lipgloss.JoinHorizontal(lipgloss.Top, middlePanel, toolBox)
		} else if len(m.store.Connections) > 0 {
			selected := m.store.Connections[m.cursor]
			rightPanel = detail.Render(selected, totalRight-2, panelHeight)
		} else {
			rightPanel = styles.Panel.Copy().Width(totalRight - 2).Height(panelHeight).
				Render("No connections yet — press 'a' to add one")
		}
	} else {
		if len(m.store.Connections) > 0 {
			selected := m.store.Connections[m.cursor]
			rightPanel = detail.Render(selected, m.width-leftWidth-2, panelHeight)
		} else {
			rightPanel = styles.Panel.Copy().Width(m.width - leftWidth - 2).Height(panelHeight).
				Render("No connections yet — press 'a' to add one")
		}
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

	fullView := lipgloss.JoinVertical(lipgloss.Left, mainView, netBar, statusBar)

	// -- overlays (narrow screen)
	var overlay string
	switch {
	case m.showHelp:
		overlay = keybinds.Render(m.width, m.keybinds)
	case m.showForm:
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(m.width-14).
			Height(m.height-14).
			Padding(1, 2).
			Render(m.form.View())
	case m.showSCP:
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(60).
			Height(15).
			Padding(1, 2).
			Render(m.scpModel.View())
	case m.showKeys:
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(60).
			Height(15).
			Padding(1, 2).
			Render(m.keysModel.View())
	case m.showNetwork:
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(60).
			Height(20).
			Padding(1, 2).
			Render(m.networkModel.View())
	case m.showImport:
		var lines string
		for _, c := range m.pendingImport {
			lines += fmt.Sprintf("  %-20s %s\n", c.Name, c.Host)
		}
		overlay = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#58a6ff")).
			Width(50).
			Height(len(m.pendingImport)+6).
			Padding(1, 2).
			Render("IMPORT FROM ~/.ssh/config\n\n" + lines + "\n[y] import  [esc] cancel")
	}

	if overlay != "" {
		return ovl.Composite(overlay, fullView, ovl.Center, ovl.Center, 0, 0)
	}

	return fullView
}
