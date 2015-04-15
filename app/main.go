package main

import (
	"flag"
	"fmt"
	"time"
	"./poller"
	"net"
)

func main() {
	var duration int
	var err error
	var ip *net.IPAddr
	dst := flag.String("dst", "8.8.8.8", "destination to ping")
	flag.Parse()
	fmt.Println("Ping on : " + *dst)
	for true {
		ip, err = poller.FromDomainNameToIp(*dst)
		if(err==nil){
			duration, err = poller.Ping(ip)
			if duration != 0 || err == nil {
				fmt.Println("It works ! Time : ")
				fmt.Println(duration)
			} else {
				fmt.Println("It failed...")
				fmt.Println(err)
				break
			}
			time.Sleep(1000000000)
		}
	}

}