package help

// Field help descriptions adapted from lazyssh by Adembc
// Source: https://github.com/Adembc/lazyssh
// Licensed under Apache 2.0 License
// Original work Copyright 2025 Adembc

// FieldHelp contains help information for SSH config fields
type FieldHelp struct {
	Field       string   // Field name
	Description string   // Brief description
	Syntax      string   // Syntax format
	Examples    []string // Usage examples
	Default     string   // Default value
	Since       string   // OpenSSH version when introduced
	Category    string   // Category for grouping
}

// HelpDisplayMode defines how help is displayed
type HelpDisplayMode int

const (
	HelpModeOff     HelpDisplayMode = iota // No help shown
	HelpModeCompact                        // Single line help
	HelpModeNormal                         // Standard help panel
	HelpModeFull                           // Detailed help with all info
)

// GetFieldHelp returns help information for a specific field
func GetFieldHelp(fieldName string) *FieldHelp {
	if help, exists := fieldHelpData[fieldName]; exists {
		return &help
	}
	return nil
}

// fieldHelpData contains help information for all SSH config fields
var fieldHelpData = map[string]FieldHelp{
	// Basic fields
	"Name": {
		Field:       "Name",
		Description: "A nickname or abbreviation for the host. This is what you type after 'ssh' command.",
		Syntax:      "any_string_without_spaces",
		Examples:    []string{"myserver", "prod-db", "dev-web-01"},
		Default:     "(required)",
		Category:    "Basic",
	},
	"Host": {
		Field:       "Host",
		Description: "The real hostname or IP address to connect to. Can be a domain name or IP address.",
		Syntax:      "hostname | ip_address",
		Examples:    []string{"example.com", "192.168.1.100", "2001:db8::1"},
		Default:     "(required)",
		Category:    "Basic",
	},
	"User": {
		Field:       "User",
		Description: "Username for logging into the remote machine. If not specified, uses current username.",
		Syntax:      "username",
		Examples:    []string{"root", "ubuntu", "admin", "deploy"},
		Default:     "current username",
		Category:    "Basic",
	},
	"Port": {
		Field:       "Port",
		Description: "The port number to connect to on the remote host. Standard SSH port is 22.",
		Syntax:      "port_number (1-65535)",
		Examples:    []string{"22", "2222", "8022"},
		Default:     "22",
		Category:    "Basic",
	},
	"IdentityFile": {
		Field:       "IdentityFile",
		Description: "Path to SSH private key files for authentication. Multiple keys can be specified.",
		Syntax:      "path[,path,...]",
		Examples:    []string{"~/.ssh/id_ed25519", "~/.ssh/id_rsa,~/.ssh/id_ed25519"},
		Default:     "~/.ssh/id_rsa, ~/.ssh/id_ed25519, etc.",
		Category:    "Basic",
	},
  "MACAddress": {
		Field:       "MAC Address",
		Description: "Hardware MAC address of the machine for Wake on LAN.",
		Syntax:      "xx:xx:xx:xx:xx:xx",
		Examples:    []string{"aa:bb:cc:dd:ee:ff"},
		Default:     "none",
		Category:    "Basic",
	}, 
	// Connection fields
	"ProxyJump": {
		Field:       "ProxyJump",
		Description: "Specifies one or more jump hosts (bastion hosts) to reach the destination. Useful for accessing servers behind firewalls.",
		Syntax:      "[user@]host[:port][,[user@]host[:port]]",
		Examples:    []string{"bastion.example.com", "jump1.com,jump2.com", "user@proxy:2222"},
		Default:     "none",
		Since:       "OpenSSH 7.3+",
		Category:    "Connection",
	},
	"ConnectTimeout": {
		Field:       "ConnectTimeout",
		Description: "Timeout in seconds for establishing the connection. Useful for slow or unreliable networks.",
		Syntax:      "seconds | none",
		Examples:    []string{"10", "30", "none"},
		Default:     "none (system default)",
		Category:    "Connection",
	},
	"ForwardAgent": {
		Field:       "ForwardAgent",
		Description: "Forward SSH agent connection to remote host. Allows using local SSH keys on remote servers.",
		Syntax:      "yes | no",
		Examples:    []string{"yes", "no"},
		Default:     "no",
		Category:    "Connection",
	},
	"ServerAliveInterval": {
		Field:       "ServerAliveInterval",
		Description: "Seconds between keepalive messages. Prevents connection drops on idle connections.",
		Syntax:      "seconds",
		Examples:    []string{"60", "120", "300"},
		Default:     "0 (disabled)",
		Category:    "Connection",
	},
	"ServerAliveCountMax": {
		Field:       "ServerAliveCountMax",
		Description: "Number of keepalive messages before disconnecting.",
		Syntax:      "count",
		Examples:    []string{"3", "5", "10"},
		Default:     "3",
		Category:    "Connection",
	},
	"TCPKeepAlive": {
		Field:       "TCPKeepAlive",
		Description: "Send TCP keepalive messages to detect broken connections.",
		Syntax:      "yes | no",
		Examples:    []string{"yes", "no"},
		Default:     "yes",
		Category:    "Connection",
	},
	// Port forwarding fields
	"LocalForward": {
		Field:       "LocalForward",
		Description: "Forward a local port to a remote address. Useful for accessing remote services through SSH tunnel.",
		Syntax:      "[bind_address:]port:host:hostport (CLI format, auto-converted for config file)",
		Examples:    []string{"8080:localhost:80", "3306:db.internal:3306", "*:8080:localhost:80"},
		Default:     "none",
		Category:    "Forwarding",
	},
	"RemoteForward": {
		Field:       "RemoteForward",
		Description: "Forward a remote port to a local address. Allows remote users to access local services.",
		Syntax:      "[bind_address:]port:host:hostport (CLI format, auto-converted for config file)",
		Examples:    []string{"8080:localhost:3000", "*:80:localhost:8080"},
		Default:     "none",
		Category:    "Forwarding",
	},
	"DynamicForward": {
		Field:       "DynamicForward",
		Description: "Create a SOCKS proxy on the specified port. Useful for routing traffic through SSH.",
		Syntax:      "[bind_address:]port",
		Examples:    []string{"1080", "localhost:1080", "*:1080"},
		Default:     "none",
		Category:    "Forwarding",
	},
	// Meta
	"Tags": {
		Field:       "Tags",
		Description: "Custom tags for organizing and filtering servers. Comma-separated list.",
		Syntax:      "tag1[,tag2,...]  ",
		Examples:    []string{"production", "development,staging", "web,frontend"},
		Default:     "none",
		Category:    "Basic",
	},
	"Notes": {
		Field:       "Notes",
		Description: "Free text notes about this connection. Reminders, credentials location, maintenance notes etc.",
		Syntax:      "any text",
		Examples:    []string{"Reboot scheduled Fridays", "Uses jump host via bastion", "API token stored in vault"},
		Default:     "none",
		Category:    "Meta",
	},
	"Args": {
		Field:       "Args",
		Description: "Additional SSH command line arguments passed directly to the SSH command.",
		Syntax:      "ssh_flag [ssh_flag ...]",
		Examples:    []string{"-p 8022", "-v", "-C -X"},
		Default:     "none",
		Category:    "Basic",
	},
}

// GetFieldsByCategory returns all fields in a specific category
func GetFieldsByCategory(category string) []string {
	// Pre-count to allocate correct capacity
	count := 0
	for _, help := range fieldHelpData {
		if help.Category == category {
			count++
		}
	}

	fields := make([]string, 0, count)
	for name, help := range fieldHelpData {
		if help.Category == category {
			fields = append(fields, name)
		}
	}
	return fields
}

// GetAllCategories returns all available help categories
func GetAllCategories() []string {
	categories := make(map[string]bool)
	for _, help := range fieldHelpData {
		categories[help.Category] = true
	}

	// Convert to slice with defined order
	orderedCategories := []string{
		"Basic", "Connection", "Forwarding", "Meta",
	}

	var result []string
	for _, cat := range orderedCategories {
		if categories[cat] {
			result = append(result, cat)
		}
	}
	return result
}
