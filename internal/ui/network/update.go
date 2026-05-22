package network

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	internalnetwork "github.com/mheaton92/quay/internal/network"
	"strconv"
	"strings"
	"time"
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

	case ToolWakeOnLAN:
		if m.conn.MACAddress != "" {
			// MAC already set, run directly
			m.runSelectedTool()
		} else {
			m.inputLabel = "MAC Address:"
			m.input.Placeholder = "aa:bb:cc:dd:ee:ff"
			m.input.SetValue("")
			m.input.Focus()
			m.showInput = true
		}

	case ToolDNSLookup:
		m.inputLabel = "Hostname to lookup:"
		m.input.Placeholder = m.conn.Host
		m.input.SetValue(m.conn.Host)
		m.input.Focus()
		m.showInput = true

	case ToolTraceroute:
		m.inputLabel = "Host to trace:"
		m.input.Placeholder = m.conn.Host
		m.input.SetValue(m.conn.Host)
		m.input.Focus()
		m.showInput = true

	case ToolSSLChecker:
		m.inputLabel = "Host to check (port 443 default)"
		m.input.Placeholder = m.conn.Host
		m.input.SetValue(m.conn.Host)
		m.input.Focus()
		m.showInput = true

	case ToolSubnetScanner:
		m.inputLabel = "Subnet to scan (e.g. 192.168.4.0):"
		m.input.Placeholder = "192.168.4.0"
		m.input.SetValue("")
		m.input.Focus()
		m.showInput = true

	case ToolBandwidthTest:
		m.inputLabel = "Press enter to start bandwidth test:"
		m.input.Placeholder = ""
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

	case ToolWakeOnLAN:
		mac := m.conn.MACAddress
		if mac == "" {
			mac = m.input.Value()
		}
		err := internalnetwork.WakeOnLAN(mac)
		if err != nil {
			m.err = err
		} else {
			m.result = "✓ Magic packet sent to " + mac
		}

	case ToolDNSLookup:
		host := m.input.Value()
		if host == "" {
			host = m.conn.Host
		}
		result := internalnetwork.DNSLookup(host)
		m.result = formatDNSResult(result)

	case ToolTraceroute:
		host := m.input.Value()
		if host == "" {
			host = m.conn.Host
		}
		hops, err := internalnetwork.Traceroute(host)
		if err != nil {
			m.err = err
		} else {
			m.result = formatTraceroute(hops)
		}

	case ToolSSLChecker:
		host := m.input.Value()
		if host == "" {
			host = m.conn.Host
		}
		result := internalnetwork.SSLCheck(host, 443)
		m.result = formatSSLResult(result)

	case ToolSubnetScanner:
		subnet := m.input.Value()
		results := internalnetwork.ScanSubnet(subnet)
		m.result = formatSubnetResults(results)

	case ToolBandwidthTest:
    result, err := internalnetwork.BandwidthTest(
        m.conn.Host,
        m.conn.User,
        m.conn.Port,
        m.conn.IdentityFile,
    )
    if err != nil {
        m.err = err
    } else {
        m.result = formatBandwidthResult(result)
    }	
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

func formatDNSResult(r internalnetwork.DNSResult) string {
	var output string
	output += fmt.Sprintf("  Host: %s\n\n", r.Host)
	if len(r.IPs) > 0 {
		output += "  A Records:\n"
		for _, ip := range r.IPs {
			output += fmt.Sprintf("    ● %s\n", ip)
		}
	}
	if len(r.CNAMEs) > 0 {
		output += "\n  CNAMEs:\n"
		for _, c := range r.CNAMEs {
			output += fmt.Sprintf("    ● %s\n", c)
		}
	}
	if len(r.MXs) > 0 {
		output += "\n  MX Records:\n"
		for _, mx := range r.MXs {
			output += fmt.Sprintf("    ● %s\n", mx)
		}
	}
	return output
}

func formatTraceroute(hops []internalnetwork.HopResult) string {
	var output string
	for _, hop := range hops {
		output += fmt.Sprintf("  %2d  %s\n", hop.Hop, hop.Host)
	}
	return output
}

func formatSSLResult(r internalnetwork.SSLResult) string {
	if !r.Valid {
		return "  ✗ Could not connect or no SSL certificate found"
	}
	color := "✓"
	if r.DaysLeft < 30 {
		color = "⚠"
	}
	return fmt.Sprintf("  %s Host:     %s\n  Issuer:   %s\n  Expires:  %s\n  Days left: %d",
		color, r.Host, r.Issuer,
		r.Expires.Format("2006-01-02"),
		r.DaysLeft)
}

func formatSubnetResults(results []internalnetwork.HostResult) string {
	var output string
	for _, r := range results {
		if r.Online {
			output += fmt.Sprintf("  ● %s\n", r.IP)
		}
	}
	if output == "" {
		return "  No hosts found"
	}
	return output
}

func formatBandwidthResult(r internalnetwork.BandwidthResult) string {
    return fmt.Sprintf("  ↓ Download: %.2f MB/s\n  ↑ Upload:   %.2f MB/s",
        r.Download,
        r.Upload)
}
