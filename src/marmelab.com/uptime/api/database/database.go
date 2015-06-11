package database

import (
	"../../poller"
	Target "../target"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func GetDb() (Db *sql.DB, err error) {
	if Db == nil {
		conf := poller.RetrieveConfDbFromJsonFile("../../conf.json")
		configdb := conf["database"]
		database := configdb.(map[string]interface{})
		Db, err = sql.Open("postgres", "host="+database["host"].(string)+" user="+database["user"].(string)+" dbname="+database["dbname"].(string)+" sslmode="+database["sslmode"].(string)+"")
	}
	return Db, err
}

func AddTarget(db *sql.DB, newTarget string) error {
	if db == nil {
		error := errors.New("db = nil ")
		return error
	}
	if newTarget == "" {
		error := errors.New("newTarget = nil ")
		return error
	}
	_, err := db.Exec("INSERT INTO Destination (destination) VALUES($1)", newTarget)
	return err
}

func GetTarget(db *sql.DB, destination string) (Target.Target_data, error) {
	var target Target.Target_data
	if db == nil {
		error := errors.New("db = nil ")
		return target, error
	}
	if destination == "" {
		error := errors.New("destination = nil ")
		return target, error
	}
	row, err := db.Query("SELECT * from destination WHERE destination = $1", destination)
	if err != nil {
		return target, err
	}
	var dest string
	var id int
	for row.Next() {
		_ = row.Scan(&id, &dest)
	}
	target = Target.Target_data{Id: id, Destination: dest}
	return target, nil
}

func GetTargets(db *sql.DB) (*sql.Rows, error) {
	if db == nil {
		error := errors.New("db = nil ")
		return nil, error
	}
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
		return nil, QueryError
	}
	return rows, nil
}

func UpdateTarget(db *sql.DB, newTarget string, oldTarget string) error {
	if db == nil {
		error := errors.New("db = nil ")
		return error
	}
	if newTarget == "" || oldTarget == "" {
		error := errors.New("string = nil")
		return error
	}
	_, err := db.Exec("UPDATE destination SET destination = $1 WHERE destination = $2", newTarget, oldTarget)
	return err
}

func DeleteTarget(db *sql.DB, target string) error {
	if db == nil {
		error := errors.New("db = nil ")
		return error
	}
	if target == "" {
		error := errors.New("target = nil ")
		return error
	}
	_, err := db.Exec("DELETE FROM Destination WHERE destination = $1", target)
	return err
}
