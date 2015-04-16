package main

import (
	"./poller"
	"fmt"
	"net"
	"time"
)

func main() {
	response := poller.Response{}
	var ip *net.IPAddr
	var err error
	var duration int
	for true {
		listOfIp := poller.RetrieveIpsFromJsonFile("/usr/src/watcher/app/poller/ips.json")
		for key, value := range listOfIp {
			ip, err = poller.FromDomainNameToIp(value)
			if err == nil {
				duration, err = poller.Ping(ip)
				if duration >= 0 {
					if err == nil {
						response.Destination = value
						response.Status = "good"
						response.Time = duration
						response.Key = key
						response.Error = err
						fmt.Println(response)
					} else {
						response.Destination = value
						response.Status = "failed"
						response.Time = duration
						response.Key = key
						response.Error = err
						fmt.Println(response)
					}
				} else {
					response.Destination = value
					response.Status = "failed"
					response.Time = duration
					response.Key = key
					response.Error = err
					fmt.Println(response)
				}
			} else {
				response.Destination = value
				response.Status = "failed"
				response.Time = duration
				response.Key = key
				response.Error = err
				fmt.Println(response)
			}
		}
		time.Sleep(10000000000)
		fmt.Println("===============================")
	}
}
