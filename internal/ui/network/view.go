package network

import (
    "github.com/charmbracelet/lipgloss"
    "github.com/mheaton92/quay/internal/ui/theme"
)

func (m *Model) View() string {
    styles := theme.DefaultStyles()
    var output string

    output += styles.Header.Render("NETWORKING TOOLS") + "\n"
    output += "Target: " + m.conn.Name + " (" + m.conn.Host + ")\n\n"

    if m.showInput {
        output += m.inputLabel + "\n"
        output += m.input.View() + "\n\n"
        output += "[enter] run  [esc] cancel"
        return output
    }

    for i, tool := range m.tools {
        if i == m.cursor {
            output += "▶ " + tool + "\n"
        } else {
            output += "  " + tool + "\n"
        }
    }

    if m.result != "" {
        output += "\n" + m.result
    }

    if m.err != nil {
        output += "\n" + lipgloss.NewStyle().
            Foreground(lipgloss.Color("#f85149")).
            Render("Error: " + m.err.Error())
    }

    output += "\n\n[enter] run  [?] help  [esc] close"

    return output
}
