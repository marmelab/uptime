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
	//c := make(chan string)
	for true {
		listOfUrl := poller.RetrieveIpsFromJsonFile("url.json")
		listOfIp := poller.RetrieveIpsFromJsonFile(listOfUrl["urlIps"])
		for key, value := range listOfIp {
			ip, err = poller.FromDomainNameToIp(value)
			response.Destination = value
			response.Key = key
			if err == nil {
				duration, err = poller.Ping(ip)
				response.Time = duration
				response.Error = err
				if (err != nil) || (duration <= 0) {
					response.Status = "failed"
					fmt.Println(response)
				} else {
					response.Status = "good"
					fmt.Println(response)
				}
			} else {
				response.Status = "failed"
				response.Time = duration
				response.Error = err
				fmt.Println(response)
			}
		}
		time.Sleep(10000000000)
		fmt.Println("===============================")
	}
}
