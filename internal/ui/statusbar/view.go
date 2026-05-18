package statusbar

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func Render(width int, confirmDelete bool, deleteName string) string {
	separator := strings.Repeat("─", width)

	if confirmDelete {
		confirmMsg := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f85149")).
			Render("Delete \"" + deleteName + "\"? [y] yes  [n] no")
		return separator + "\n" + confirmMsg
	}

	keys := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#484f58")).
		Render("[enter] connect  [a] add  [e] edit  [d] delete [s] scp [j/k] navigate  [q] quit")

	return separator + "\n" + keys
}
