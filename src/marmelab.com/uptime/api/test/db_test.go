package test

import (
	"../Database.go"
	"reflect"
	"testing"
	"database/sql"
	_ "github.com/lib/pq"
	"../target"
)

func TestGetDbShouldNotTriggerError(t *testing.T) {
	db, err := Database.getDb()
	if err != nil {
		t.Error("getDb should not raise a error")
	}
	if db ==nil {
		t.Error("getDB should return a valid db")
	}
}

func TestAddValidTargetShouldNotTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	errorAddTarget := Database.AddTarget(db, expectedTarget.Destination)
	if(errorAddTarget != nil) {
		t.Error("Error add a valid target should not trigger a error")
	}
	row, error := db.Query("SELECT destination FROM testDestination WHERE destination = $1", expectedTarget.Destination)
	defer row.Close()
		target := make([]target.Target_data, 0)
		for row.Next() {
			var dest string
			error := rows.Scan(&dest)
			if error != nil {
				t.Error("Error addTarget doesn't add the target in database")
			}
	actualTarget := target.Target_data{Destination=dest}
	if !reflect.DeepEqual(expectedTarget, actualTarget) {
		t.Error("Error addTarget doesn't add the target in database")
	}
	_, _ = db.Exec("DELETE FROM testDestination WHERE destination = $1",expectedTarget.Destination)
}

func TestAddInvalidTargetShouldTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination=nil}
	errorAddTarget := Database.AddTarget(db, expectedTarget.Destination)
	if(errorAddTarget == nil) {
		t.Error("Error add a invalid target should trigger a error")
	}	
	row, error := db.Query("SELECT destination FROM testDestination WHERE destination = $1", expectedTarget.Destination)
	defer row.Close()
		target := make([]target.Target_data, 0)
		for row.Next() {
			var dest string
			error := rows.Scan(&dest)
			if error == nil {
				t.Error("Error add a invalid target should trigger a error")
			}
	actualTarget := target.Target_data{Destination=dest}
	if reflect.DeepEqual(expectedTarget, actualTarget) {
		t.Error("Error add a invalid target should trigger a error")
	}
}

func TestGetValidTargetShouldNotTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	_,_ = db.Exec("INSERT testDestination (destination) VALUES($1)", expectedTarget.Destination)
	actualTarget, error := Database.getTarget(db, 2)
	if error != nil {
		t.Error("Error get a valid target should not raise a error", error)
	}
	if !reflect.DeepEqual(expectedTarget,actualTarget) {
		t.Error("Error get a valid target should return the same target as inserted")
	}
	_, _ = db.Exec("DELETE FROM testDestination WHERE destination = $1",expectedTarget.Destination)	
}

func TestGetInvalidTargetShouldTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	_,_ = db.Exec("INSERT testDestination (destination) VALUES($1)", expectedTarget.Destination)
	actualTarget, error := Database.getTarget(db, 4)
	if error == nil {
		t.Error("Error get a inexistante target should raise a error", error)
	}
	if reflect.DeepEqual(expectedTarget,actualTarget) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	_, _ = db.Exec("DELETE FROM testDestination WHERE destination = $1",expectedTarget.Destination)
}

func TestGetTargetsShouldNotTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	 var expectedTarget2 target.Target_data{Destination="facebook.com"}
	_,_ = db.Exec("INSERT testDestination (destination) VALUES ($1), ($2)", expectedTarget.Destination, expectedTarget2.Destination)
	actualTargets, error := Database.getTargets(db)
	if error != nil {
		t.Error("Error get targets should not raise a error", error)
	}
	rows, _ := db.Query("SELECT Destination FROM testDestination")
	if reflect.DeepEqual(actualTargets,rows) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	_, _ = db.Exec("DELETE FROM testDestination WHERE destination = $1",expectedTarget.Destination)
	_, _ = db.Exec("DELETE FROM testDestination WHERE destination = $1",expectedTarget2.Destination)
}

func TestUpdateValideTargetShouldNotTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	_, _ = db.Exec("INSERT testDestination (destination) VALUES ($1)", expectedTarget.Destination)
	var expectedTarget2 target.Target_data{Destination="facebook.com"}
	error := UpdateTarget(db,expectedTarget2.Destination, expectedTarget.Destination)
	if error != nil {
		t.Error("replace a valid target with a valid target should not raise a error")
	}
	row, _ := db.Query("SELECT destination FROM testDestination WHERE destination = $1", expectedTarget2.Destination)
	if !reflect.DeepEqual(actualTargets,rows) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	_, _ = db.Exec("DELETE FROM testDestination WHERE destination = $1",expectedTarget.Destination)
}

func TestUpdateInvalideTargetShouldTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	_, _ = db.Exec("INSERT testDestination (destination) VALUES ($1)", expectedTarget.Destination)
	var expectedTarget2 target.Target_data{Destination="facebook.com"}
	error := UpdateTarget(db,expectedTarget2.Destination, nil)
	if error == nil {
		t.Error("replace a invalid target should raise a error")
	}
	row, _ := db.Query("SELECT destination FROM testDestination WHERE destination = $1", expectedTarget2.Destination)
	if reflect.DeepEqual(actualTargets,rows) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	_, _ = db.Exec("DELETE FROM testDestination WHERE destination = $1",expectedTarget.Destination)
}

func TestDeleteValidTargetShouldNotTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	_, _ = db.Exec("INSERT testDestination (destination) VALUES ($1)", expectedTarget.Destination)
	error := DeleteTarget(db, expectedTarget.Destination)
	if error != nil {
		t.Error("delete a valid target should not raise a error")
	}
	_, err := db.Query("SELECT destination FROM testDestination WHERE destination = $1", expectedTarget.Destination)
	if err == nil {
		t.Error("get a deleted target should raise a error")
	}
}

func TestDeleteInvalidTargetShouldNotTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination="youtube.fr"}
	_, _ = db.Exec("INSERT testDestination (destination) VALUES ($1)", expectedTarget.Destination)
	error := DeleteTarget(db, nil)
	if error == nil {
		t.Error("delete a invalid target should raise a error")
	}
	_, err := db.Query("SELECT destination FROM testDestination WHERE destination = $1", expectedTarget.Destination)
	if err != nil {
		t.Error("get a target should not raise a error")
	}
}
