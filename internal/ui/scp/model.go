package scp

import (
	"github.com/mheaton92/quay/internal/connection"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/filepicker"
)

type Step int

const (
	StepDirection Step = iota
	StepFilePicker
	StepRemotePath
	StepTransferring
)

type Model struct {
	conn 		connection.Connection
	upload 		bool
	step 		Step
	localPath 	string
	remotePath 	string
	status 		string
	err 		error
	fp 			filepicker.Model
}

func NewSCP(conn connection.Connection) *Model {
    fp := filepicker.New()
    fp.CurrentDirectory = "/"
    return &Model{
        conn: conn,
        step: StepDirection,
        fp:   fp,
    }
}

func (m *Model) Init() tea.Cmd {
    return m.fp.Init()
}
func (m *Model) Done() bool {
	return m.step == StepTransferring && m.status != ""
}