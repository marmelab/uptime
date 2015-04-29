package main

import (
	"../poller"
	"./model"
	"database/sql"
	"encoding/json"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8000", "port for the api listen")
	flag.Parse()
	db, err := sql.Open("postgres", "host=db user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("error open db")
	}
	http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return
		}
		rows, _ := db.Query("Select * from Destination")
		leng, _ := db.Query("Select Count(*) from Destination")
		var length int
		for leng.Next() {
		_:
			leng.Scan(&length)
		defer rows.Close()
		ips := make([]model.Ip, length)
		i := 0
		for rows.Next() {
			var dest string
			error := rows.Scan(&dest)
			if error != nil {
				log.Print(error)
			}
			log.Print(dest)
			ips[i].Destination = dest
			i++
		}
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
		_,_=db.Exec("INSERT INTO Results (destination,status,time) VALUES($1,$2,$3)",newResult.Destination,newResult.Status,newResult.Time)
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
