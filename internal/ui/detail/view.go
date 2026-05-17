package detail

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui/theme"
	//"strings"
)

func Render(conn connection.Connection, width int ) string {
	styles := theme.DefaultStyles()

	var output string
	output += styles.Header.Render(conn.Name) + "\n"
	output += fmt.Sprintf("%s@%s port %d", conn.User, conn.Host, conn.Port) + "    Keys: " + conn.IdentityFile + "\n"
	output += fmt.Sprintf("Last: %s\n", conn.LastConnected.Format("2006-01-02 15:04:05") + "    Count: " + fmt.Sprintf("%d", conn.ConnectionCount)) + "\n"

	panelStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#58a6ff")).
		Width(width-10).
		Padding(0, 1)

	return panelStyle.Render(output)
}
