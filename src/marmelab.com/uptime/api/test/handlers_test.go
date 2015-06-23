package test

import (
	Router "../router"
	"../target"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func newServer() {
	config := map[string]string{"port": "8384"}
	router := Router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+config["port"], router))
}

func TestBeforeTest(t *testing.T) {
	go newServer()
}
func TestGetTargetsShouldNotTriggerError(t *testing.T) {
	addTarget("google.fr")
	addTarget("youtube.fr")
	var listOfTarget []target.Target_data
	response, err := http.Get("http://localhost:8384/targets")
	if response.StatusCode != http.StatusOK {
		t.Error("Error, GetTargets should not return a error", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Error("Error, GetTargets should not return a error", err)
	}
	err = json.Unmarshal(body, &listOfTarget)
	if err != nil {
		t.Error("Error, GetTargets should not return a error", err)
	}
	if listOfTarget[0].Destination != "google.fr" {
		t.Error("Error, GetTargets should not return a error", err)
	}
	db, _ := connectToDB()
	emptyDatabase(db)
}

func TestGetTargetWithValidIdShouldNotTriggerError(t *testing.T) {
	res, db := addTarget("google.fr")
	response, err := http.Get("http://localhost:8384/targets/" + strconv.Itoa(res.Id)) // TODO
	if response.StatusCode != http.StatusOK {
		t.Error("Error, ShowTarget should not return a error", err)
	}
	var getTarget target.Target_data
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	err = json.Unmarshal(body, &getTarget)
	log.Print(getTarget)
	if getTarget.Destination != "google.fr" {
		t.Error("Error, GetTarget should not return a error", err)
	}
	emptyDatabase(db)
}

func TestGetTargetWithBadIdsShouldTriggerError(t *testing.T) {
	ids := [3]string{"-1", "0", "14235581485"}
	for i := 0; i < 3; i++ {
		response, err := http.Get("http://localhost:8384/targets/" + ids[i])
		if response.StatusCode == http.StatusOK {
			t.Error("Error, ShowTarget should return a error", err)
		}
	}
}

func TestCreateTargetWithValidDataShouldNotTriggerError(t *testing.T) {
	newTarget := target.Target_data{Destination: "blablabla", Status: false}
	data, _ := json.Marshal(newTarget)
	req, error := http.NewRequest("POST", "http://localhost:8384/targets", bytes.NewBuffer(data))
	client := &http.Client{}
	resp, err := client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	var actualTarget target.Target_data
	errDecode := decoder.Decode(&actualTarget)
	if errDecode != nil {
		t.Error("Error : ", errDecode)
	}
	if actualTarget.Destination != newTarget.Destination {
		t.Error("Error, target updated is different from exepected ")
	}
	if resp.StatusCode != 200 {
		t.Error("Error, UpdateTarget should not return a error", resp.StatusCode)
	}
	if err != nil {
		t.Error("Error, UpdateTarget should not return a error", err)
	}
	if error != nil {
		t.Error("Error, CreateTarget should not return a error", error)
	}
	defer resp.Body.Close()
	db, _ := connectToDB()
	emptyDatabase(db)
}

func TestUpdateTargetWithValidIdShouldTriggerError(t *testing.T) {
	res, db := addTarget("google.fr")
	targetForUpdate := target.Target_data{Destination: "testUpdateTarget"}
	data, _ := json.Marshal(targetForUpdate)
	req, error := http.NewRequest("PUT", "http://localhost:8384/targets/"+strconv.Itoa(res.Id), bytes.NewBuffer(data))
	client := &http.Client{}
	resp, err := client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	var newTarget target.Target_data
	errDecode := decoder.Decode(&newTarget)
	if errDecode != nil {
		t.Error("Error : ", errDecode)
	}
	if newTarget.Destination != targetForUpdate.Destination {
		t.Error("Error, target updated is different from exepected ")
	}
	if resp.StatusCode != 200 {
		t.Error("Error, UpdateTarget should not return a error", resp.StatusCode)
	}
	if err != nil {
		t.Error("Error, UpdateTarget should not return a error", err)
	}
	if error != nil {
		t.Error("Error, UpdateTarget should not return a error", error)
	}
	emptyDatabase(db)
}

func TestDeleteTargetWithValidIdShouldTriggerError(t *testing.T) {
	res, _ := addTarget("google.fr")
	req, error := http.NewRequest("DELETE", "http://localhost:8384/targets/"+strconv.Itoa(res.Id), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	var newTarget target.Target_data
	errDecode := decoder.Decode(&newTarget)
	if errDecode != nil {
		t.Error("Error : ", errDecode)
	}
	if newTarget.Destination != "google.fr" {
		t.Error("Error, target updated is different from exepected ")
	}
	if resp.StatusCode != 200 {
		t.Error("Error, UpdateTarget should not return a error", resp.StatusCode)
	}
	if error != nil {
		t.Error("Error, DeleteTarget should not return a error", err)
	}
}
