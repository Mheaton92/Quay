package ui

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
    "github.com/mheaton92/quay/internal/ui/connectionlist"
    "github.com/mheaton92/quay/internal/ui/detail"
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

    if m.showForm {
        w := m.width
        v := m.height
        if w == 0 { w = 80 }
        if v == 0 { v = 20 }
        formStyle := lipgloss.NewStyle().
            Border(lipgloss.RoundedBorder()).
            BorderForeground(lipgloss.Color("#58a6ff")).
            Width(w - 14).
            Height(v - 14).
            Padding(1, 2)
        return formStyle.Render(m.form.View())
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

    mainView := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
    
    deleteName := ""
    if m.confirmDelete && len(m.store.Connections) > 0 {
        deleteName = m.store.Connections[m.cursor].Name
    }
    statusBar := statusbar.Render(m.width, m.confirmDelete, deleteName)
    return lipgloss.JoinVertical(lipgloss.Left, mainView, statusBar)
}