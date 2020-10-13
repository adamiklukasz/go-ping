package ping

import (
	"errors"
	"testing"
)

var (
	exampleWin = []byte(`
Pinging solarwinds.com [74.115.13.112] with 32 bytes of data:
Reply from 74.115.13.112: bytes=32 time=164ms TTL=236

Ping statistics for 74.115.13.112:
    Packets: Sent = 1, Received = 1, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 164ms, Maximum = 164ms, Average = 164ms

`)

	exampleLinux = []byte(`
PING solarwinds.com (74.115.13.112) 56(84) bytes of data.
64 bytes from 74.115.13.112 (74.115.13.112): icmp_seq=1 ttl=235 time=172 ms

--- solarwinds.com ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 172.409/172.409/172.409/0.000 ms
`)

	exampleWinFail   = []byte(`Ping request could not find host solarwinds.com. Please check the name and try again.`)
	exampleLinuxFail = []byte(`ping: solarwinds.com: Temporary failure in name resolution`)
)

func TestPingWinOk(t *testing.T) {
	executeCmd = func(cmdArgs ...string) ([]byte, error) {
		return exampleWin, nil
	}

	d, err := pingWin("solarwinds.com")

	if d.String() != "164ms" {
		t.Errorf("invalid value")
	}

	if err != nil {
		t.Errorf("error raised")
	}
}

func TestPingLinuxOk(t *testing.T) {
	executeCmd = func(cmdArgs ...string) ([]byte, error) {
		return exampleLinux, nil
	}

	d, err := pingLinux("solarwinds.com")

	if d.String() != "172ms" {
		t.Errorf("invalid value")
	}

	if err != nil {
		t.Errorf("error raised")
	}
}

func TestPingWinFail(t *testing.T) {
	executeCmd = func(cmdArgs ...string) ([]byte, error) {
		return exampleWinFail, errors.New("")
	}

	_, err := pingWin("solarwinds.com")

	if err == nil {
		t.Errorf("error expected")
	}
}

func TestPingLinuxFail(t *testing.T) {
	executeCmd = func(cmdArgs ...string) ([]byte, error) {
		return exampleLinuxFail, errors.New("")
	}

	_, err := pingLinux("solarwinds.com")

	if err == nil {
		t.Errorf("error expected")
	}
}
