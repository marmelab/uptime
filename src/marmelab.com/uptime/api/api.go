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
	"time"
)

func SetCors(w *http.Header) {
	w.Set("Access-Control-Allow-Origin", "*")
	w.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	port := flag.String("port", "8383", "port for the api listen")
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
		header := w.Header()
		SetCors(&header)
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return
		}

		rows, errr := db.Query(`
			WITH last_results AS (
				SELECT *, ROW_NUMBER() OVER(
					PARTITION BY destination
					ORDER BY created_at DESC
				) AS rank
				FROM results
			)
			SELECT D.id, D.destination, LR.status = 'good'
			FROM destination D
			LEFT JOIN last_results LR ON (D.destination = LR.destination AND rank = 1);
		`)
		if(errr != nil) {
			log.Print("request error ", errr)
		}
		defer rows.Close()
		ips := make([]target.Target_data, 0)
		for rows.Next() {
			var id int
			var dest string
			var status bool
			error := rows.Scan(&id, &dest, &status)
			log.Print(error)
			if error != nil {
				log.Print(error)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			ips = append(ips, target.Target_data{Id: id, Destination: dest, Status: status})
		}
		json.NewEncoder(w).Encode(ips)
	})

	http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		SetCors(&header)
		if (r.Method != "POST") && (r.Method != "GET") {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		if r.Method == "GET" {
			rows, _ := db.Query("SELECT destination, status, duration,  created_at, target_id FROM results")
			defer rows.Close()
			res := make([]poller.Response, 0)
			for rows.Next() {
				var dest string
				var sta string
				var tim int
				var created_at time.Time
				var target_id int
				error := rows.Scan(&dest, &sta, &tim, &created_at, &target_id)
				if error != nil {
					log.Print(error)
					http.Error(w, http.StatusText(500), 500)
					return
				}
				res = append(res,poller.Response{Destination: dest, Status: sta, Time: tim, Created_at: created_at, Target_id: target_id})
			}
			json.NewEncoder(w).Encode(res)
		} else {
			decoder := json.NewDecoder(r.Body)
			newResult := poller.Response{}
			error := decoder.Decode(&newResult)
			if error != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			log.Print(newResult.Target_id)
			_, _ = db.Exec("INSERT INTO Results (destination, status, duration, target_id) VALUES($1, $2, $3, $4)", newResult.Destination, newResult.Status, newResult.Time, newResult.Target_id)
		}
	})

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
