package form

import (
	"github.com/charmbracelet/huh"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    form, cmd := m.form.Update(msg)
    if f, ok := form.(*huh.Form); ok {
        m.form = f
        if m.form.State == huh.StateCompleted {
            m.done = true
        }
    }
    return m, cmd
}