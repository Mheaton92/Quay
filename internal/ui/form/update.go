package form

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "tab", "right":
			m.page++
			if m.page > 3 {
				m.page = 3
			}
			return m, m.form.NextGroup()
		case "shift+tab", "left":
			m.page--
			if m.page < 0 {
				m.page = 0
			}
			return m, m.form.PrevGroup()
		case "up":
			return m, m.form.PrevField()
		case "down":
			return m, m.form.NextField()
		case "ctrl+s":
			m.done = true
			return m, nil
		}
	}
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		if m.form.State == huh.StateCompleted {
			m.done = true
		}
	}
	return m, cmd
}
