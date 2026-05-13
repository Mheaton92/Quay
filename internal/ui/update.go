package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "j", "down":
            if m.cursor < len(m.store.Connections)-1 {
                m.cursor++
            }
        case "k", "up":
            if m.cursor > 0 {
                m.cursor--
            }
        }
    }
    return m, nil
}