package form

import "github.com/charmbracelet/lipgloss"

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

	return tabBar + "\n\n" + view
}
