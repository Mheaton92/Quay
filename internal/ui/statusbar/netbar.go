package statusbar

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
    "github.com/mheaton92/quay/internal/monitor"
)

func RenderNetBar(host string, stats *monitor.HostStats, width int) string {
    if stats == nil {
        return lipgloss.NewStyle().
            Border(lipgloss.RoundedBorder()).
            BorderForeground(lipgloss.Color("#58a6ff")).
            Width(width - 2).
            Height(3).
            Padding(0, 1).
            Render("Waiting for stats...")
    }

    sparkWidth := width - 50
    if sparkWidth < 10 {
        sparkWidth = 10
    }

    spark := monitor.Sparkline(stats.History, sparkWidth)

    var latency string
    if stats.Online {
        ms := float64(stats.Latency.Microseconds()) / 1000.0
        if ms < 1 {
            latency = fmt.Sprintf("%.2fms", ms)
        } else {
            latency = fmt.Sprintf("%.0fms", ms)
        }
    } else {
        latency = "offline"
    }

    lossColor := lipgloss.Color("#3fb950")
    if stats.PacketLoss > 0 {
        lossColor = lipgloss.Color("#f85149")
    }

    header := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#58a6ff")).
        Bold(true).
        Render("● LIVE  " + host)

    statsLine := fmt.Sprintf("  latency: %-8s  loss: %-8s  %s",
        latency,
        fmt.Sprintf("%.0f%%", stats.PacketLoss),
        spark)

    content := header + "\n" +
        lipgloss.NewStyle().Foreground(lossColor).Render(statsLine)

    return lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#58a6ff")).
        Width(width - 2).
        Height(3).
        Padding(0, 1).
        Render(content)
}