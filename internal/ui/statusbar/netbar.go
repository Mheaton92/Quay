package statusbar

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/mheaton92/quay/internal/monitor"
	"strings"
)

func renderHostLine(label string, host string, stats *monitor.HostStats, width int) string {
	if stats == nil {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("#484f58")).
			Render(label + "  " + host + "  waiting...")
	}

	sparkWidth := width - 60
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

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#58a6ff")).
		Bold(true).
		Width(8)

	hostStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c9d1d9")).
		Width(14)

	latencyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#58a6ff")).
		Width(10)

	lossStyle := lipgloss.NewStyle().
		Foreground(lossColor).
		Width(10)

	sparkStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#58a6ff"))

	return labelStyle.Render(label) +
		hostStyle.Render(host) +
		latencyStyle.Render(latency) +
		lossStyle.Render(fmt.Sprintf("%.0f%% loss", stats.PacketLoss)) +
		sparkStyle.Render(spark)
}

func RenderNetBar(pinnedHosts []string, pinnedStats map[string]*monitor.HostStats, liveHost string, liveStats *monitor.HostStats, width int) string {
	separator := strings.Repeat("─", width-4)
	var lines []string

	for _, host := range pinnedHosts {
		lines = append(lines, renderHostLine("PINNED", host, pinnedStats[host], width))
	}

	if liveHost != "" {
		if len(lines) > 0 {
			lines = append(lines, separator)
		}
		lines = append(lines, renderHostLine("LIVE", liveHost, liveStats, width))
	}

	content := strings.Join(lines, "\n")

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#58a6ff")).
		Width(width-2).
		Padding(0, 1).
		Render(content)
}
