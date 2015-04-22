package main

import (
	"../poller"
	"./model"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			listIp := model.Ips{
				model.Ip{Destination: "google.fr"},
				model.Ip{Destination: "youtube.fr"},
				model.Ip{Destination: "bing.fr"},
				model.Ip{Destination: "szszdzdadafdff.fr"},
			}
			json.NewEncoder(w).Encode(listIp)
		}

	})
	http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
		}
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			newResultat := poller.Response{}
			error := decoder.Decode(&newResultat)
			if error != nil {
				log.Fatal(error)
			}
			log.Print(newResultat)
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
