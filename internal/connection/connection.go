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
