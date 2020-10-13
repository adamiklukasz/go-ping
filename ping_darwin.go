package ping

import (
	"time"
)

func parsePingCmd(targetAddr string) (time.Duration, error) {
	return pingLinux(targetAddr)
}
