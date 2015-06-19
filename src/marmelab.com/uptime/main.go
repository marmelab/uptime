package main

import (
	"marmelab.com/uptime/api/target"
	"marmelab.com/uptime/config"
	"marmelab.com/uptime/poller"
	"encoding/json"
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
	var id int
	for true {
		conf, configErr := config.GetConfig("./conf.json")
		if configErr != nil {
			log.Fatal(configErr)
		}

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
			c := make(chan *(target.Target_data))

			go func() {
				for true {
					t := <-c
					url = t.Destination
					id = t.Id
					log.Printf("consummed %v", url)
					response.Destination = url
					response.Target_id = id
					duration, err = poller.HttpPing(url, "http")
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
			}()
			for _, value := range listOfDestination {
				c <- &value
				log.Printf("sending %v", &value)
			}
			time.Sleep(time.Second * 10)
		}
	}
}
