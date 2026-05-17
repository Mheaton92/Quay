package keys

import (
    "github.com/mheaton92/quay/internal/connection"
    internalkeys "github.com/mheaton92/quay/internal/keys"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/textinput"
)

type Mode int

const (
    ModeView Mode = iota
    ModeGenerate
    ModeDeploy
	ModeDelete
)



type Model struct {
    keys        []internalkeys.Key
    cursor      int
    mode        Mode
    status      string
    err         error
    connections []connection.Connection
    newKeyName    string
    newKeyComment string
    connCursor    int
    nameInput     textinput.Model
    commentInput  textinput.Model
    activeInput   int // 0 = name, 1 = comment
}
func NewKeys(connections []connection.Connection) (*Model, error) {
    keys, err := internalkeys.List()
    ni := textinput.New()
    ni.Placeholder = "Key name (e.g. work_key)"
    ni.Focus()
    ni.CharLimit = 64
    
    ci := textinput.New()
    ci.Placeholder = "e.g. user@host"
    ci.CharLimit = 128
    
    if err != nil {
        return nil, err
    }
    return &Model{
        keys:        keys,
        cursor:      0,
        mode:        ModeView,
        connections: connections,
        nameInput:   textinput.New(),
        commentInput: textinput.New(),
        activeInput:  0,
    }, nil
}
func (m *Model) Done() bool {
    return false
}

func (m *Model) Init() tea.Cmd {
    return nil
}
func (m *Model) IsInView() bool {
    return m.mode == ModeView
}