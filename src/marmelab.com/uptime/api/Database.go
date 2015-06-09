package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"../poller"
	"./target"
	"errors"
	"log"
)

var Db *sql.DB

func getDb() (*sql.DB, error) {
	if Db == nil{
		conf := poller.RetrieveConfDbFromJsonFile("/usr/src/api/src/marmelab.com/uptime/conf.json")
		configdb := conf["database"]
		database := configdb.(map[string]interface{})
		Db, err := sql.Open("postgres", "host="+database["host"].(string)+" user="+database["user"].(string)+" dbname="+database["dbname"].(string)+" sslmode="+database["sslmode"].(string)+"")
	}
	return Db, err
}

func AddTarget(db *sql.DB, newTarget string) (error){
	if(db == nil) {
		error := errors.New("db = nil ")
		return error		
	}
	if(newTarget == nil) {
		error := errors.New("newTarget = nil ")
		return error		
	}
	_, err := db.Exec("INSERT INTO Destination (destination) VALUES($1)", newTarget)
	return err;
}

func GetTarget(db *sql.DB, id int) (sql.Result, error) {
	if(db == nil) {
		error := errors.New("db = nil ")
		return nil, error		
	}
	if(id == nil) {
		error := errors.New("id = nil ")
		return nil, error		
	}
	row, err := db.Query("SELECT * from destination WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func GetTargets(db *sql.DB) (sql.Result, error) {
	if(db == nil) {
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

func UpdateTarget(db *sql.DB, newTarget string, oldTarget string) (error){
	db := getDb()
	if(db == nil) {
		error := errors.New("db = nil ")
		return error		
	}
	if(newTarget == nil || oldTarget == nil) {
		error := errors.New("string = nil ")
		return error		
	}	
	_, err := db.Exec("UPDATE Destination SET destination = $1 WHERE destination = $2",newTarget, oldTarget)
	return err
}

func DeleteTarget(db *sql.DB, target string) {
	db := getDb()
	if(db == nil) {
		error := errors.New("db = nil ")
		return error		
	}
	if(target == nil) {
		error := errors.New("target = nil ")
		return error		
	}		
	_, err := db.Exec("DELETE FROM Destination WHERE destination = $1",target)
	return err
}
