package scp

import (
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mheaton92/quay/internal/connection"
	"os"
)

type Step int

const (
	StepDirection Step = iota
	StepFilePicker
	StepRemotePath
	StepLocalPath
	StepTransferring
)

type Model struct {
	conn        connection.Connection
	upload      bool
	step        Step
	localPath   string
	remotePath  string
	status      string
	err         error
	fp          filepicker.Model
	remoteInput textinput.Model
	localInput  textinput.Model
}

func NewSCP(conn connection.Connection) *Model {
	fp := filepicker.New()
	fp.CurrentDirectory = os.Getenv("HOME")
	fp.Height = 20
	ti := textinput.New()
	ti.Placeholder = "Enter remote path (e.g., /home/user/file.txt)"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50
	li := textinput.New()
	li.Placeholder = "/home/user/"
	li.Focus()
	li.CharLimit = 256
	li.Width = 50
	return &Model{
		conn:        conn,
		step:        StepDirection,
		fp:          fp,
		remoteInput: ti,
		localInput:  li,
	}
}

func (m *Model) Init() tea.Cmd {
	return m.fp.Init()
}
func (m *Model) Done() bool {
	return false // never auto-close, user must press esc
}
