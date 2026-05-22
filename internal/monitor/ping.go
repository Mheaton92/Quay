package monitor

import (
	"time"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Pinger interface {
	Ping(host string) (time.Duration, error)
}

type ExecPinger struct{}

func NewExecPinger() *ExecPinger {
	return &ExecPinger{}
}

func (p *ExecPinger) Ping(host string) (time.Duration, error) {
	cmd := exec.Command("ping", "-c", "1", "-W", "2", host)
	out, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("host unreachable")
	}

	// Parse "time=1.23 ms" from ping output
	re := regexp.MustCompile(`time=(\d+\.?\d*)\s*ms`)
	matches := re.FindSubmatch(out)
	if len(matches) < 2 {
		return 0, fmt.Errorf("could not parse ping output")

	}

	ms, err := strconv.ParseFloat(string(matches[1]), 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(ms * float64(time.Millisecond)), nil
}

type ICMPPinger struct{}

func NewICMPPinger() *ICMPPinger {
	return &ICMPPinger{}
}

func (p *ICMPPinger) Ping(host string) (time.Duration, error) {
	// raw ICMP implementation later
	return 0, fmt.Errorf("ICMP pinger not yet implemented")
}

func NewPinger(useRawICMP bool) Pinger {
	if useRawICMP {
		return NewICMPPinger()
	}
	return NewExecPinger()
}
