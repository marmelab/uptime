package test

import (
	"../database"
	Target "../target"
	"database/sql"
	_ "github.com/lib/pq"
	"reflect"
	"testing"
)

func TestGetDbShouldNotTriggerError(t *testing.T) {
	db, err := database.GetDb()
	if err != nil {
		t.Error("getDb should not raise a error", err)
	}
	if db == nil {
		t.Error("getDB should return a valid db")
	}
}

func TestAddValidTargetShouldNotTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptimeTest sslmode=disable")
	errorAddTarget := database.AddTarget(db, "youtube.fr")
	if errorAddTarget != nil {
		t.Error("Error add a valid target should not trigger a error")
	}
	row, err := db.Query("SELECT destination FROM Destination WHERE destination = $1", "youtube.fr")
	defer row.Close()
	var destination string
	for row.Next() {
		error := row.Scan(&destination)
		if error != nil {
			t.Error("Error addTarget doesn't add the target in database", err)
		}
	}
	if !reflect.DeepEqual("youtube.fr", destination) {
		t.Error("the row returned is different to the row inserted")
	}
	Target.DeleteTarget(db, "youtube.fr")
}

func TestAddInvalidTargetShouldTriggerError(t *testing.T) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptimeTest sslmode=disable")
	errorAddTarget := database.AddTarget(db, "")
	if errorAddTarget == nil {
		t.Error("Error add a invalid target should trigger a error", errorAddTarget)
	}
	row, _ := db.Query("SELECT destination FROM Destination WHERE destination = $1", "")
	defer row.Close()
	var destination string
	for row.Next() {
		error := row.Scan(&destination)
		if error == nil {
			t.Error("Error scan a  invalid target should trigger a error", error)
		}
	}
}

func TestGetValidTargetShouldNotTriggerError(t *testing.T) {
	db := Target.AddTarget("youtube.com")
	actualTarget, error := database.GetTarget(db, "youtube.com")
	if error != nil {
		t.Error("Error get a valid target should not raise a error", error)
	}
	if !reflect.DeepEqual("youtube.com", actualTarget.Destination) {
		t.Error("Error get a valid target should return the same target as inserted")
	}
	Target.DeleteTarget(db, "youtube.com")
}

func TestGetInvalidTargetShouldTriggerError(t *testing.T) {
	db := Target.AddTarget("youtube.com")
	actualTarget, error := database.GetTarget(db, "")
	if error == nil {
		t.Error("Error get a inexistante target should raise a error", error)
	}
	if reflect.DeepEqual("youtube.com", actualTarget.Destination) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	Target.DeleteTarget(db, "youtube.com")
}

func TestGetTargetsShouldNotTriggerError(t *testing.T) {
	db := Target.AddTarget("youtube.com")
	_ = Target.AddTarget("facebook.com")
	actualTargets, error := database.GetTargets(db)
	if error != nil {
		t.Error("Error get targets should not raise a error", error)
	}
	rows, _ := db.Query("SELECT Destination FROM Destination")
	if reflect.DeepEqual(actualTargets, rows) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	Target.DeleteTarget(db, "youtube.com")
	Target.DeleteTarget(db, "facebook.com")
}

func TestUpdateValideTargetShouldNotTriggerError(t *testing.T) {
	db := Target.AddTarget("youtube.com")
	expectedTarget2 := "facebook.com"
	error := database.UpdateTarget(db, expectedTarget2, "youtube.com")
	if error != nil {
		t.Error("replace a valid target with a valid target should not raise a error")
	}
	row, _ := db.Query("SELECT destination FROM Destination WHERE destination = $1", expectedTarget2)
	var destination string
	for row.Next() {
		_ = row.Scan(&destination)
	}
	if !reflect.DeepEqual(expectedTarget2, destination) {
		t.Error("Error expectedTarget and actualTarget are different")
	}
	Target.DeleteTarget(db, "facebook.com	")
}

func TestUpdateInvalideTargetShouldTriggerError(t *testing.T) {
	db := Target.AddTarget("youtube.com")
	expectedTarget2 := "facebook.com"
	error := database.UpdateTarget(db, expectedTarget2, "")
	if error == nil {
		t.Error("replace a invalid target should raise a error")
	}
	row, _ := db.Query("SELECT destination FROM Destination WHERE destination = $1", expectedTarget2)
	if reflect.DeepEqual("youtube.com", row) {
		t.Error("Error expectedTarget actualTarget should be different")
	}
	Target.DeleteTarget(db, "youtube.com")
}

func TestDeleteValidTargetShouldNotTriggerError(t *testing.T) {
	db := Target.AddTarget("youtube.com")
	error := database.DeleteTarget(db, "youtube.com")
	if error != nil {
		t.Error("delete a valid target should not raise a error")
	}
	row, _ := db.Query("SELECT destination FROM Destination WHERE destination = $1", "youtube.com")
	var destination string
	for row.Next() {
		_ = row.Scan(&destination)
	}
	if reflect.DeepEqual("youtube", destination) {
		t.Error("Error expectedTarget and actualTarget are different")
	}
}

func TestDeleteInvalidTargetShouldNotTriggerError(t *testing.T) {
	db := Target.AddTarget("youtube.com")
	error := database.DeleteTarget(db, "")
	if error == nil {
		t.Error("delete a invalid target should raise a error")
	}
	_, err := db.Query("SELECT destination FROM Destination WHERE destination = $1", "youtube.com")
	if err != nil {
		t.Error("get a target should not raise a error")
	}
	Target.DeleteTarget(db, "youtube.com")
}
