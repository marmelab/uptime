package poller

import (
	"net"
	"net/http/httptest"
	"testing"
	"net/http"
)

func TestPingValidDestination(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
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
	_, err := net.ResolveIPAddr("ip", "localhost?Ithasnosense")
	if err == nil {
		t.Error("Expected error, got", err)
	}

}

func TestPingNoDestination(t *testing.T) {
	ip, err := net.ResolveIPAddr("ip","")
	if err == nil {
		duration, error := Ping(ip)
		if error == nil {
			t.Error("Expected error, got", error)
		} else if duration > 0 {
			t.Error("Expected duration <= 0, got", duration)
		}
	}
}
