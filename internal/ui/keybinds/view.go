package keybinds

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/config"
)

func Render(width int, keybinds config.Keybinds) string {
	var output string

	for _, e := range getData(keybinds) {
		if e.subtitle != "" {
			output += "\n" + lipgloss.NewStyle().
				Foreground(lipgloss.Color("#d29922")).
				PaddingLeft(12).
				Render(e.subtitle) + "\n\n"
		} else {
			key := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#58a6ff")).
				Width(10).
				Align(lipgloss.Right).
				Render(e.hotkey)
			output += key + "  " + e.description + "\n"
		}
	}

	panelStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#58a6ff")).
		Padding(0, 1)

	return panelStyle.Render(output)
}

type entry struct {
	hotkey      string
	description string
	subtitle    string
}

func getData(keybinds config.Keybinds) []entry {
	data := []entry{
		{subtitle: "General"},
		{
			hotkey:      keybinds.Connect,
			description: "Connect to selected host",
		},
		{
			hotkey:      keybinds.Add,
			description: "Add new host",
		},
		{
			hotkey:      keybinds.Edit,
			description: "Edit selected host",
		},
		{
			hotkey:      keybinds.Delete,
			description: "Delete selected host",
		},
		{
			hotkey:      keybinds.SCP,
			description: "Open SCP menu for selected host",
		},
		{
			hotkey:      keybinds.Keys,
			description: "Open SSH key menu",
		},
		{
			hotkey:      keybinds.Help,
			description: "Open this help menu",
		},
		{
			hotkey:      keybinds.Up,
			description: "up",
		},
		{
			hotkey:      keybinds.Down,
			description: "Down",
		},
		{
			hotkey:      keybinds.NextTab,
			description: "Move to next tab in menus",
		},
		{
			hotkey:      keybinds.PrevTab,
			description: "Move back a tab in menus",
		},
		{
			hotkey:      keybinds.Networking,
			description: "Open networking tools",
		},
		{
			hotkey:      keybinds.PinSession,
			description: "Pin current host to bottom bar (session)",
		},
		{
			hotkey:      keybinds.PinPersistent,
			description: "Pin current host to bottom bar (persistent)",
		},
	}
	return data
}
