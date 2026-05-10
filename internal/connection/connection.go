package connection

import (
	"encoding/json"
	"github.com/mheaton92/quay/internal/config"
	"os"
)

type Connection struct {
	Name   string
	IP     string
	User   string
	Port   int
	Online bool
	Tags   []string
	Notes  string
	Args   string
}

func NewConnection(name, ip, user string, port int) Connection {
	return Connection{
		Name: name,
		IP:   ip,
		User: user,
		Port: port,
	}
}

func Save(connections []Connection) error {
	dir, err := config.ConfigDir()
	if err != nil {
		return err
	}
	path := dir + "/connections.json"
	data, err := json.Marshal(connections)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func Load() ([]Connection, error) {
	dir, err := config.ConfigDir()
	if err != nil {
		return nil, err
	}
	path := dir + "/connections.json"
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var connections []Connection
	err = json.Unmarshal(content, &connections)
	if err != nil {
		return nil, err
	}
	return connections, nil
}
