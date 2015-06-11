package target

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Target_data struct {
	Id          int    `json:"id"`
	Destination string `json:"destination"`
	Status      bool   `json:"status"`
}

func AddTarget(destination string) *sql.DB {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptimeTest sslmode=disable")
	expectedTarget := Target_data{Destination: destination}
	_, _ = db.Exec("INSERT INTO Destination (destination) VALUES($1)", expectedTarget.Destination)
	return db
}

func DeleteTarget(db *sql.DB, destination string) {
	_, _ = db.Exec("DELETE FROM Destination WHERE destination = $1", destination)
}
