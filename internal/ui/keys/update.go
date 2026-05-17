package keys

import (
    "os"
    "os/exec"
    "path/filepath"
    tea "github.com/charmbracelet/bubbletea"
    internalkeys "github.com/mheaton92/quay/internal/keys"
)

type keyGenDoneMsg struct {
    err  error
    name string
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
    switch m.mode {
    case ModeView:
        switch msg := msg.(type) {
        case keyGenDoneMsg:
            m.keys, _ = internalkeys.List()
            if msg.err != nil {
                m.status = "Error: " + msg.err.Error()
            } else {
                m.status = "Key generated: " + msg.name
            }
        case tea.KeyMsg:
            switch msg.String() {
            case "k", "up":
                if m.cursor > 0 {
                    m.cursor--
                }
            case "j", "down":
                if m.cursor < len(m.keys)-1 {
                    m.cursor++
                }
            case "g":
                m.mode = ModeGenerate
                m.activeInput = 0
                m.nameInput.Focus()
                m.commentInput.Blur()
                m.nameInput.SetValue("")
                m.commentInput.SetValue("")
            case "d":
                m.mode = ModeDeploy
                m.connCursor = 0
            case "c":
                if len(m.keys) > 0 {
                    internalkeys.CopyToClipboard(m.keys[m.cursor])
                    m.status = "Public key copied to clipboard!"
                }
            case "x":
                if len(m.keys) > 0 {
                    m.mode = ModeDelete
                }
            }
        }

    case ModeGenerate:
        var cmd tea.Cmd
        if m.activeInput == 0 {
            m.nameInput, cmd = m.nameInput.Update(msg)
        } else {
            m.commentInput, cmd = m.commentInput.Update(msg)
        }
        if key, ok := msg.(tea.KeyMsg); ok {
            switch key.String() {
            case "tab":
                if m.activeInput == 0 {
                    m.activeInput = 1
                    m.nameInput.Blur()
                    m.commentInput.Focus()
                } else {
                    m.activeInput = 0
                    m.commentInput.Blur()
                    m.nameInput.Focus()
                }
            case "enter":
                m.newKeyName = m.nameInput.Value()
                m.newKeyComment = m.commentInput.Value()
                if m.newKeyName != "" {
                    m.mode = ModeView
                    sshCmd := exec.Command("ssh-keygen",
                        "-t", "ed25519",
                        "-f", filepath.Join(os.Getenv("HOME"), ".ssh", m.newKeyName),
                        "-C", m.newKeyComment,
                        "-N", "")
                    name := m.newKeyName
                    return m, tea.ExecProcess(sshCmd, func(err error) tea.Msg {
                        return keyGenDoneMsg{err: err, name: name}
                    })
                }
            case "esc":
                m.mode = ModeView
            }
        }
        return m, cmd

    case ModeDeploy:
        if key, ok := msg.(tea.KeyMsg); ok {
            switch key.String() {
            case "j", "down":
                if m.connCursor < len(m.connections)-1 {
                    m.connCursor++
                }
            case "k", "up":
                if m.connCursor > 0 {
                    m.connCursor--
                }
            case "enter":
                if len(m.connections) > 0 {
                    conn := m.connections[m.connCursor]
                    err := internalkeys.Deploy(m.keys[m.cursor], conn.Host, conn.User, conn.Port)
                    if err != nil {
                        m.status = "Deploy failed: " + err.Error()
                    } else {
                        m.status = "Key deployed!"
                    }
                    m.mode = ModeView
                }
            case "esc":
                m.mode = ModeView
            }
        }

    case ModeDelete:
        if key, ok := msg.(tea.KeyMsg); ok {
            switch key.String() {
            case "y":
                err := internalkeys.Delete(m.keys[m.cursor])
                if err != nil {
                    m.status = "Delete failed: " + err.Error()
                } else {
                    m.status = "Key deleted!"
                    m.keys, _ = internalkeys.List()
                    if m.cursor > 0 {
                        m.cursor--
                    }
                }
                m.mode = ModeView
            case "n", "esc":
                m.mode = ModeView
            }
        }
    }
    return m, nil
}