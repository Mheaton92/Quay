package detail

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui/theme"
)

// Render draws connection info and homelab stacked in a single panel.
// width is the inner content width; the panel border adds 2 to the rendered size.
func Render(conn connection.Connection, width int, height int) string {
	styles := theme.DefaultStyles()
	content := connectionContent(conn) + "\n" + homelabContent()
	return styles.Panel.Copy().
		Width(width).
		Height(height).
		Padding(0, 1).
		Render(content)
}

func RenderConnection(conn connection.Connection, width, height int) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true, true, false, true).
		BorderForeground(lipgloss.Color("#58a6ff")).
		Width(width).
		Height(height).
		Padding(0, 1).
		Render(connectionContent(conn))
}

func RenderHomelab(conn connection.Connection, width, height int) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), false, true, true, true).
		BorderForeground(lipgloss.Color("#58a6ff")).
		Width(width).
		Height(height).
		Padding(0, 1).
		Render(homelabContent())
}

func connectionContent(conn connection.Connection) string {
	var b strings.Builder
	b.WriteString(sectionTitle("CONNECTION") + "\n")
	b.WriteString(row("host", conn.Host))
	b.WriteString(row("user", conn.User))
	b.WriteString(row("port", fmt.Sprintf("%d", conn.Port)))
	b.WriteString(row("key", shortenPath(conn.IdentityFile)))
	if conn.ProxyJump != "" {
		b.WriteString(row("proxy", conn.ProxyJump))
	}
	b.WriteString("\n")
	b.WriteString(sectionTitle("HISTORY") + "\n")
	b.WriteString(row("last", conn.LastConnected.Format("2006-01-02 15:04")))
	b.WriteString(row("count", fmt.Sprintf("%d", conn.ConnectionCount)))
	if len(conn.Tags) > 0 {
		b.WriteString("\n" + sectionTitle("TAGS") + "\n")
		b.WriteString("  " + strings.Join(conn.Tags, ", ") + "\n")
	}
	if conn.Notes != "" {
		b.WriteString("\n" + sectionTitle("NOTES") + "\n")
		b.WriteString("  " + conn.Notes + "\n")
	}
	return b.String()
}

func homelabContent() string {
	var b strings.Builder
	b.WriteString(sectionTitle("HOMELAB") + "\n")
	b.WriteString(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#484f58")).
		Render("  no integration configured") + "\n")
	return b.String()
}

func sectionTitle(title string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#484f58")).
		Render(title)
}

func row(label, value string) string {
	return fmt.Sprintf("  %-8s %s\n",
		lipgloss.NewStyle().Foreground(lipgloss.Color("#484f58")).Render(label),
		value)
}

func shortenPath(path string) string {
	home := os.Getenv("HOME")
	if home != "" && strings.HasPrefix(path, home) {
		return "~" + path[len(home):]
	}
	termuxHome := "/data/data/com.termux/files/home"
	if strings.HasPrefix(path, termuxHome) {
		return "~" + path[len(termuxHome):]
	}
	return path
}
