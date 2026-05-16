package form

import (
	"github.com/charmbracelet/huh"
	"github.com/mheaton92/quay/internal/connection"
)

type Model struct {
	conn connection.Connection
	portStr string // temporary string for port input
	tagsStr string // temporary string for tags input
	form *huh.Form
	done bool
}


func NewForm(conn connection.Connection) Model {
    m := Model{conn: conn}
    m.form = huh.NewForm(
		// Basic feilds
        huh.NewGroup( 
            huh.NewInput().
                Title("Name").
                Value(&m.conn.Name),
            huh.NewInput().
                Title("Host").
                Value(&m.conn.Host),
            huh.NewInput().
                Title("User").
                Value(&m.conn.User),
            huh.NewInput().
                Title("Port").
                Value(&m.portStr),
			huh.NewInput().
				Title("Key").
				Value(&m.conn.IdentityFile),
        ),
		// Connection fields
		huh.NewGroup(
			huh.NewInput().
				Title("Proxyjump").
				Value(&m.conn.ProxyJump),
			huh.NewInput().
				Title("Connect Timeout").
				Value(&m.conn.ConnectTimeout),
			huh.NewInput().
				Title("Forward Agent").
				Value(&m.conn.ForwardAgent),
			huh.NewInput().
				Title("Server Alive Interval").
				Value(&m.conn.ServerAliveInterval),
			huh.NewInput().
				Title("Server Alive Count Max").
				Value(&m.conn.ServerAliveCountMax),
			huh.NewInput().
				Title("TCP Keep Alive").
				Value(&m.conn.TCPKeepAlive),
			
		),
		// Forwarding fields
		huh.NewGroup(
			huh.NewInput().
				Title("Local Forward").
				Value(&m.conn.LocalForward),
			huh.NewInput().
				Title("Remote Forward").
				Value(&m.conn.RemoteForward),
			huh.NewInput().
				Title("Dynamic Forward").
				Value(&m.conn.DynamicForward),
		),
		// Meta fields
		huh.NewGroup(
			huh.NewInput().
				Title("Tags").
				Value(&m.tagsStr),
			huh.NewInput().
				Title("Notes").
				Value(&m.conn.Notes),
			huh.NewInput().
				Title("Args").
				Value(&m.conn.Args),
		),
    )
    return m
}