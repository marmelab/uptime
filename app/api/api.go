package main

import (
	"../poller"
	"./model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db,errordb := sql.Open("postgres","user=postgres dbname=uptime sslmode=verify-full")
	if(errordb!=nil){
		log.Fatal(errordb)
	}
	http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			rows,errorrows := db.Query("Select * from Destination")
			if(errorrows!=nil){
				log.Print(errorrows)
			}
			log.Print(rows)
			listIp := model.Ips{
				model.Ip{Destination: "google.fr"},
				model.Ip{Destination: "youtube.fr"},
				model.Ip{Destination: "bing.fr"},
				model.Ip{Destination: "szszdzdadafdff.fr"},
			}
			json.NewEncoder(w).Encode(listIp)
		}
		if r.Method == "POST" {
			newDestination := model.Ip{}
			body, err := ioutil.ReadAll(r.Body)
			r.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(body, &newDestination)
			if err != nil {
				log.Fatal(err)
			}
			listIp := model.Ips{
				model.Ip{Destination: "google.fr"},
				model.Ip{Destination: "youtube.fr"},
				model.Ip{Destination: "bing.fr"},
				model.Ip{Destination: "szszdzdadafdff.fr"},
				model.Ip{Destination: newDestination.Destination},
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
