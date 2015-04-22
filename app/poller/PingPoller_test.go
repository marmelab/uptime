package poller

import (
	"golang.org/x/net/icmp"
	"net"
	"testing"
)

type Packet struct {
}

type IcmpMock interface {
	ListenPacket(network string, address string) *icmp.PacketConn
}

func (pack Packet) ListenPacket(n string, a string) *icmp.PacketConn {
	var p icmp.PacketConn
	ptr := &p
	return ptr
}

func TestPingValidDestination(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected no error, got", error)
		} else if duration <= 0 {
			t.Error("Expected duration > 0, got", duration)
		}
}


type Net struct {
}

type NetMock interface {
	ResolveIPAddr(proto string, address string) (*net.IPAddr, error)
}

func (ipad Net) ResolveIPAddr(p string, a string) (*net.IPAddr, error) {
	var i net.IPAddr
	i.Zone = "france"
	i.IP = net.ParseIP("localhost")
	ptr := &i
	return ptr, nil
}

func TestPingWithValidIpShouldNotTriggerError(t *testing.T) {
	myIcmp := Packet{}
	myNet := Net{}
	myIp, _ := myNet.ResolveIPAddr("ip", "google.fr")
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(myIp, myPacket)
	if err != nil {
		t.Error("Pinging a valid IP should not raise an error, got ", err)
	}
}


func TestPingingWithNoIPConnShouldTriggerError(t *testing.T) {
	myIcmp := Packet{}
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(nil, myPacket)
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)est failed again
	}
}
