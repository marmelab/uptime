package poller

import (
	"golang.org/x/net/icmp"
	"net"
	"testing"
)

type IcmpMock interface {
	ListenPacket(network string, address string) *icmp.PacketConn
}

func (pack Packet) ListenPacket(n string, a string) *icmp.PacketConn {
	var p icmp.PacketConn
	ptr := &p
	return ptr
}

type NetMock interface {
	ResolveIPAddr(proto string, address string) (*net.IPAddr, error)
}
type Packet struct {
}
type Net struct {
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
	myIp, _ := myNet.ResolveIPAddr("ip", "")
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(myIp, myPacket)
	if err != nil {
		t.Error("Pinging a valid IP should not raise an error, got ", err)
	}
}

func TestPingWithNoIpShouldTriggerError(t *testing.T) {
	myIcmp := Packet{}
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(nil, myPacket)
	if err == nil {
		t.Error("Pinging a nil IP should raise an error, got ", err)
	}
}
func TestPingingWithNoIPConnShouldNotTriggerError(t *testing.T) {
	myNet := Net{}
	myIp, _ := myNet.ResolveIPAddr("ip", "")
	_, err := Ping(myIp, nil)
	if err != nil {
		t.Error("Pinging a nil PacketConn should not raise an error got", err)
	}
}

func TestHttpPingWithValidUrlShouldNotTriggerError(t *testing.T) {
	_, err := HttpPing("google.fr","http")
	if(err != nil) {
		t.Error("Pinging a valid url with http method should not raise an error, got ", err)
	}	
}

func TestHttpsPingWithValidUrlShouldNotTriggerError(t *testing.T) {
	_, err := HttpPing("google.fr","https")
	if(err != nil) {
		t.Error("Pinging a valid url with http method should not raise an error, got ", err)
	}	
}
