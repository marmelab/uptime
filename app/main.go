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
	var url string
	var err error
	var duration int
	for true {
		listOfUrl := poller.RetrieveIpsFromJsonFile("url.json")
		listOfIp := poller.RetrieveIpsFromJsonFile(listOfUrl["urlIps"])
		c := make(chan string, len(listOfIp))
		go func() {
			for _, value := range listOfIp {
				c <- value
			}
		}()
		for range listOfIp {
			url = <-c
			ip, err = poller.FromDomainNameToIp(url)
			response.Destination = url
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
		/*fr key, value := range listOfIp {
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
		}*/
		time.Sleep(time.Second * 10)
		fmt.Println("===============================")
	}
}
