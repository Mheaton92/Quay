package connection

import (
	"encoding/json"
	"errors"
	"github.com/mheaton92/quay/internal/config"
	"os"
)

type Connection struct {
	// Basic
	Name         string
	Host				 string
	User         string
	Port         int
	IdentityFile string

	// Connection
	ProxyJump           string
	ConnectTimeout      string
	ForwardAgent        string
	ServerAliveInterval string

	// Forwarding
	LocalForward   string
	RemoteForward  string
	DynamicForward string

	// Meta
	Online bool
	Tags   []string
	Notes  string
	Args   string
}

type Store struct {
	Connections []Connection
}

func NewConnection(name, host, user string, port int) Connection {
	return Connection{
		Name: name,
		Host: host,
		User: user,
		Port: port,
	}
}

func (s *Store) Save() error {
	dir, err := config.ConfigDir()
	if err != nil {
		return err
	}
	path := dir + "/connections.json"
	data, err := json.Marshal(s.Connections)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (s *Store) Load() error {
	dir, err := config.ConfigDir()
	if err != nil {
		return err
	}
	path := dir + "/connections.json"
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			s.Connections = []Connection{}
			return nil
		}
		return err
	}
	err = json.Unmarshal(content, &s.Connections)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Add(c Connection) error {
	for _, conn := range s.Connections {
		if conn.Name == c.Name {
			return errors.New("Connection already exists")
		}
	}
	s.Connections = append(s.Connections, c)
	return s.Save()
}

func (s *Store) Delete(name string) error {
	for i, conn := range s.Connections {
		if conn.Name == name {
			s.Connections = append(s.Connections[:i], s.Connections[i+1:]...)
			return s.Save()
		}
	}
	return errors.New("connection not found")
}
func NewStore() (*Store, error) {
	store := &Store{}
	err := store.Load()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) Edit(name string, updated Connection) error {
	for i, conn := range s.Connections {
		if conn.Name == name {
			s.Connections[i] = updated
			return s.Save()
		}
	}
	return errors.New("connection not found")
}
