package keys

import (
	"fmt"
	"github.com/mheaton92/quay/internal/ui/theme"
)

func (m *Model) View() string {
	styles := theme.DefaultStyles()
	var output string
	switch m.mode {
	case ModeView:
		output += styles.Header.Render("SSH KEYS") + "\n\n"
		if len(m.keys) == 0 {
			output += "No keys found in ~/.ssh/\n"
		}
		for i, key := range m.keys {
			if i == m.cursor {
				output += "▶ "
			} else {
				output += "  "
			}
			output += fmt.Sprintf("%-20s %s\n", key.Name, key.Type)
		}
		output += "\n[g] generate  [d] deploy  [c] copy  [x] delete  [esc] close"

	case ModeGenerate:
		output += styles.Header.Render("GENERATE NEW KEY") + "\n\n"
		output += "Type: ed25519 (recommended)\n\n"
		output += "Name:\n" + m.nameInput.View() + "\n\n"
		output += "Comment:\n" + m.commentInput.View() + "\n\n"
		output += "[tab] switch field  [enter] generate  [esc] back"

	case ModeDeploy:
		output += styles.Header.Render("DEPLOY KEY") + "\n\n"
		output += "Key: " + m.keys[m.cursor].Name + "\n\n"
		for i, conn := range m.connections {
			if i == m.connCursor {
				output += "▶ "
			} else {
				output += "  "
			}
			output += conn.Name + "\n"
		}
		output += "\n[enter] deploy  [esc] cancel"

	case ModeDelete:
		output += styles.Header.Render("DELETE KEY") + "\n\n"
		output += "Delete: " + m.keys[m.cursor].Name + "?\n\n"
		output += "[y] yes  [n] no"
	}
	return output
}
