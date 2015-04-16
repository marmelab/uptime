package poller

import "testing"
import "net"

func TestPingValidDestination(t *testing.T) {
	ip, err := net.ResolveIPAddr("ip", "localhost")
	if err == nil {
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected no error, got", err)
		} else if duration <= 0 {
			t.Error("Expected duration > 0, got", duration)
		}
	}
}

func TestPingWrongDestination(t *testing.T) {
	ip, err := net.ResolveIPAddr("ip", "localhost?Ithasnosense")
	if err == nil {
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected error, got", err)
		} else if duration <= 0 {
			t.Error("Expected duration < 0, got", duration)
		}
	}
}

func TestPingNoDestination(t *testing.T) {
	ip, err := net.ResolveIPAddr("ip", "")
	if err == nil {
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected no error, got", err)
		} else if duration <= 0 {
			t.Error("Expected duration > 0, got", duration)
		}
	}
}
