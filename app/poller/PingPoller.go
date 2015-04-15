package poller

import (
	"golang.org/x/net/icmp"
	"net"
	"time"
)

type Request struct {
	DoPing          bool
	TimeBetweenPing time.Duration
}

func Ping(domainName string) (int, error) {
	var duration int
	var data []byte
	ip, err := net.ResolveIPAddr("ip", domainName)
	if err == nil {
		var destination net.Addr = ip
		packetConn, err := icmp.ListenPacket("ip4:icmp", "")
		if err == nil {
			timeNow := time.Now().Nanosecond()
			errorCode, err := packetConn.WriteTo(data, destination)
			duration = time.Now().Nanosecond() - timeNow
			if errorCode == 0 {
				return duration / 1000, err
			}
			if err != nil {
				return duration, err
			}
		} else {
			return duration, err
		}
	} else {
		return duration, err
	}

	return duration / 1000, err
}

func (request *Request) SetdoPing(doPing bool) {
	request.DoPing = doPing
}

func (request *Request) SettimeBetweenPing(timeBetweenPing time.Duration) {
	request.TimeBetweenPing = timeBetweenPing
}


