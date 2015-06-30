package repositories

import (
	"../../poller"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
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
	errorQuery := db.QueryRow("INSERT INTO results (target_id, destination, status, duration) VALUES($1, $2, $3, $4) RETURNING *", newResult.Target_id, newResult.Destination, newResult.Status, newResult.Time).Scan(&result.Id, &result.Target_id, &result.Destination, &result.Status, &result.Time, &result.Created_at)
	if errorQuery != nil {
		return result, errorQuery
	}
	return result, nil
}

func GetResult(db *sql.DB, id int) (poller.Response, error) {
	var result poller.Response
	var errScan error
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
		errScan = row.Scan(&result.Id, &result.Target_id, &result.Destination, &result.Status, &result.Time, &result.Created_at)
		if(errScan != nil) {
			return result, errScan
		}
	}
	return result, nil
}

func GetResults(db *sql.DB, page int, perPage int) (*sql.Rows, error) {
	if db == nil {
		error := errors.New("db = nil ")
		return nil, error
	}
	if page < 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 20
	}
	rows, QueryError := db.Query("SELECT * from results LIMIT $1 OFFSET $2", perPage, page)
	if QueryError != nil {
		log.Print("request error ", QueryError)
		return nil, QueryError
	}
	return rows, nil
}

func UpdateResult(db *sql.DB, newResult poller.Response, oldTargetId int) (poller.Response, error) {
	var result poller.Response
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if newResult.Destination == "" || oldTargetId <= 0 {
		error := errors.New("newResult = nil or oldTarget is wrong")
		return result, error
	}
	_, err := db.Exec("UPDATE results SET destination = $1, status = $2, duration = $3 WHERE id = $4", newResult.Destination, newResult.Status, newResult.Time, oldTargetId)
	return newResult, err
}

func DeleteResult(db *sql.DB, resultId int) (poller.Response, error) {
	var result poller.Response
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	if resultId <= 0 {
		error := errors.New("resultId is wrong ")
		return result, error
	}
	errorQuery := db.QueryRow("DELETE FROM results WHERE id = $1 RETURNING *", resultId).Scan(&result.Id, &result.Target_id, &result.Destination, &result.Status, &result.Time, &result.Created_at)
	if errorQuery != nil {
		return result, errorQuery
	}
	return result, nil
}
