package monitor

import (
	"time"
)

const MaxHistory = 60 // keep last 60 pings

type HostStats struct {
	Host		string
	Latency	time.Duration
	History	[]time.Duration
	PacketLoss	float64
	Online 			bool
	LastChecked	time.Time
}

func (s *HostStats) AddSample(d time.Duration, err error) {
	if err != nil {
		s.Online = false
		s.History = append(s.History, 0)
	} else {
		s.Online = true
		s.Latency = d
		s.History = append(s.History, d)
	}
	if len(s.History) > MaxHistory {
		s.History = s.History[1:]
	}
	s.LastChecked = time.Now()
	s.updatePacketLoss()
}

func (s *HostStats) updatePacketLoss() {
	if len(s.History) == 0 {
		return
	}
	lost := 0
	for _, d := range s.History {
		if d == 0 {
			lost++
		}
	}
	s.PacketLoss = float64(lost) / float64(len(s.History)) * 100
}
