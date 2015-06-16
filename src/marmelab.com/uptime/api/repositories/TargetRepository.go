package repositories

import (
	"../../config"
	"../target"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func GetDb() (db *sql.DB, err error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/usr/src/api/src/marmelab.com/uptime/conf.json"
	}
	config, configErr := config.GetConfig(configPath)
	if configErr != nil {
		return nil, configErr
	}
	if config["database"] != nil {
		configdb := config["database"].(map[string]interface{})
		if db == nil {
			db, err := sql.Open("postgres", "host="+configdb["host"].(string)+" user="+configdb["user"].(string)+" dbname="+configdb["dbname"].(string)+" sslmode="+configdb["sslmode"].(string)+"")
			return db, err
		}
	} else {
		if db == nil {
			db, err := sql.Open("postgres", "host="+config["host"].(string)+" user="+config["user"].(string)+" dbname="+config["dbname"].(string)+" sslmode="+config["sslmode"].(string)+"")
			return db, err
		}
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
	queryError := db.QueryRow("INSERT INTO Destination (destination) VALUES($1) RETURNING id, destination", newTarget.Destination).Scan(&result.Id, &result.Destination)
	if queryError != nil {
		return result, queryError
	}
	return result, nil
}

func GetTarget(db *sql.DB, id int) (target.Target_data, error) {
	var target_data target.Target_data
	var errScan error
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
		errScan = row.Scan(&target_data.Id, &target_data.Destination)
		if errScan != nil {
			return target_data, errScan
		}
	}
	return target_data, nil
}

func GetTargetsWithLastResult(db *sql.DB, page int, perPage int) (*sql.Rows, error) {
	if db == nil {
		error := errors.New("db = nil ")
		return nil, error
	}
	rows, queryError := db.Query(`
		WITH last_results AS (
			SELECT *, ROW_NUMBER() OVER(
				PARTITION BY destination
				ORDER BY created_at DESC
			) AS rank
			FROM results
		)
		SELECT D.id, D.destination, LR.status = 'good' AS reachable
		FROM destination D
		LEFT JOIN last_results LR ON (D.destination = LR.destination AND rank = 1);
	`)
	if queryError != nil {
		log.Print("request error ", queryError)
		return nil, queryError
	}
	return rows, nil
}

func UpdateTarget(db *sql.DB, newTarget target.Target_data, oldTargetId int) (target.Target_data, error) {
	var result target.Target_data
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if newTarget.Destination == "" || oldTargetId <= 0 {
		error := errors.New("newTarget = nil or oldTarget is wrong")
		return result, error
	}
	_, err := db.Exec("UPDATE destination SET destination = $1 WHERE id = $2 ", newTarget.Destination, oldTargetId)
	result.Destination = newTarget.Destination
	result.Id = oldTargetId
	return result, err
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
	queryError := db.QueryRow("DELETE FROM Destination WHERE id = $1 RETURNING id, destination", target_dataId).Scan(&result.Id, &result.Destination)
	if queryError != nil {
		return result, queryError
	}
	return result, nil
}
