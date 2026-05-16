package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ui/form"
	"github.com/mheaton92/quay/internal/ssh"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle escpae/quit at the top level
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			if m.showForm {
				m.showForm = false
				return m, nil
			}
			return m, tea.Quit
		}
	}
	if m.showForm {
		var cmd tea.Cmd
		m.form, cmd = m.form.Update(msg)
		if m.form != nil && m.form.Done() {
			m.store.Add(m.form.Connection())
			m.showForm = false
		}
		return m, cmd
	}

	switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor < len(m.store.Connections)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "a":
			m.form = form.NewForm(connection.Connection{})
			m.showForm = true

		case "enter":
			if len(m.store.Connections) > 0 {
				selected := m.store.Connections[m.cursor]
				err := ssh.Connect(selected)
				if err != nil {
					m.err = err
				}
			}

		case "d":
			if len(m.store.Connections) > 0 {
				selected := m.store.Connections[m.cursor]
				m.store.Delete(selected.Name)
				if m.cursor > 0 {
					m.cursor--
				}
			}
		}
		
	}
	return m, nil
}
