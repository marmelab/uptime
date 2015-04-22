package poller

import (
	"net"
	"testing"
)

func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
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
