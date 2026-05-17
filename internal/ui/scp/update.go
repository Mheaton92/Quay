package scp

import (
    tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
    switch m.step {
    case StepDirection:
        if key, ok := msg.(tea.KeyMsg); ok {
            switch key.String() {
            case "u":
                m.upload = true
                m.step = StepFilePicker
                return m, m.fp.Init()
            case "d":
                m.upload = false
                m.step = StepRemotePath
            }
        }
    case StepFilePicker:
        var cmd tea.Cmd
        m.fp, cmd = m.fp.Update(msg)
        if didSelect, path := m.fp.DidSelectFile(msg); didSelect {
            m.localPath = path
            m.step = StepRemotePath
        }
        return m, cmd
    }
    return m, nil
}