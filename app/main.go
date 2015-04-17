package main

import (
	"./api/model"
	"./poller"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	response := poller.Response{}
	listOfDestination := model.Ips{}
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
					poller.DoPostOn(&response,"http://localhost:8000/ips/results")
				} else {
					duration, error := poller.Ping(ip)
					if error != nil {
						response.Destination = listOfDestination[i].Destination
						response.Status = "failed"
						response.Time = duration
						response.Error = err
						poller.DoPostOn(&response,"http://localhost:8000/ips/results")
					}
					if duration <= 0 {
						response.Destination = listOfDestination[i].Destination
						response.Status = "failed"
						response.Time = duration
						response.Error = err
						poller.DoPostOn(&response,"http://localhost:8000/ips/results")
					}
					response.Destination = listOfDestination[i].Destination
					response.Status = "good"
					response.Time = duration
					response.Error = err
					poller.DoPostOn(&response,"http://localhost:8000/ips/results")
				}

			}

		}
		time.Sleep(10000000000)
		log.Println("===============================")
	}
}
