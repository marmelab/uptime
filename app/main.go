package main

import (
	"fmt"
	"time"
	"./poller"
	"net"
)

func main() {
	var ip *net.IPAddr
	var err error
	var duration int
	for true{
		listOfIp := poller.RetrieveIpsFromJsonFile("/usr/src/watcher/app/poller/ips.json")
		for key,value := range listOfIp["ips"] {
			ip,err = poller.FromDomainNameToIp(value)
				if(err==nil){
					duration,err = poller.Ping(ip)
					if(duration >= 0){
						if(err ==nil){
							fmt.Println("It works ! Ping on ", value,"key : ",key, " time : ", duration)
						}else{
						fmt.Println(" it failed....", err)						}
					}else {
						fmt.Println(" it failed....", err)
					}
				}else{
					fmt.Println("it failed... error wrong name : ",err)
				}
			}
			time.Sleep(1000000000)
			fmt.Println("===============================")
		}
}