package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mheaton92/quay/internal/connection"
	"github.com/mheaton92/quay/internal/ssh"
	"github.com/mheaton92/quay/internal/ui/form"
	uikeys "github.com/mheaton92/quay/internal/ui/keys"
	"github.com/mheaton92/quay/internal/ui/scp"
)

type sshExitMsg struct {
	connName string
	err      error
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// ctrl+c always quits
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "esc" {
			if m.showHelp {
				m.showHelp = false
				return m, nil
			}
			if m.showSCP {
				m.showSCP = false
				return m, nil
			}
			if m.showForm {
				m.showForm = false
				return m, nil
			}
			if m.showKeys && m.keysModel.IsInView() {
				m.showKeys = false
				return m, nil
			}
		}
		if msg.String() == m.keybinds.Quit {
			if m.showHelp {
				m.showHelp = false
				return m, nil
			}
			if !m.showForm && !m.showSCP && !m.showKeys {
				return m, tea.Quit
			}
		}
	}

	// SCP gets messages
	if m.showSCP {
		var cmd tea.Cmd
		m.scpModel, cmd = m.scpModel.Update(msg)
		if m.scpModel != nil && m.scpModel.Done() {
			m.showSCP = false
		}
		return m, cmd
	}

	// Form gets messages
	if m.showForm {
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

	// Keys gets messages
	if m.showKeys {
		var cmd tea.Cmd
		m.keysModel, cmd = m.keysModel.Update(msg)
		return m, cmd
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
		key := msg.String()
		if key == m.keybinds.Down || key == "down" {
			if m.cursor < len(m.store.Connections)-1 {
				m.cursor++
			}
		} else if key == m.keybinds.Up || key == "up" {
			if m.cursor > 0 {
				m.cursor--
			}
		} else if key == m.keybinds.Add {
			m.form = form.NewForm(connection.Connection{})
			m.showForm = true
			return m, m.form.Init()
		} else if key == m.keybinds.Connect {
			if len(m.store.Connections) > 0 {
				selected := m.store.Connections[m.cursor]
				sshCmd := ssh.BuildCmd(selected)
				return m, tea.ExecProcess(sshCmd, func(err error) tea.Msg {
					return sshExitMsg{connName: selected.Name, err: err}
				})
			}
		} else if key == m.keybinds.Delete {
			if len(m.store.Connections) > 0 {
				m.confirmDelete = true
			}
		} else if key == "y" {
			if m.confirmDelete {
				selected := m.store.Connections[m.cursor]
				m.store.Delete(selected.Name)
				if m.cursor > 0 {
					m.cursor--
				}
				m.confirmDelete = false
			}
		} else if key == "n" {
			m.confirmDelete = false
		} else if key == m.keybinds.Edit {
			if len(m.store.Connections) > 0 {
				selected := m.store.Connections[m.cursor]
				m.form = form.NewForm(selected)
				m.form.SetEditing(true)
				m.showForm = true
				return m, m.form.Init()
			}
		} else if key == m.keybinds.SCP {
			if len(m.store.Connections) > 0 {
				selected := m.store.Connections[m.cursor]
				m.scpModel = scp.NewSCP(selected)
				m.showSCP = true
				return m, m.scpModel.Init()
			}
		} else if key == m.keybinds.Keys {
			keysModel, err := uikeys.NewKeys(m.store.Connections)
			if err != nil {
				m.err = err
			} else {
				m.keysModel = keysModel
				m.showKeys = true
			}
		} else if key == m.keybinds.Help {
			m.showHelp = !m.showHelp
		}
	}
	return m, nil
}