package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func initDb(){
	conf := poller.RetrieveConfDbFromJsonFile("/usr/src/api/src/marmelab.com/uptime/conf.json")
	configdb := conf["database"]
	database := configdb.(map[string]interface{})
	db, err := sql.Open("postgres", "host="+database["host"].(string)+" user="+database["user"].(string)+" dbname="+database["dbname"].(string)+" sslmode="+database["sslmode"].(string)+"")
	return db
}

func AddTarget(db *DB){
	decoder := json.NewDecoder(r.Body)
	var newTarget string
	error := decoder.Decode(&newTarget)
	if error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	_, _ = db.Exec("INSERT INTO Destination (destination) VALUES($1)", newTarget)

}

func GetTarget(){

}

func GetTargets(db *sql.DB) {
	rows, QueryError := db.Query(`
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
	if QueryError != nil {
		log.Print("request error ", QueryError)
	}
	defer rows.Close()
	return rows	
}

func UpdateTarget() {
			if r.Method == "PUT" {
			decoder := json.NewDecoder(r.Body)
			var newTarget string
			var oldTarget string
			error := decoder.Decode(&newTarget, &oldTarget)
			if error != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			_, _ = db.Exec("UPDATE Destination SET destination = $1 WHERE destination = $2",newTarget, oldTarget)
		}
}

func DeleteTarget() {
			if r.Method == "DELETE" {
			decoder := json.NewDecoder(r.Body)
			var target string
			error := decoder.Decode(&target)
			if error != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			_, _ = db.Exec("DELETE FROM Destination WHERE destination = $1",target)
		}
}
