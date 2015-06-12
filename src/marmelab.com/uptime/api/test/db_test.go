package test

import (
	"../repositories"
	"../target"
	"database/sql"
	_ "github.com/lib/pq"
	"reflect"
	"testing"
	"log"
	"../../poller"
)

func connectToDB() (*sql.DB, error){
	configdb := poller.RetrieveConfDbFromJsonFile("./conf_test.json")
	return sql.Open("postgres", "host="+configdb["host"].(string)+" user="+configdb["user"].(string)+" dbname="+configdb["dbname"].(string)+" sslmode="+configdb["sslmode"].(string)+"")
}
func addTarget(destination string) (target.Target_data, *sql.DB) {
	var result target.Target_data
	db, _ := connectToDB()
	expectedTarget := target.Target_data{Destination: destination}
	_ = db.QueryRow("INSERT INTO Destination (destination) VALUES($1) RETURNING id", expectedTarget.Destination).Scan(&result.Id)
	return result, db
}

func emptyDatabase(db *sql.DB) {
	_, _ = db.Exec("TRUNCATE Destination")
}

func TestGetDbShouldNotTriggerError(t *testing.T) {
	db, err := repositories.GetDb()
	if err != nil {
		t.Error("getDb should not raise a error", err)
	}
	if db == nil {
		t.Error("getDB should return a valid db")
	}
}

func TestAddValidTargetShouldNotTriggerError(t *testing.T) {
	db, _ := connectToDB()
	newTarget := target.Target_data{Destination: "youtube.fr"}
	newTargetAdded, errorAddTarget := repositories.AddTarget(db, newTarget)
	if errorAddTarget != nil {
		t.Error("Error add a valid target should not trigger a error", errorAddTarget)
	}
	row, err := db.Query("SELECT * FROM Destination WHERE id = $1", newTargetAdded.Id)
	defer row.Close()
	var actualTarget target.Target_data
	for row.Next() {
		error := row.Scan(&actualTarget.Id, &actualTarget.Destination)
		if error != nil {
			t.Error("Error addTarget doesn't add the target in database", err)
		}
	}
	if !reflect.DeepEqual(actualTarget.Destination, newTarget.Destination) {
		t.Error("the row returned is different to the row inserted")
	}
	emptyDatabase(db)
}

func TestAddInvalidTargetShouldTriggerError(t *testing.T) {
	db, _ := connectToDB()
	newTarget := target.Target_data{Destination: ""}
	_, errorAddTarget := repositories.AddTarget(db, newTarget)
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
	result, db := addTarget("youtube.com")
	actualTarget, error := repositories.GetTarget(db, result.Id)
	if error != nil {
		t.Error("Error get a valid target should not raise a error", error)
	}
	if !reflect.DeepEqual("youtube.com", actualTarget.Destination) {
		t.Error("Error get a valid target should return the same target as inserted")
	}
	emptyDatabase(db)
}

func TestGetInvalidTargetShouldTriggerError(t *testing.T) {
	_, db := addTarget("youtube.com")
	actualTarget, error := repositories.GetTarget(db, -1)
	if error == nil {
		t.Error("Error get a inexistant target should raise a error", error)
	}
	if reflect.DeepEqual("youtube.com", actualTarget.Destination) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	emptyDatabase(db)
}

func TestGetTargetsWithLastResultShouldNotTriggerError(t *testing.T) {
	_, db := addTarget("youtube.com")
	_, _ = addTarget("facebook.com")
	actualTargets, error := repositories.GetTargetsWithLastResult(db)
	if error != nil {
		t.Error("Error get targets should not raise a error", error)
	}
	rows, _ := db.Query("SELECT Destination FROM Destination")
	if reflect.DeepEqual(actualTargets, rows) {
		t.Error("Error expectedTarget actualTarget should be nil")
	}
	emptyDatabase(db)
	emptyDatabase(db)
}

func TestUpdateValideTargetShouldNotTriggerError(t *testing.T) {
	returned, db := addTarget("youtube.com")
	expectedTarget := target.Target_data{Destination: "facebook.com"}
	log.Print(returned.Id)
	error := repositories.UpdateTarget(db, expectedTarget,returned.Id)
	if error != nil {
		t.Error("replace a valid target with a valid target should not raise a error")
	}
	row, _ := db.Query("SELECT destination FROM Destination WHERE destination = $1", expectedTarget.Destination)
	var destination string
	for row.Next() {
		_ = row.Scan(&destination)
	}
	if !reflect.DeepEqual(expectedTarget.Destination, destination) {
		t.Error("Error expectedTarget and actualTarget are different")
	}
	emptyDatabase(db)
}

func TestUpdateInvalideTargetShouldTriggerError(t *testing.T) {
	_, db := addTarget("youtube.com")
	expectedTarget := target.Target_data{Destination: "facebook.com"}
	error := repositories.UpdateTarget(db, expectedTarget, -4)
	if error == nil {
		t.Error("replace a invalid target should raise a error")
	}
	row, _ := db.Query("SELECT destination FROM Destination WHERE destination = $1", expectedTarget.Destination)
	if reflect.DeepEqual("youtube.com", row) {
		t.Error("Error expectedTarget actualTarget should be different")
	}
	emptyDatabase(db)
}

func TestDeleteValidTargetShouldNotTriggerError(t *testing.T) {
	returned, db := addTarget("youtube.com")
	returned2, error := repositories.DeleteTarget(db, returned.Id)
	expectedTarget := target.Target_data{Id: returned2.Id, Destination: "youtube.com"}
	if error != nil {
		t.Error("delete a valid target should not raise a error")
	}
	row, _ := db.Query("SELECT * FROM Destination WHERE id = $1", returned2.Id)
	var actualTarget target.Target_data
	for row.Next() {
		_ = row.Scan(&actualTarget.Id, & actualTarget.Destination)
	}
	if reflect.DeepEqual(expectedTarget, actualTarget) {
		t.Error("Error expectedTarget and actualTarget are different")
	}
}

func TestDeleteInvalidTargetShouldNotTriggerError(t *testing.T) {
	_, db := addTarget("youtube.com")
	_, error := repositories.DeleteTarget(db, -1)
	if error == nil {
		t.Error("delete a invalid target should raise a error")
	}
	_, err := db.Query("SELECT destination FROM Destination WHERE id = $1", 1)
	if err != nil {
		t.Error("get a target should not raise a error")
	}
	emptyDatabase(db)
}
