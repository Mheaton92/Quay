package monitor

import "time"

var blocks = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}

func Sparkline(history []time.Duration, width int) string {
	if len(history) == 0 {
		return ""
	}

	// Find max for scaling
	var max time.Duration
	for _, d := range history {
		if d > max {
			max = d
		}
	}
	if max == 0 {
		max = 1
	}

	// Take last width samples
	samples := history
	if len(samples) > width {
		samples = samples[len(samples)-width:]
	}

	result := ""
	for _, d := range samples {
		if d == 0 {
			result += " "
			continue
		}
		idx := int(float64(d) / float64(max) * float64(len(blocks)-1))
		if idx >= len(blocks) {
			idx = len(blocks) - 1
		}
		result += blocks[idx]
	}
	return result
}
