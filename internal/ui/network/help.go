package network

type toolHelp struct {
    name        string
    description string
    usage       string
    whyUseIt    string
}

var toolHelpData = map[Tool]toolHelp{
    ToolPortScanner: {
        name:        "Port Scanner",
        description: "Checks if specific TCP ports are open on a target host.",
        usage:       "Enter comma separated ports e.g. 22,80,443,8006",
        whyUseIt:    "Verify services are running, troubleshoot connectivity, check firewall rules",
    },
    ToolWakeOnLAN: {
        name:        "Wake on LAN",
        description: "Sends a magic packet to wake a sleeping machine.",
        usage:       "Requires MAC address of target machine. Must be on same network.",
        whyUseIt:    "Wake lab machines remotely without physical access",
    },
    ToolDNSLookup: {
        name:        "DNS Lookup",
        description: "Resolves hostnames and shows A, CNAME, and MX records.",
        usage:       "Enter any hostname to resolve",
        whyUseIt:    "Check DNS resolution, verify AdGuard is working, debug connectivity",
    },
    ToolTraceroute: {
        name:        "Traceroute",
        description: "Shows the network path to a host hop by hop.",
        usage:       "Enter target hostname or IP",
        whyUseIt:    "Find where latency is introduced, debug routing issues",
    },
    ToolSSLChecker: {
        name:        "SSL Checker",
        description: "Checks SSL certificate validity and expiry date.",
        usage:       "Enter hostname — checks port 443 by default",
        whyUseIt:    "Monitor cert expiry on your services before they break",
    },
    ToolSubnetScanner: {
        name:        "Subnet Scanner",
        description: "Scans a subnet for active hosts using ping.",
        usage:       "Enter subnet e.g. 192.168.4.0",
        whyUseIt:    "Discover devices on your network, find new machines",
    },
    ToolBandwidthTest: {
        name:        "Bandwidth Test",
        description: "Measures network throughput between local and remote host over SSH.",
        usage:       "Press enter to start — uses dd over SSH to measure speed",
        whyUseIt:    "Benchmark network performance between lab machines",
    },
}
