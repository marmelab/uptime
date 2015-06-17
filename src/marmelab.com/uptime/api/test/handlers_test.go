 package test

// import (
// 	Router "../router"
// 	"bytes"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"testing"
// 	"runtime"
// 	"path/filepath"
// 	"../../poller"
// 	"io/ioutil"
// )

// func newServer() {
// 	_, filename, _, _ := runtime.Caller(1)
// 	_ = filepath.Join(filepath.Dir(filename), "conf.json")
// 	log.Print(">>>>>>>>>>>>>>>1", filepath.Dir(filename))
// 	config := poller.RetrieveConfDbFromJsonFile("./conf_test.json")
// 	router := Router.NewRouter()
// 	log.Fatal(http.ListenAndServe(":"+config["port"].(string), router))
// }

// func TestBeforeTest(t *testing.T) {
// 	go newServer()
// }
// func TestRetrieveTargetsShouldNotTriggerError(t *testing.T) {
// 	_, db := addTarget("testaddtarget1")
// 	_, _ = addTarget("testaddtarget2")
// 	var listOfTarget []string
// 	response, err := http.Get("http://localhost:8384/targets")
// 	if response.StatusCode != http.StatusOK {
// 		t.Error("Error, RetrieveTargets should not return a error", err)
// 	}
// 	body, err := ioutil.ReadAll(response.Body)
// 	response.Body.Close()
// 	if err != nil {
// 		t.Error("Error, RetrieveTargets should not return a error", err)
// 	}
// 	err = json.Unmarshal(body, &listOfTarget)
// 	if err != nil {
// 		t.Error("Error, RetrieveTargets should not return a error", err)
// 	}
// 	if(listOfTarget[0] != "testaddtarget1" || listOfTarget[1] != "testaddtarget2") {
// 		t.Error("Error, RetrieveTargets should not return a error", err)
// 	}
// 	emptyDatabase(db)
// }

// func TestShowTargetWithValidIdShouldNotTriggerError(t *testing.T) {
// 	response, err := http.Get("http://localhost:8384/targets/1")
// 	if response.StatusCode != http.StatusOK {
// 		t.Error("Error, ShowTarget should not return a error", err)
// 	}
// }

// func TestShowTargetWithBadIdsShouldTriggerError(t *testing.T) {
// 	ids := [3]string{"-1", "0", "14235581485"}
// 	for i := 0; i < 3; i++ {
// 		response, err := http.Get("http://localhost:8384/targets/" + ids[i])
// 		if response.StatusCode == http.StatusOK {
// 			t.Error("Error, ShowTarget should return a error", err)
// 		}
// 	}
// }

// func TestCreateTargetWithValidDataShouldNotTriggerError(t *testing.T) {
// 	data, _ := json.Marshal("testCreateTarget")
// 	_, error := http.NewRequest("POST", "http://localhost:8384/targets", bytes.NewBuffer(data))
// 	if error != nil {
// 		t.Error("Error, CreateTarget should not return a error", error)
// 	}
// }

// func TestUpdateTargetWithNullIdShouldTriggerError(t *testing.T) {
// 	data, _ := json.Marshal("testUpdateTarget")
// 	_, error := http.NewRequest("PUT", "http://localhost:8384/targets/1", bytes.NewBuffer(data))
// 	if error != nil {
// 		t.Error("Error, UpdateTarget should not return a error", error)
// 	}
// }

// func TestDeleteTargetWithValidIdShouldTriggerError(t *testing.T) {
// 	_, error := http.NewRequest("DELETE", "http://localhost:8384/targets/1", nil)
// 	if error != nil {
// 		t.Error("Error, DeleteTarget should not return a error", error)
// 	}
// }
