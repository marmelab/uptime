package poller

import (
	"golang.org/x/net/icmp"
	"net"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"net/http"
	"net/http/httptest"
>>>>>>> add PingPoller_test.go
=======
>>>>>>> tests fixed
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
	"testing"
)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
<<<<<<< HEAD
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
=======
	ip, _ := net.ResolveIPAddr("ip", "localhost")
>>>>>>> tests fixed
<<<<<<< HEAD
<<<<<<< HEAD
=======
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
>>>>>>> test failed again
=======
>>>>>>> tests fixed
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected no error, got", error)
		} else if duration <= 0 {
			t.Error("Expected duration > 0, got", duration)
		}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
type Packet struct {
>>>>>>>  code and test fixed
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
=======
type Packet struct {
>>>>>>>  code and test fixed
=======
type Packet struct {
>>>>>>> d09aa0ab1d64898cc222c40cd23103b41a95c1e5
}

>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
type IcmpMock interface {
	ListenPacket(network string, address string) *icmp.PacketConn
}

func (pack Packet) ListenPacket(n string, a string) *icmp.PacketConn {
	var p icmp.PacketConn
	ptr := &p
	return ptr
}
<<<<<<< HEAD

type Net struct {
}
<<<<<<< HEAD
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
	}
>>>>>>> add PingPoller_test.go
=======
<<<<<<< HEAD

type Net struct {
}
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19

type NetMock interface {
	ResolveIPAddr(proto string, address string) (*net.IPAddr, error)
}

<<<<<<< HEAD
<<<<<<< HEAD
func (ipad Net) ResolveIPAddr(p string, a string) (*net.IPAddr, error) {
	var i net.IPAddr
	i.Zone = "france"
	i.IP = net.ParseIP("localhost")
	ptr := &i
	return ptr, nil
}

=======
=======

type Net struct {
}

type NetMock interface {
	ResolveIPAddr(proto string, address string) (*net.IPAddr, error)
}

>>>>>>> d09aa0ab1d64898cc222c40cd23103b41a95c1e5
func (ipad Net) ResolveIPAddr(p string, a string) (*net.IPAddr, error) {
	var i net.IPAddr
	i.Zone = "france"
	i.IP = net.ParseIP("localhost")
	ptr := &i
	return ptr, nil
}

>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
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
<<<<<<< HEAD
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
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
=======

type NetMock interface {
	ResolveIPAddr(proto string, address string) (*net.IPAddr, error)
>>>>>>>  code and test fixed
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
func TestPingingWithNoIpShouldTriggerError(t *testing.T) {
<<<<<<< HEAD
	_, err := Ping(nil)
	if (err == nil) {
		t.Error("Pinging a nil IP should raise an error got", err);
>>>>>>> test failed again
=======
	myNet := Net{}
	myIp, _ := myNet.ResolveIPAddr("ip", "google.fr")
	_, err := Ping(myIp, nil)
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
	}
}
=======
>>>>>>>  add instance of packetConn if it is nil

func TestPingingWithNoIPConnShouldTriggerError(t *testing.T) {
	myIcmp := Packet{}
	myPacket := myIcmp.ListenPacket("ip4:icmp", "")
	_, err := Ping(nil, myPacket)
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
>>>>>>>  code and test fixed
=======
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
	}
}
