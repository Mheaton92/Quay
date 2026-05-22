package network

import (
    "fmt"
    "strconv"
    "strings"
    "time"
    tea "github.com/charmbracelet/bubbletea"
    internalnetwork "github.com/mheaton92/quay/internal/network"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
    if m.showInput {
        var cmd tea.Cmd
        m.input, cmd = m.input.Update(msg)
        if key, ok := msg.(tea.KeyMsg); ok {
            switch key.String() {
            case "enter":
                m.showInput = false
                m.runSelectedTool()
            case "esc":
                m.showInput = false
                m.input.SetValue("")
            }
        }
        return m, cmd
    }

    if key, ok := msg.(tea.KeyMsg); ok {
        switch key.String() {
        case "j", "down":
            if m.cursor < len(m.tools)-1 {
                m.cursor++
            }
        case "k", "up":
            if m.cursor > 0 {
                m.cursor--
            }
        case "?":
            m.showHelp = !m.showHelp
        case "enter":
            m.promptForInput()
        }
    }
    return m, nil
}

func (m *Model) promptForInput() {
    switch Tool(m.cursor) {
    case ToolPortScanner:
        m.inputLabel = "Ports to scan (comma separated):"
        m.input.Placeholder = "22,80,443,8006"
        m.input.SetValue("")
        m.input.Focus()
        m.showInput = true
    }
}

func (m *Model) runSelectedTool() {
    switch Tool(m.cursor) {
    case ToolPortScanner:
        ports := parsePorts(m.input.Value())
        results := internalnetwork.ScanPorts(m.conn.Host, ports, 2*time.Second)
        m.result = formatPortResults(results)
    }
}

func parsePorts(input string) []int {
    var ports []int
    for _, p := range strings.Split(input, ",") {
        p = strings.TrimSpace(p)
        if port, err := strconv.Atoi(p); err == nil {
            ports = append(ports, port)
        }
    }
    return ports
}

func formatPortResults(results []internalnetwork.PortResult) string {
    var output string
    for _, r := range results {
        if r.Open {
            output += fmt.Sprintf("  ● port %-6d open\n", r.Port)
        } else {
            output += fmt.Sprintf("  ○ port %-6d closed\n", r.Port)
        }
    }
    return output
}
