package scp

import (
	"github.com/mheaton92/quay/internal/ui/theme"
)

func (m *Model) View() string {
	styles := theme.DefaultStyles()
	switch m.step {
	case StepDirection:
		return styles.Header.Render("SCP — "+m.conn.Name) + "\n\n" +
			"[u] Upload   local → remote\n" +
			"[d] Download remote → local\n\n" +
			"[esc] Cancel"
	case StepFilePicker:
		return m.fp.View()
	case StepRemotePath:
		return "── Remote path ──\n\n" +
			m.remoteInput.View() + "\n\n" +
			"[enter] confirm  [esc] cancel"
	case StepLocalPath:
		return "── Local save path ──\n\n" +
			m.localInput.View() + "\n\n" +
			"[enter] confirm  [esc] cancel"
	case StepTransferring:
		if m.err != nil {
			return styles.Header.Render("✗ Transfer Failed") + "\n\n" +
				m.err.Error() + "\n\n[esc] close"
		}
		if m.upload {
			return styles.Header.Render("✓ "+m.status) + "\n\n" +
				"Uploaded to: " + m.conn.User + "@" + m.conn.Host + ":" + m.remotePath + "\n\n[esc] close"
		}
		return styles.Header.Render("✓ "+m.status) + "\n\n" +
			"Saved to: " + m.localPath + "\n\n[esc] close"
	}
	return ""
}
