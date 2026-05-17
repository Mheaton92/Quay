package statusbar

import (
    "github.com/charmbracelet/lipgloss"
	"strings"
)

func Render(width int) string {
    separator := strings.Repeat("─", width)
    
    keys := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#484f58")).
        Render("[enter] connect  [a] add  [e] edit  [d] delete  [j/k] navigate  [q] quit")
    
    return separator + "\n" + keys
}