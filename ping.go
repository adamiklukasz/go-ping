package ping

import (
	"time"
)

type Stats struct {
	Min time.Duration
	Max time.Duration
	Avg time.Duration
}

// returns one ping duration for a specific address
func Ping(targetAddr string) (time.Duration, error) {
	return parsePingCmd(targetAddr)
}

// returns average ping durations for a specific address
func PingN(targetAddr string, count int, interval time.Duration) (Stats, error) {
	s := Stats{
		Min: 1 * time.Hour,
		Max: 0,
		Avg: 0,
	}

	if count < 1 {
		panic("count should be a positive number")
	}

	var sum time.Duration

	for i := 0; i < count; i++ {
		d, err := Ping(targetAddr)
		if err != nil {
			return s, nil
		}

		sum += d

		if s.Min > d {
			s.Min = d
		}

		if s.Max < d {
			s.Max = d
		}

		time.Sleep(interval)
	}

	s.Avg = sum / time.Duration(count)

	return s, nil
}
