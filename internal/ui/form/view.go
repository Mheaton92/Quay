package form

func (m Model) View() string {
    tabs := []string{"Basic", "Connection", "Forwarding", "Meta"}
    current := m.page

    var tabBar string
    for i, tab := range tabs {
        if i == current {
            tabBar += "[" + tab + "]"
        } else {
            tabBar += " " + tab + " "
        }
    }
    return tabBar + "\n\n" + m.form.View()
}   