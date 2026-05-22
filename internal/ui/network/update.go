package network

import (
    tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
    if key, ok := msg.(tea.KeyMsg); ok {
        switch key.String() {
        case "j", "down":
            if m.cursor < len(m.tools)-1 {
                m.cursor++
            }
        case "k", "up":
            if m.cursor > 0 {
                m.cursor--
            }
        case "?":
            m.showHelp = !m.showHelp
        case "enter":
            return m, m.runTool()
        }
    }
    return m, nil
}

func (m *Model) runTool() tea.Cmd {
    return func() tea.Msg {
        return nil
    }
}
