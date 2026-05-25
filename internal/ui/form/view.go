package form

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/ui/help"
	"github.com/mheaton92/quay/internal/ui/theme"
)

func (m *Model) View() string {
	tabs := []string{"Basic", "Connection", "Forwarding", "Meta"}
	current := m.page
	var tabBar string
	for i, tab := range tabs {
		if i == current {
			tabBar += "[" + tab + "]"
		} else {
			tabBar += " " + tab + " "
		}
	}

	view := m.form.View()
	if m.validationError != "" {
		view += "\n\n" + lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f85149")).
			Render("✗ "+m.validationError)
	}

	return tabBar + "\n\n" + view + "\n\n" + m.renderFieldHelp()
}

func (m *Model) renderFieldHelp() string {
    if !m.showFieldHelp {
        return ""
    }
    h := help.GetFieldHelp(fieldNames[m.field])
    if h == nil {
        return ""
    }
    return lipgloss.NewStyle().
				Foreground(theme.ColorBlue).
				Bold(true).
				Render(h.Field) + ":  " +
				lipgloss.NewStyle().
				Foreground(theme.ColorWhite).
				Render(h.Description)
}
