package ui

func (m Model) View() string {
	var output string
	for _, conn := range m.store.Connections {
		output += conn.Name + " " + conn.IP + "\n"
	}
	return output
}