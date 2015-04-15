package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/icmp"
	"net"
	"time"
)

type Request struct{
	doPing bool
	timeBetweenPing time.Duration
}

func Ping(domainName string) int {
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
				return duration / 1000
			}
			if err != nil {
				fmt.Println("error on WriteTo")
				fmt.Println(err)
			}
		} else {
			fmt.Println("error on ListenPacket")
			fmt.Println(err)
		}
	} else {
		fmt.Println("error on domainName")
		fmt.Println(err)
	}

	return duration / 1000
}

func (request *Request) setdoPing(doPing bool) {
	request.doPing = doPing
}

func (request *Request) settimeBetweenPing(timeBetweenPing time.Duration) {
	request.timeBetweenPing = timeBetweenPing
}

func main() {
	request :=  Request{}
	var duration int
	dst := flag.String("dst", "8.8.8.8", "destination to ping")
	flag.Parse()
	fmt.Println("Ping on : " + *dst)
	request.setdoPing(true)
	request.settimeBetweenPing(1000000000)
	for(request.doPing){
	duration = Ping(*dst)
		if duration != 0 {
			fmt.Println("It works ! Time : ")
			fmt.Println(duration)
		} else {
			fmt.Println("It failed...")
		}
		time.Sleep(request.timeBetweenPing)
	}

}
