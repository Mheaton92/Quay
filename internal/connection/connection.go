package connection

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
