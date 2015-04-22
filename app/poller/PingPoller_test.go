package poller

import (
	"net"
	"testing"
)

func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
<<<<<<< HEAD
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
=======
	ip, _ := net.ResolveIPAddr("ip", "localhost")
>>>>>>> tests fixed
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected no error, got", error)
		} else if duration <= 0 {
			t.Error("Expected duration > 0, got", duration)
		}
}


func TestPingingWithNoIpShouldTriggerError(t *testing.T) {
	_, err := Ping(nil)
	if (err == nil) {
		t.Error("Pinging a nil IP should raise an error got", err);
	}
}
