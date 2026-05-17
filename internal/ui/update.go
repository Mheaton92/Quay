package ui

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/mheaton92/quay/internal/connection"
    "github.com/mheaton92/quay/internal/ui/form"
    "github.com/mheaton92/quay/internal/ui/scp"
    "github.com/mheaton92/quay/internal/ssh"
    "time"
)

type sshExitMsg struct {
    connName string
    err      error
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // ctrl+c always quits no matter what
    if msg, ok := msg.(tea.KeyMsg); ok {
        if msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
    }

    // SCP gets messages before anything else
    if m.showSCP {
        if msg, ok := msg.(tea.KeyMsg); ok {
            if msg.String() == "esc" {
                m.showSCP = false
                return m, nil
            }
        }
        var cmd tea.Cmd
        m.scpModel, cmd = m.scpModel.Update(msg)
        if m.scpModel != nil && m.scpModel.Done() {
            m.showSCP = false
        }
        return m, cmd
    }

    // Form gets messages next
    if m.showForm {
        if msg, ok := msg.(tea.KeyMsg); ok {
            if msg.String() == "esc" {
                m.showForm = false
                return m, nil
            }
        }
        var cmd tea.Cmd
        m.form, cmd = m.form.Update(msg)
        if m.form != nil && m.form.Done() {
            if m.form.IsEditing() {
                m.store.Edit(m.form.OriginalName(), m.form.Connection())
            } else {
                m.store.Add(m.form.Connection())
            }
            m.showForm = false
        }
        return m, cmd
    }

    // Main menu key handling
    if msg, ok := msg.(tea.KeyMsg); ok {
        if msg.String() == "q" {
            return m, tea.Quit
        }
    }

    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    case sshExitMsg:
        for i, conn := range m.store.Connections {
            if conn.Name == msg.connName {
                m.store.Connections[i].LastConnected = time.Now()
                m.store.Connections[i].ConnectionCount++
                m.store.Save()
                break
            }
        }
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
            return m, m.form.Init()
        case "enter":
            if len(m.store.Connections) > 0 {
                selected := m.store.Connections[m.cursor]
                sshCmd := ssh.BuildCmd(selected)
                return m, tea.ExecProcess(sshCmd, func(err error) tea.Msg {
                    return sshExitMsg{connName: selected.Name, err: err}
                })
            }
        case "d":
            if len(m.store.Connections) > 0 {
                m.confirmDelete = true
            }
        case "y":
            if m.confirmDelete {
                selected := m.store.Connections[m.cursor]
                m.store.Delete(selected.Name)
                if m.cursor > 0 {
                    m.cursor--
                }
                m.confirmDelete = false
            }
        case "n":
            m.confirmDelete = false
        case "e":
            if len(m.store.Connections) > 0 {
                selected := m.store.Connections[m.cursor]
                m.form = form.NewForm(selected)
                m.form.SetEditing(true)
                m.showForm = true
                return m, m.form.Init()
            }
        case "s":
            if len(m.store.Connections) > 0 {
                selected := m.store.Connections[m.cursor]
                m.scpModel = scp.NewSCP(selected)
                m.showSCP = true
                return m, m.scpModel.Init()
            }
        }
    }
    return m, nil
}