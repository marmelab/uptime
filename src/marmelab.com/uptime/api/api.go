package main

import (
	"../poller"
	"./model"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"flag"
	_"github.com/lib/pq"
	"database/sql"
)

func main() {
	port := flag.String("port","8000","port for the api listen")
	flag.Parse()
<<<<<<< HEAD
	db,err := sql.Open("postgres","user=postgres dbname=uptime sslmode=verify-full")
	http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return
=======
	db,err := sql.Open("postgres","host=db user=postgres dbname=postgres sslmode=disable")
	if(err!=nil){
		log.Fatal("error open db")
	}
	http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET"{
			w.WriteHeader(http.StatusNotFound)
		}
		if r.Method == "GET" {
			raw,error:= db.Query("SELECT * FROM Destination")
			if(error!=nil){
				log.Print(error)
			}
			log.Print(raw)
			var ips [2]model.Ip
			ips[0].Destination = "google.fr"
			ips[1].Destination = "failfailfail.fail"
			json.NewEncoder(w).Encode(ips)
>>>>>>>  link db to api and connect to db via api
		}
		var ips [2]model.Ip
		ips[0].Destination = "google.fr"
		ips[1].Destination = "failfailfail.fail"
		json.NewEncoder(w).Encode(ips)
	})
	http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
		decoder := json.NewDecoder(r.Body)
		newResult := poller.Response{}
		error := decoder.Decode(&newResult)
		if error != nil {
			log.Print(error)
			return
		}
		log.Print(newResult)
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
