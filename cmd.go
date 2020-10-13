package ping

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"
)

const (
	cmdExecutionTimeout = 10 * time.Second
)

var (
	errorParse = errors.New("unexpected ping output")
)

var executeCmd = func(cmdArgs ...string) ([]byte, error) {
	ctx, _ := context.WithTimeout(context.Background(), cmdExecutionTimeout)
	cmd := exec.CommandContext(ctx, cmdArgs[0], cmdArgs[1:]...)

	return cmd.CombinedOutput()
}

func pingWin(targetAddr string) (time.Duration, error) {
	out, err := executeCmd("ping", "-n", "1", targetAddr)
	if err != nil {
		return 0, err
	}

	i := bytes.LastIndex(out, []byte("= "))
	if i == -1 {
		return 0, errorParse
	}

	d, err := time.ParseDuration(strings.TrimSpace(string(out[i+1:])))
	if err != nil {
		return 0, err
	}

	return d, nil
}

func pingLinux(targetAddr string) (time.Duration, error) {
	out, err := executeCmd("ping", "-c", "1", targetAddr)
	if err != nil {
		return 0, err
	}

	i := bytes.LastIndex(out, []byte("time="))
	if i == -1 {
		return 0, errorParse
	}

	j := bytes.Index(out[i:], []byte("\n"))
	if j == -1 {
		return 0, errorParse
	}

	durStr := strings.ReplaceAll(string(out[i+5:i+j]), " ", "")

	dur, err := time.ParseDuration(durStr)
	if err != nil {
		return 0, err
	}

	return dur, nil
}
