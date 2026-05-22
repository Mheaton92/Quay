package form

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/mheaton92/quay/internal/connection"
	"os"
	"strconv"
)

type Model struct {
	conn            connection.Connection
	portStr         string // temporary string for port input
	tagsStr         string // temporary string for tags input
	form            *huh.Form
	done            bool
	editing         bool // true if editing an existing connection, false if creating new
	page            int  // current page index
	originalName    string
	validationError string
}

func (m *Model) Init() tea.Cmd {
	return m.form.Init()
}

func NewForm(conn connection.Connection) *Model {
	m := &Model{
		conn: conn,
		portStr: func() string {
			if conn.Port == 0 {
				return "22"
			}
			return fmt.Sprintf("%d", conn.Port)
		}(),
		tagsStr: "",
	}
	m.originalName = conn.Name
	defaultKey := ""
	if _, err := os.Stat(os.Getenv("HOME") + "/.ssh/id_ed25519"); err == nil {
		defaultKey = os.Getenv("HOME") + "/.ssh/id_ed25519"
	}

	if m.conn.IdentityFile == "" {
		m.conn.IdentityFile = defaultKey
	}
	if m.conn.User == "" {
		m.conn.User = os.Getenv("USER")
	}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Placeholder("Alias for this connection").
				Value(&m.conn.Name),

			huh.NewInput().
				Title("Host").
				Placeholder("Hostname or IP address").
				Value(&m.conn.Host),

			huh.NewInput().Title("User").Value(&m.conn.User),

			huh.NewInput().
				Title("Port").
				Value(&m.portStr),

			huh.NewInput().Title("Key").Placeholder(defaultKey).Value(&m.conn.IdentityFile),
		),
		huh.NewGroup(
			huh.NewInput().Title("ProxyJump").Value(&m.conn.ProxyJump),
			huh.NewInput().Title("Connect Timeout").Value(&m.conn.ConnectTimeout),
			huh.NewInput().Title("Forward Agent").Value(&m.conn.ForwardAgent),
			huh.NewInput().Title("Server Alive Interval").Value(&m.conn.ServerAliveInterval),
			huh.NewInput().Title("Server Alive Count Max").Value(&m.conn.ServerAliveCountMax),
			huh.NewInput().Title("TCP Keep Alive").Value(&m.conn.TCPKeepAlive),
		),
		huh.NewGroup(
			huh.NewInput().Title("Local Forward").Value(&m.conn.LocalForward),
			huh.NewInput().Title("Remote Forward").Value(&m.conn.RemoteForward),
			huh.NewInput().Title("Dynamic Forward").Value(&m.conn.DynamicForward),
		),
		huh.NewGroup(
			huh.NewInput().Title("Tags").Value(&m.tagsStr),
			huh.NewInput().Title("Notes").Value(&m.conn.Notes),
			huh.NewInput().Title("Args").Value(&m.conn.Args),
		),
	).WithShowHelp(false).WithShowErrors(false)
	return m
}

func (m *Model) Done() bool {
	return m.done
}

func (m *Model) Connection() connection.Connection {
	port, err := strconv.Atoi(m.portStr)
	if err != nil {
		port = 22
	}
	m.conn.Port = port
	return m.conn
}
func (m *Model) IsEditing() bool {
	return m.editing
}
func (m *Model) SetEditing(editing bool) {
	m.editing = editing
}
func (m *Model) OriginalName() string {
	return m.originalName
}
