package main

import (
	"./api/model"
	"./poller"
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"golang.org/x/net/icmp"
=======
>>>>>>> add PingPoller_test.go
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	response := poller.Response{}
<<<<<<< HEAD
	var listOfDestination []model.Ip
	var duration int
	var url string
	c := make(chan string, len(listOfDestination))
	packetConn, _ := icmp.ListenPacket("ip4:icmp", "")
	go func() {
		for {
			select {
			case url = <-c:
				ip, err := poller.FromDomainNameToIp(url)
				response.Destination = url
				if err == nil {
					duration, err = poller.Ping(ip, packetConn)
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
				break
			}
		}
	}()
	for {
		res, err := http.Get(poller.RetrieveIpsFromJsonFile("url.json")["urlApiIp"])
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
			for i := 0; i < len(listOfDestination); i++ {
				c <- listOfDestination[i].Destination
			}
=======
	listOfDestination := model.Ips{}
	var duration int
	var url string
	for true {
		res, err := http.Get(poller.RetrieveIpsFromJsonFile("url.json")["urlApiIp"])
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
				for i := 0; i < len(listOfDestination); i++ {
					c <- listOfDestination[i].Destination
				}
			}()
			for i := 0; i < len(listOfDestination); i++ {
				url = <-c
				ip, err := poller.FromDomainNameToIp(url)
				response.Destination = listOfDestination[i].Destination
				if err == nil {
					duration, err = poller.Ping(ip)
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
>>>>>>> add PingPoller_test.go
		}
		time.Sleep(time.Second * 10)
		fmt.Println("===============================")
	}
}
