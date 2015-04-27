package poller

import (
	"golang.org/x/net/icmp"
	"net"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>> tests fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
	"net/http"
	"net/http/httptest"
>>>>>>> add PingPoller_test.go
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> tests fixed
=======
>>>>>>> add PingPoller_test.go
=======
=======
>>>>>>> tests fixed
>>>>>>> tests fixed
=======
=======
>>>>>>> tests fixed
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	"testing"
)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
>>>>>>>  code and test fixed
=======
>>>>>>> test failed again
=======
=======
>>>>>>>  code and test fixed
>>>>>>>  code and test fixed
=======
=======
>>>>>>>  code and test fixed
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
<<<<<<< HEAD
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>>  code and test fixed
=======
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
=======
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
>>>>>>> test failed again
=======
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
	ip, _ := net.ResolveIPAddr("ip", "localhost")
>>>>>>> tests fixed
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> tests fixed
=======
=======
>>>>>>> tests fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
func TestPingWithValidIpShouldNotTrigger(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ip, err := net.ResolveIPAddr("ip", ts.URL)
>>>>>>> test failed again
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> test failed again
=======
=======
>>>>>>> tests fixed
>>>>>>> tests fixed
=======
=======
>>>>>>> tests fixed
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
		duration, error := Ping(ip)
		if error != nil {
			t.Error("Expected no error, got", error)
		} else if duration <= 0 {
			t.Error("Expected duration > 0, got", duration)
		}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
>>>>>>> test failed again
=======
=======
>>>>>>>  code and test fixed
=======
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
>>>>>>>  code and test fixed
=======
type Packet struct {
>>>>>>>  code and test fixed
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
<<<<<<< HEAD
=======
=======
=======
type Packet struct {
>>>>>>>  code and test fixed
=======
type Packet struct {
>>>>>>> d09aa0ab1d64898cc222c40cd23103b41a95c1e5
}

>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
type IcmpMock interface {
	ListenPacket(network string, address string) *icmp.PacketConn
}

func (pack Packet) ListenPacket(n string, a string) *icmp.PacketConn {
	var p icmp.PacketConn
	ptr := &p
	return ptr
<<<<<<< HEAD
}

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

type NetMock interface {
	ResolveIPAddr(proto string, address string) (*net.IPAddr, error)
<<<<<<< HEAD
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
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
}
<<<<<<< HEAD

<<<<<<< HEAD
<<<<<<< HEAD
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
type Net struct {
}
<<<<<<< HEAD
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
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
<<<<<<< HEAD
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
<<<<<<< HEAD
>>>>>>>  code and test fixed
=======
>>>>>>> add PingPoller_test.go
=======
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
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
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
>>>>>>> add PingPoller_test.go
	}
=======
=======
<<<<<<< HEAD
>>>>>>> test failed again
=======
=======
=======
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df

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
<<<<<<< HEAD
>>>>>>>  code and test fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
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
<<<<<<< HEAD
=======
=======
	if err == nil {
		t.Error("Pinging a nil IP should raise an error got", err)
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	}
}
