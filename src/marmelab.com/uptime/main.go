package main

import (
	"./poller"
	"fmt"
	"net"
	"time"
	"golang.org/x/net/icmp"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"./api/model"
	"os"
)

func main() {
	response := poller.Response{}
	var listOfDestination []model.Ip
	var ip *net.IPAddr
	var url string
	var duration int
	packetConn, _ := icmp.ListenPacket("ip4:icmp", "")
	for true {
		res, err := http.Get(poller.RetrieveIpsFromJsonFile(os.Getenv("CONFIG_PATH"))["urlApiIp"])
		if err != nil {
			log.Fatal(err)
		} else {
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(body, &listOfDestination)
			if err != nil {
				log.Fatal(err)
			}
		c := make(chan string, len(listOfDestination))
		go func() {
			for _, value := range listOfDestination {
				c <- value.Destination
			}
		}()
		for range listOfDestination {
			url = <-c
			ip, err = poller.FromDomainNameToIp(url)
			response.Destination = url
			if err == nil {
				duration, err = poller.Ping(ip,packetConn)
				response.Time = duration
				response.Error = err
				if (err != nil) || (duration <= 0) {
					response.Status = "failed"
				err = poller.DoPostOn(&response, poller.RetrieveIpsFromJsonFile("url.json")["urlApiResults"])
					if err != nil {
						log.Print(err)
					}
				} else {
					response.Status = "good"
				err = poller.DoPostOn(&response, poller.RetrieveIpsFromJsonFile("url.json")["urlApiResults"])
					if err != nil {
						log.Print(err)
					}
				}
			} else {
				response.Status = "failed"
				response.Time = duration
				response.Error = err
				err = poller.DoPostOn(&response, poller.RetrieveIpsFromJsonFile("url.json")["urlApiResults"])
					if err != nil {
						log.Print(err)
					}
			}
		}
		time.Sleep(time.Second * 10)
		fmt.Println("===============================")
	}
}
}
