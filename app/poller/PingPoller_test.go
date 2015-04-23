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

func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
	myIcmp := Packet{}
	myNet := Net{}
	myIp, _ := myNet.ResolveIPAddr("ip", "google.fr")
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(myIp, myPacket)
	if err != nil {
		t.Error("Pinging a valid IP should not raise an error, got ", err)
	}
}

func TestPingingWithNoIpShouldTriggerError(t *testing.T) {
	myNet := Net{}
	myIp, _ := myNet.ResolveIPAddr("ip", "google.fr")
	_, err := Ping(myIp, nil)
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
	}
}

func TestPingingWithNoPacketConnShouldTriggerError(t *testing.T) {
	myIcmp := Packet{}
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(nil, myPacket)
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
	}
}
