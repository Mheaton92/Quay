package ui

func (m Model) View() string {
	var output string
	for i, conn := range m.store.Connections {
	if i == m.cursor {
		output += "▶ "
	} else {
		output += "  "
	}
	output += conn.Name + " " + conn.IP + "\n"
}
return output
}