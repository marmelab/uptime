package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"net"
	"time"
	"flag"
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
				return duration/1000
			}
			if err != nil {
				fmt.Println("error on writeTo")
			}
		} else {
			fmt.Println("error on ListenPacket")
		}
	}else {
		fmt.Println("error on domainName")
	}

	return duration/1000
}

func main() {
	dst := flag.String("dst","8.8.8.8","destination to ping")
	flag.Parse()
	fmt.Println("Ping on : " + *dst)
	duration := Ping(*dst)
	fmt.Println("It works ! Time : ")
	fmt.Println(duration)

}
