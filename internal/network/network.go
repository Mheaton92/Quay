package network

import (
	"crypto/tls"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type PortResult struct {
	Port   int
	Open   bool
	Banner string
}

type BandwidthResult struct {
    Download float64 // MB/s
    Upload   float64 // MB/s
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

func WakeOnLAN(macAddr string) error {
	mac, err := net.ParseMAC(macAddr)
	if err != nil {
		return fmt.Errorf("invalid MAC address: %s", macAddr)
	}

	// Build magic packet
	packet := make([]byte, 102)
	// First 6 bytes are 0xFF
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}
	// Repeat MAC address 16 times
	for i := 1; i <= 16; i++ {
		copy(packet[i*6:], mac)
	}

	// Send UDP broadcast
	conn, err := net.Dial("udp", "255.255.255.255:9")
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Write(packet)
	return err
}

type DNSResult struct {
	Host   string
	IPs    []string
	CNAMEs []string
	MXs    []string
}

func DNSLookup(host string) DNSResult {
    result := DNSResult{Host: host}

    // Check if input is an IP — do reverse lookup
    if net.ParseIP(host) != nil {
        names, err := net.LookupAddr(host)
        if err == nil {
            result.CNAMEs = names
        }
        result.IPs = []string{host}
        return result
    }

    // Regular hostname lookup
    ips, err := net.LookupHost(host)
    if err == nil {
        result.IPs = ips
    }

    cname, err := net.LookupCNAME(host)
    if err == nil && cname != host+"." {
        result.CNAMEs = []string{cname}
    }

    mxs, err := net.LookupMX(host)
    if err == nil {
        for _, mx := range mxs {
            result.MXs = append(result.MXs, fmt.Sprintf("%s (priority %d)", mx.Host, mx.Pref))
        }
    }

    return result
}

type HopResult struct {
	Hop     int
	Host    string
	Latency string
}

func Traceroute(host string) ([]HopResult, error) {
    cmd := exec.Command("traceroute", "-m", "30", "-w", "2", host)
    out, err := cmd.Output()
    if err != nil {
        // try tracepath as fallback
        cmd = exec.Command("tracepath", host)
        out, err = cmd.Output()
        if err != nil {
            return nil, fmt.Errorf("traceroute not available")
        }
        return parseTracepath(string(out)), nil
    }
    return parseTraceroute(string(out)), nil
}

func parseTraceroute(output string) []HopResult {
    var hops []HopResult
    lines := strings.Split(output, "\n")
    for _, line := range lines[1:] {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }
        hops = append(hops, HopResult{
            Hop:  len(hops) + 1,
            Host: line,
        })
    }
    return hops
}

func parseTracepath(output string) []HopResult {
    var hops []HopResult
    lines := strings.Split(output, "\n")
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" || strings.HasPrefix(line, "Resume:") {
            continue
        }
        hops = append(hops, HopResult{
            Hop:  len(hops) + 1,
            Host: line,
        })
    }
    return hops
}

type SSLResult struct {
	Host     string
	Issuer   string
	Expires  time.Time
	DaysLeft int
	Valid    bool
}

func SSLCheck(host string, port int) SSLResult {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: false,
	})
	if err != nil {
		return SSLResult{Host: host, Valid: false}
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)

	return SSLResult{
		Host:     host,
		Issuer:   cert.Issuer.CommonName,
		Expires:  cert.NotAfter,
		DaysLeft: daysLeft,
		Valid:    true,
	}
}

type HostResult struct {
	IP     string
	Online bool
}

func ScanSubnet(subnet string) []HostResult {
	var results []HostResult
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Generate all IPs in subnet e.g. 192.168.4.1-254
	parts := strings.Split(subnet, ".")
	if len(parts) != 4 {
		return nil
	}
	base := parts[0] + "." + parts[1] + "." + parts[2] + "."

	for i := 1; i <= 254; i++ {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
			err := cmd.Run()
			mu.Lock()
			results = append(results, HostResult{
				IP:     ip,
				Online: err == nil,
			})
			mu.Unlock()
		}(fmt.Sprintf("%s%d", base, i))
	}
	wg.Wait()
	return results
}

func BandwidthTest(host string, user string, port int, keyFile string) (BandwidthResult, error) {
    // Download test — pull from remote
    downloadCmd := exec.Command("ssh",
        "-p", fmt.Sprintf("%d", port),
        "-i", keyFile,
        "-o", "StrictHostKeyChecking=no",
        fmt.Sprintf("%s@%s", user, host),
        "dd if=/dev/zero bs=1M count=10 2>/dev/null")

    start := time.Now()
    out, err := downloadCmd.Output()
    elapsed := time.Since(start).Seconds()
    if err != nil {
        return BandwidthResult{}, fmt.Errorf("download test failed: %s", err)
    }
    downloadMBps := float64(len(out)) / elapsed / 1024 / 1024

    // Upload test — push to remote
    uploadCmd := exec.Command("ssh",
        "-p", fmt.Sprintf("%d", port),
        "-i", keyFile,
        "-o", "StrictHostKeyChecking=no",
        fmt.Sprintf("%s@%s", user, host),
        "dd of=/dev/null bs=1M 2>/dev/null")

    data := make([]byte, 10*1024*1024) // 10MB
    uploadCmd.Stdin = strings.NewReader(string(data))

    start = time.Now()
    err = uploadCmd.Run()
    elapsed = time.Since(start).Seconds()
    if err != nil {
        return BandwidthResult{Download: downloadMBps}, nil
    }
    uploadMBps := 10.0 / elapsed

    return BandwidthResult{
        Download: downloadMBps,
        Upload:   uploadMBps,
    }, nil
}
