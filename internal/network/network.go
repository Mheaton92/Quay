package network

import (
	"fmt"
	"net"
	"time"
)

type PortResult struct {
	Port   int
	Open   bool
	Banner string
}

func ScanPorts(host string, ports []int, timeout time.Duration) []PortResult {
	results := make([]PortResult, 0, len(ports))
	for _, port := range ports {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", addr, timeout)
		if err != nil {
			results = append(results, PortResult{Port: port, Open: false})
			continue
		}
		conn.Close()
		results = append(results, PortResult{Port: port, Open: true})
	}
	return results
}
