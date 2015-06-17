package repositories

import (
	"../../poller"
	"../target"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	"runtime"
)

var db *sql.DB

func GetDb() (db *sql.DB, err error) {
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), "../../conf.json")
	if db == nil {
		configdb := poller.RetrieveConfDbFromJsonFile(path)["database"]
		database := configdb.(map[string]interface{})
		db, err = sql.Open("postgres", "host="+database["host"].(string)+" user="+database["user"].(string)+" dbname="+database["dbname"].(string)+" sslmode="+database["sslmode"].(string)+"")
	}
	return db, err
}

func AddTarget(db *sql.DB, newTarget target.Target_data) (target.Target_data, error) {
	var result target.Target_data
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if newTarget.Destination == "" {
		error := errors.New("newTarget = nil ")
		return result, error
	}
	_ = db.QueryRow("INSERT INTO Destination (destination) VALUES($1) RETURNING id", newTarget.Destination).Scan(&result.Id)
	return result, nil
}

func GetTarget(db *sql.DB, id int) (target.Target_data, error) {
	var target_data target.Target_data
	if db == nil {
		error := errors.New("db = nil ")
		return target_data, error
	}
	if id <= 0 {
		error := errors.New("id invalid ")
		return target_data, error
	}
	row, err := db.Query("SELECT * from destination WHERE id = $1", id)
	if err != nil {
		return target_data, err
	}
	for row.Next() {
		_ = row.Scan(&target_data.Id, &target_data.Destination)
	}
	return target_data, nil
}

func GetTargetsWithLastResult(db *sql.DB) (*sql.Rows, error) {
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

func UpdateTarget(db *sql.DB, newTarget target.Target_data, oldTargetId int) error {
	if db == nil {
		error := errors.New("db = nil ")
		return error
	}
	if newTarget.Destination == "" || oldTargetId <= 0 {
		error := errors.New("newTarget = nil or oldTarget is wrong")
		return error
	}
	_, err := db.Exec("UPDATE destination SET destination = $1 WHERE id = $2", newTarget.Destination, oldTargetId)
	return err
}

func DeleteTarget(db *sql.DB, target_dataId int) (target.Target_data, error) {
	var result target.Target_data
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if target_dataId <= 0 {
		error := errors.New("target_dataId is wrong ")
		return result, error
	}
	_ = db.QueryRow("DELETE FROM Destination WHERE id = $1 RETURNING *", target_dataId).Scan(&result.Id)
	return result, nil
}
