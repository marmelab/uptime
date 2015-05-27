package main

import (
	"./api/target"
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
	var listOfDestination []target.Target_data
	var url string
	var duration int	
	for true {
		conf := poller.RetrieveConfDbFromJsonFile("/usr/src/watcher/src/marmelab.com/uptime/conf.json")
		res, err := http.Get(conf["urlApiIp"].(string))
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
				response.Destination = url
					duration, err = poller.HttpPing(url,false)
					response.Time = duration
					if (err != nil) || (duration <= 0) {
						response.Status = "failed"
						if err != nil {
							log.Print(err)
						}
					} else {
						response.Status = "good"
						if err != nil {
							log.Print(err)
						}
					}
				err = poller.DoPostOn(&response, conf["urlApiResults"].(string))
			}
			time.Sleep(time.Second * 10)
			fmt.Println("===============================")
		}
	}
}
