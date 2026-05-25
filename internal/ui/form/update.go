package form

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"strconv"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "tab", "right":
			m.page++
			if m.page > 3 {
				m.page = 3
			}
			m.field = []int{0, 6, 12, 15}[m.page]
			return m, m.form.NextGroup()
		case "shift+tab", "left":
			m.page--
			if m.page < 0 {
				m.page = 0
			}
			m.field = []int{0, 6, 12, 15}[m.page]
			return m, m.form.PrevGroup()
		case "?":
			m.showFieldHelp = !m.showFieldHelp
			return m, nil
		case "up":
			if m.field > 0 {
				m.field--
			}
			return m, m.form.PrevField()
		case "down":
			if m.field < len(fieldNames)-1 {
				m.field++
			}
			return m, m.form.NextField()
		case "ctrl+s":
			if m.conn.Name == "" {
				m.validationError = "Name is required"
				return m, nil
			}
			if m.conn.Host == "" {
				m.validationError = "Host is required"
				return m, nil
			}
			port, _ := strconv.Atoi(m.portStr)
			if port < 1 || port > 65535 {
				m.validationError = "Port must be between 1 and 65535"
				return m, nil
			}
			m.validationError = ""
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
