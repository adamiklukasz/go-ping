package ping

import (
	"time"
)

// returns one ping duration for a specific address
func Ping(targetAddr string) (time.Duration, error) {
	return parsePingCmd(targetAddr)
}

// returns average ping durations for a specific address
func PingN(targetAddr string, count int, interval time.Duration) (time.Duration, error) {
	if count < 1 {
		panic("count should be a positive number")
	}

	var sum time.Duration

	for i := 0; i < count; i++ {
		d, err := Ping(targetAddr)
		if err != nil {
			return 0, nil
		}

		sum += d

		time.Sleep(interval)
	}

	return sum / time.Duration(count), nil
}
