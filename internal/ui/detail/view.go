package detail

import (
    "fmt"
    "os"
    "strings"
    "github.com/charmbracelet/lipgloss"
    "github.com/mheaton92/quay/internal/connection"
    "github.com/mheaton92/quay/internal/ui/theme"
)

func Render(conn connection.Connection, width int, height int) string {
    styles := theme.DefaultStyles()
    header := styles.Header.Render(conn.Name) + "  " +
        lipgloss.NewStyle().Foreground(lipgloss.Color("#484f58")).
            Render(conn.User+"@"+conn.Host+"  port "+fmt.Sprintf("%d", conn.Port))

    leftWidth := width/2 - 4
    rightWidth := width/2 - 4

    leftPanel := lipgloss.NewStyle().
        Width(leftWidth).
        Height(height - 4).
        Border(lipgloss.NormalBorder(), false, true, false, false).
        BorderForeground(lipgloss.Color("#21262d")).
        Padding(0, 1).
        Render(connectionContent(conn))

    rightPanel := lipgloss.NewStyle().
        Width(rightWidth).
        Height(height - 4).
        Padding(0, 1).
        Render(homelabContent())

    columns := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

    return styles.Panel.Copy().
        Width(width - 10).
        Height(height).
        Render(header + "\n\n" + columns)
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
    var content string
    content += sectionTitle("CONNECTION") + "\n"
    content += row("host", conn.Host)
    content += row("user", conn.User)
    content += row("port", fmt.Sprintf("%d", conn.Port))
    content += row("key", shortenPath(conn.IdentityFile))
    if conn.ProxyJump != "" {
        content += row("proxy", conn.ProxyJump)
    }
    content += "\n"
    content += sectionTitle("HISTORY") + "\n"
    content += row("last", conn.LastConnected.Format("2006-01-02 15:04"))
    content += row("count", fmt.Sprintf("%d", conn.ConnectionCount))
    if len(conn.Tags) > 0 {
        content += "\n" + sectionTitle("TAGS") + "\n"
        content += "  " + strings.Join(conn.Tags, ", ") + "\n"
    }
    if conn.Notes != "" {
        content += "\n" + sectionTitle("NOTES") + "\n"
        content += "  " + conn.Notes + "\n"
    }
    return content
}

func homelabContent() string {
    var content string
    content += sectionTitle("HOMELAB") + "\n"
    content += lipgloss.NewStyle().
        Foreground(lipgloss.Color("#484f58")).
        Render("  no integration configured") + "\n"
    return content
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
