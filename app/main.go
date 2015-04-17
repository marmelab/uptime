package main

import (
	"./api/model"
	"./poller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	response := poller.Response{}
	listOfDestination := api.Ips{}
	for true {
		res, err := http.Get("http://localhost:8000/ips/")
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
				ip, err := poller.FromDomainNameToIp(listOfDestination[i].Destination)
				if err != nil {
					response.Destination = listOfDestination[i].Destination
					response.Status = "failed"
					response.Time = -1
					response.Error = err
					fmt.Println(response)
				} else {
					duration, error := poller.Ping(ip)
					if error != nil {
						response.Destination = listOfDestination[i].Destination
						response.Status = "failed"
						response.Time = duration
						response.Error = err
						fmt.Println(response)
					}
					if duration <= 0 {
						response.Destination = listOfDestination[i].Destination
						response.Status = "failed"
						response.Time = duration
						response.Error = err
						fmt.Println(response)
					}
					response.Destination = listOfDestination[i].Destination
					response.Status = "good"
					response.Time = duration
					response.Error = err
					fmt.Println(response)
				}

			}

		}
		time.Sleep(10000000000)
		fmt.Println("===============================")
	}
}
