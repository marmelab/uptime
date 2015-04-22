package poller

import (
	"golang.org/x/net/icmp"
	"net"
<<<<<<< HEAD
=======
	"net/http"
	"net/http/httptest"
>>>>>>> add PingPoller_test.go
	"testing"
)

<<<<<<< HEAD
<<<<<<< HEAD
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
=======
<<<<<<< HEAD
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
<<<<<<< HEAD
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
>>>>>>>  code and test fixed
=======
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
=======
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
>>>>>>> test failed again
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected no error, got", error)
		} else if duration <= 0 {
			t.Error("Expected duration > 0, got", duration)
		}
<<<<<<< HEAD
=======
type Packet struct {
>>>>>>>  code and test fixed
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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
func TestPingWrongDestination(t *testing.T) {
	_, err := net.ResolveIPAddr("ip", "localhost?Ithasnosense")
	if err == nil {
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected error, got", error)
		} else if duration <= 0 {
			t.Error("Expected duration < 0, got", duration)
		}
>>>>>>> add PingPoller_test.go
	}
}


func TestPingingWithNoIPConnShouldTriggerError(t *testing.T) {
=======
func TestPingingWithNoIpShouldTriggerError(t *testing.T) {
	myNet := Net{}
	myIp, _ := myNet.ResolveIPAddr("ip", "google.fr")
	_, err := Ping(myIp, nil)
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
	}
}

<<<<<<< HEAD
func TestPingingWithNoPacketConnShouldTriggerError(t *testing.T) {
>>>>>>>  code and test fixed
=======

func TestPingingWithNoIPConnShouldTriggerError(t *testing.T) {
>>>>>>>  add instance of packetConn if it is nil
	myIcmp := Packet{}
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(nil, myPacket)
	if err == nil {
<<<<<<< HEAD
		t.Error("Pinging a nil IP should raise an error got", err)est failed again
=======
		t.Error("Pinging a nil IP should raise an error got", err)
>>>>>>>  code and test fixed
=======
func TestPingNoDestination(t *testing.T) {
	ip, err := net.ResolveIPAddr("ip", "")
	if err == nil {
		duration, error := Ping(ip)
		if error == nil {
			t.Error("Expected no error, got", error)
		} else if duration > 0 {
			t.Error("Expected duration < 0, got", duration)
		}
>>>>>>> add PingPoller_test.go
=======
}


func TestPingingWithNoIpShouldTriggerError(t *testing.T) {
	_, err := Ping(nil)
	if (err == nil) {
		t.Error("Pinging a nil IP should raise an error got", err);
>>>>>>> test failed again
	}
}
