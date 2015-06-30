package test

import (
	"../../poller"
	"../repositories"
	"database/sql"
	_ "github.com/lib/pq"
	"reflect"
	"testing"
)

func addResult(expectedTarget poller.Response) (poller.Response, *sql.DB) {
	var result poller.Response
	db, _ := connectToDB()
	_ = db.QueryRow("INSERT INTO results (target_id, destination, status, duration) VALUES($1, $2, $3, $4) RETURNING id", expectedTarget.Target_id, expectedTarget.Destination, &expectedTarget.Status, &expectedTarget.Time).Scan(&result.Id)
	return result, db
}

func TestAddValidResultShouldNotTriggerError(t *testing.T) {
	targetAdded, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	newResultAdded, errorAddResult := repositories.AddResult(db, newResult)
	if errorAddResult != nil {
		t.Error("Error add a valid result should not trigger a error", errorAddResult)
	}
	row, _ := db.Query("SELECT * FROM results WHERE id = $1", newResultAdded.Id)
	defer row.Close()
	var actualResult poller.Response
	for row.Next() {
		error := row.Scan(&actualResult.Id, &actualResult.Target_id, &actualResult.Destination, &actualResult.Status, &actualResult.Time, &actualResult.Created_at)
		if error != nil {
			t.Error("Error addResult doesn't add the target in database", error)
		}
	}
	if !reflect.DeepEqual(actualResult.Destination, newResult.Destination) {
		t.Error("the row returned is different to the row inserted")
	}
	if !reflect.DeepEqual(actualResult.Target_id, newResult.Target_id) {
		t.Error("the row returned is different to the row inserted")
	}
	if !reflect.DeepEqual(actualResult.Time, newResult.Time) {
		t.Error("the row returned is different to the row inserted")
	}
	if !reflect.DeepEqual(actualResult.Status, newResult.Status) {
		t.Error("the row returned is different to the row inserted")
	}
	emptyDatabase(db)
}

func TestAddInvalidResultShouldTriggerError(t *testing.T) {
	_, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: 0, Destination: "", Status: "", Time: -1}
	_, errorAddResult := repositories.AddResult(db, newResult)
	if errorAddResult == nil {
		t.Error("Error add a invalid result should trigger a error", errorAddResult)
	}
	emptyDatabase(db)
}

func TestGetValidResultShouldNotTriggerError(t *testing.T) {
	targetAdded, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	result, db := addResult(newResult)
	actualResult, error := repositories.GetResult(db, result.Id)
	if error != nil {
		t.Error("Error get a valid result should not raise a error", error)
	}
	if !reflect.DeepEqual(newResult.Destination, actualResult.Destination) {
		t.Error("Error get a valid result should return the same result as inserted")
	}
	if !reflect.DeepEqual(newResult.Status, actualResult.Status) {
		t.Error("Error get a valid result should return the same result as inserted")
	}
	if !reflect.DeepEqual(newResult.Time, actualResult.Time) {
		t.Error("Error get a valid result should return the same result as inserted")
	}
	emptyDatabase(db)
}

func TestGetInvalidResultShouldTriggerError(t *testing.T) {
	newResult := poller.Response{Target_id: 1, Destination: "AddValidTarget", Status: "good", Time: 1111}
	_, db := addResult(newResult)
	_, error := repositories.GetResult(db, -1)
	if error == nil {
		t.Error("Error get a invalid result should raise a error", error)
	}
	emptyDatabase(db)
}

func TestGetResultsShouldNotTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	targetAdded2, _ := addTarget("testAdd2")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	newResult2 := poller.Response{Target_id: targetAdded2.Id, Destination: "AddValidTarget2", Status: "failed", Time: 1111}
	_, _ = addResult(newResult)
	_, _ = addResult(newResult2)
	actualResults, error := repositories.GetResults(db, 0, 0)
	if error != nil {
		t.Error("Error get targets should not raise a error", error)
	}
	rows, _ := db.Query("SELECT target_id, destination, status, duration FROM results")
	if reflect.DeepEqual(actualResults, rows) {
		t.Error("Error expectedTarget actualResult should be nil")
	}
	emptyDatabase(db)
}

func TestUpdateValideResultShouldNotTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	result, _ := addResult(newResult)
	newResult2 := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget2", Status: "good", Time: 1111}
	_, error := repositories.UpdateResult(db, newResult2, result.Id)
	if error != nil {
		t.Error("replace a valid result with a valid result should not raise a error")
	}
	row, _ := db.Query("SELECT destination FROM results WHERE id = $1", result.Id)
	var destination string
	for row.Next() {
		_ = row.Scan(&destination)
	}
	if !reflect.DeepEqual(newResult2.Destination, destination) {
		t.Error("Error expectedTarget and actualResult are different")
	}
	emptyDatabase(db)
}

func TestUpdateInvalideResultShouldTriggerError(t *testing.T) {
	newResult := poller.Response{Target_id: 1, Destination: "AddValidTarget", Status: "good", Time: 1111}
	_, db := addResult(newResult)
	newResult2 := poller.Response{Target_id: 1, Destination: "AddValidTarget2", Status: "good", Time: 1111}
	_, error := repositories.UpdateResult(db, newResult2, -1)
	if error == nil {
		t.Error("replace a invalid result with a valid resultResultd raise a error")
	}
}

func TestDeleteValidResultShouldNotTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	result, _ := addResult(newResult)
	returned2, error := repositories.DeleteResult(db, result.Id)
	if error != nil {
		t.Error("delete a valid resultResultd should not raise a error")
	}
	_, errorQuery := db.Query("SELECT destination FROM result WHERE id = $1", returned2.Id)
	if errorQuery == nil {
		t.Error("Select a deleted result should return a error")
	}
	emptyDatabase(db)
}

func TestDeleteInvalidResultShouldTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	_, _ = addResult(newResult)
	_, error := repositories.DeleteResult(db, -1)
	if error == nil {
		t.Error("delete a invalid result should raise a error", error)
	}
	emptyDatabase(db)
}
