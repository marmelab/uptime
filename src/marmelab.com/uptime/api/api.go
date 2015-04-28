package main

import (
	"../poller"
	"./model"
	"encoding/json"
	"log"
	"net/http"
	"flag"
	_"github.com/lib/pq"
	"database/sql"
)

func main() {
	port := flag.String("port","8000","port for the api listen")
	flag.Parse()
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
		}

	})
	http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			newResultat := poller.Response{}
			error := decoder.Decode(&newResultat)
			if error != nil {
				w.Header().Set("Statuscode","500")
				log.Fatal(error)
			}
			log.Print(newResultat)
		}
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
