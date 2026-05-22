package monitor

import (
	"sync"
	"time"
)

const PingInterval = 5 * time.Second

type Monitor struct {
	hosts  map[string]*HostStats
	pinger Pinger
	mu     sync.RWMutex
	done   chan struct{}
}

func NewMonitor(useRawICMP bool) *Monitor {
	return &Monitor{
		hosts:  make(map[string]*HostStats),
		pinger: NewPinger(useRawICMP),
		done:   make(chan struct{}),
	}
}

func (m *Monitor) Add(host string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.hosts[host]; !exists {
		m.hosts[host] = &HostStats{Host: host}
	}
}

func (m *Monitor) Remove(host string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.hosts, host)
}

func (m *Monitor) Stats(host string) *HostStats {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hosts[host]
}

func (m *Monitor) Start() {
	go func() {
		m.pingAll() // ping immediately on start
		ticker := time.NewTicker(PingInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				m.pingAll()
			case <-m.done:
				return
			}
		}
	}()
}

func (m *Monitor) Stop() {
	close(m.done)
}

func (m *Monitor) pingAll() {
	m.mu.RLock()
	hosts := make([]string, 0, len(m.hosts))
	for host := range m.hosts {
		hosts = append(hosts, host)
	}
	m.mu.RUnlock()

	var wg sync.WaitGroup
	for _, host := range hosts {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			d, err := m.pinger.Ping(h)
			m.mu.Lock()
			if stats, ok := m.hosts[h]; ok {
				stats.AddSample(d, err)
			}
			m.mu.Unlock()
		}(host)
	}
	wg.Wait()
}
