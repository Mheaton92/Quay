package detail

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui/theme"
	"strings"
)

func Render(conn connection.Connection, width int ) string {
	styles := theme.DefaultStyles()

	var output string
	output += styles.Header.Render(conn.Name) + "\n"
	output += fmt.Sprintf("%s@%s port %d", conn.User, conn.Host, conn.Port) + "\n"
	output += "Key: " + conn.IdentityFile + "\n"
	output += "Tags: " + strings.Join(conn.Tags, ", ") + "\n"

	panelStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#58a6ff")).
		Width(width-10).
		Padding(0, 1)

	return panelStyle.Render(output)
}
