package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/icmp"
	"net"
	"time"
)

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

func main() {
	dst := flag.String("dst", "8.8.8.8", "destination to ping")
	flag.Parse()
	fmt.Println("Ping on : " + *dst)
	duration := Ping(*dst)
	if duration != 0 {
		fmt.Println("It works ! Time : ")
		fmt.Println(duration)
	} else {
		fmt.Println("It failed...")
	}

}
