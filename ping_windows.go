package ping

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"time"
)

func parsePingCmd(targetAddr string) (time.Duration, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cmd := exec.CommandContext(ctx, "ping", "-n", "1", targetAddr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}

	i := bytes.LastIndex(out, []byte("= "))
	d, err := time.ParseDuration(strings.TrimSpace(string(out[i+1:])))
	if err != nil {
		return 0, err
	}

	return d, nil
}
