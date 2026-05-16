package theme 

import (
	"github.com/charmbracelet/lipgloss"
)

// Colors
var (
	ColorBlue   = lipgloss.Color("#58a6ff")
	ColorGreen  = lipgloss.Color("#3fb950")
	ColorYellow = lipgloss.Color("#f0c674")
	ColorRed    = lipgloss.Color("#e5534b")
	ColorDim	= lipgloss.Color("#484f58")
	ColorWhite  = lipgloss.Color("#c9d1d9")
)

// Styles
type Styles struct {
	Header lipgloss.Style
	Separator lipgloss.Style
	Selected lipgloss.Style
	Dot lipgloss.Style
	Panel lipgloss.Style
}

func DefaultStyles() Styles {
	return Styles{
		Header: lipgloss.NewStyle().Bold(true).Foreground(ColorBlue),
		Separator: lipgloss.NewStyle().Foreground(ColorDim),
		Selected: lipgloss.NewStyle().Bold(true).Foreground(ColorBlue),
		Dot: lipgloss.NewStyle().Foreground(ColorGreen),
		Panel: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBlue).
			Padding(0, 1),
	}
}