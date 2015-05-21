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
	"errors"
)

func setAllowCORS(w *http.ResponseWriter) (error) {
	if w != nil {
		(*w).Header().Set("Access-Control-Allow-Origin","*")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(*w).Header().Set("Access-Control-Allow-Headers","Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		return nil
	} else {
		error := errors.New("http.ResponseWriter = nil ")
		return error
	}
}

func return500(w *http.ResponseWriter) (error) {
	if w != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		(*w).Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return nil	
	} else {
		error := errors.New("http.ResponseWriter = nil ")
		return error
	}

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
		setAllowCORS(&w)
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return
		}

		rows, _ := db.Query(`
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
		`);

		defer rows.Close()
		leng, _ := db.Query("SELECT COUNT(*) FROM destination")
		defer leng.Close()
		var length int
		for leng.Next() {
			_  = leng.Scan(&length)
		}
		ips := make([]target.Target_data, length)
		i := 0
		for rows.Next() {
			var id int
			var dest string
			var status bool
			error := rows.Scan(&id, &dest, &status)
			if error != nil {
				return500(&w)
				return
			}

			ips[i] = target.Target_data{ Id: id, Destination: dest, Status: status }
			i++
		}
		json.NewEncoder(w).Encode(ips)
	})

	http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
		setAllowCORS(&w)
		if (r.Method != "POST") && (r.Method != "GET") {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		if r.Method == "GET" {
			rows, _ := db.Query("SELECT * FROM results")
			defer rows.Close()
			leng, _ := db.Query("SELECT COUNT(destination) FROM results")
			defer leng.Close()
			var length int
			for leng.Next() {
			_:
				leng.Scan(&length)
			}
			res := make([]poller.Response, length)
			i := 0
			for rows.Next() {
				var dest string
				var sta string
				var tim int
				error := rows.Scan(&dest, &sta, &tim)
				if error != nil {
					return500(&w)
					return
				}
				res[i].Destination = dest
				res[i].Status = sta
				res[i].Time = tim
				i++
			}
			json.NewEncoder(w).Encode(res)
		} else {
			decoder := json.NewDecoder(r.Body)
			newResult := poller.Response{}
			error := decoder.Decode(&newResult)
			if error != nil {
				return500(&w)
				return
			}
			_, _ = db.Exec("INSERT INTO Results (destination, status, time) VALUES($1, $2, $3)", newResult.Destination, newResult.Status, newResult.Time)
		}
	})

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
