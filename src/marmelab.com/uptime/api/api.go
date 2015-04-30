package main

import (
	"../poller"
	"./target"
	"database/sql"
	"encoding/json"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "9000", "port for the api listen")
	flag.Parse()
	conf := poller.RetrieveConfDbFromJsonFile("/usr/src/api/src/marmelab.com/uptime/conf.json")
	configdb := conf["database"]
	database := configdb.(map[string]interface{})
	db, err := sql.Open("postgres", "host="+database["host"].(string)+" user="+database["user"].(string)+" dbname="+database["dbname"].(string)+" sslmode="+database["sslmode"].(string)+"")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return
		}
		rows, _ := db.Query("SELECT destination FROM destination")
		defer rows.Close()
		leng, _ := db.Query("SELECT COUNT(*) FROM destination")
		defer leng.Close()
		var length int
		for leng.Next() {
		_:
			leng.Scan(&length)
		}
		ips := make([]target.Ip, length)
		i := 0
		for rows.Next() {
			var dest string
			error := rows.Scan(&dest)
			if error != nil {
				log.Print(error)
				return
			}
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
		_, _ = db.Exec("INSERT INTO Results (destination, status, time) VALUES($1, $2, $3)", newResult.Destination, newResult.Status, newResult.Time)
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
