package main

import (
	"./api/model"
	"./poller"
	"encoding/json"
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	"golang.org/x/net/icmp"
=======
>>>>>>> add PingPoller_test.go
=======
	"golang.org/x/net/icmp"
>>>>>>>  code and test fixed
=======
	"golang.org/x/net/icmp"
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
>>>>>>> add PingPoller_test.go
=======
	"golang.org/x/net/icmp"
=======
>>>>>>> add PingPoller_test.go
=======
	"golang.org/x/net/icmp"
>>>>>>>  code and test fixed
>>>>>>>  code and test fixed
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	response := poller.Response{}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
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
<<<<<<< HEAD
			}
			for i := 0; i < len(listOfDestination); i++ {
				c <- listOfDestination[i].Destination
			}
=======
	listOfDestination := model.Ips{}
=======
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
	var listOfDestination []model.Ip
>>>>>>>  code and test fixed
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
=======
<<<<<<< HEAD
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
			}
			for i := 0; i < len(listOfDestination); i++ {
				c <- listOfDestination[i].Destination
			}
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
	listOfDestination := model.Ips{}
=======
	var listOfDestination []model.Ip
>>>>>>>  code and test fixed
<<<<<<< HEAD
>>>>>>>  code and test fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
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
<<<<<<< HEAD
			}
			err = json.Unmarshal(body, &listOfDestination)
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < len(listOfDestination); i++ {
				c <- listOfDestination[i].Destination
			}
>>>>>>> add PingPoller_test.go
>>>>>>> add PingPoller_test.go
=======
			}
			err = json.Unmarshal(body, &listOfDestination)
			if err != nil {
				log.Fatal(err)
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
			}
			for i := 0; i < len(listOfDestination); i++ {
				c <- listOfDestination[i].Destination
			}
<<<<<<< HEAD
>>>>>>> add PingPoller_test.go
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
		}
		time.Sleep(time.Second * 10)
		fmt.Println("===============================")
	}
}
