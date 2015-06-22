package repositories

import (
	"../../poller"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func AddResult(db *sql.DB, newResult poller.Response) (poller.Response, error) {
	var result poller.Response
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if newResult.Destination == "" {
		error := errors.New("newResult = nil ")
		return result, error
	}
	_ = db.QueryRow("INSERT INTO results (destination, status, duration) VALUES($1, $2, $3) RETURNING id", newResult.Destination, newResult.Status, newResult.Time).Scan(&result.Id)
	return result, nil
}

func GetResult(db *sql.DB, id int) (poller.Response, error) {
	var result poller.Response
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if id <= 0 {
		error := errors.New("id invalid ")
		return result, error
	}
	row, err := db.Query("SELECT * from results WHERE id = $1", id)
	if err != nil {
		return result, err
	}
	for row.Next() {
		_ = row.Scan(&result.Id, &result.Destination)
	}
	return result, nil
}

func GetResults(db *sql.DB) (*sql.Rows, error) {
	if db == nil {
		error := errors.New("db = nil ")
		return nil, error
	}
	rows, err := db.Query("SELECT * from results")
	if err != nil {
		return result, err
	}
	if QueryError != nil {
		log.Print("request error ", QueryError)
		return nil, QueryError
	}
	return rows, nil
}

func UpdateResults(db *sql.DB, newResult poller.Response, oldTargetId int) error {
	if db == nil {
		error := errors.New("db = nil ")
		return error
	}
	if newResult.Destination == "" || oldTargetId <= 0 {
		error := errors.New("newResult = nil or oldTarget is wrong")
		return error
	}
	_, err := db.Exec("UPDATE results SET destination = $1, status = $2, duration = $3 WHERE id = $4", newResult.Destination, newResult.Status, newResult.Time, oldTargetId)
	return err
}

func DeleteResult(db *sql.DB, target_dataId int) (poller.Response, error) {
	var result poller.Response
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if target_dataId <= 0 {
		error := errors.New("target_dataId is wrong ")
		return result, error
	}
	_ = db.QueryRow("DELETE FROM results WHERE id = $1 RETURNING *", target_dataId).Scan(&result.Id)
	return result, nil
}
