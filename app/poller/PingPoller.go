package poller

import (
	"golang.org/x/net/icmp"
	"net"
	"time"
)


func FromDomainNameToIp(domainName string) (ip *net.IPAddr,err error){
	ip, err = net.ResolveIPAddr("ip", domainName) 
	return ip, err
}

func Ping(ip *net.IPAddr) (int, error) {
	var duration int
	var data []byte
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

	return duration / 1000, err
}




