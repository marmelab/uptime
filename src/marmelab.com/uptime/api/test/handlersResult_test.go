package test

import (
	"../../poller"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestGetResultsHandlerShouldNotTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	targetAdded2, _ := addTarget("testAdd2")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	newResult2 := poller.Response{Target_id: targetAdded2.Id, Destination: "AddValidTarge2", Status: "good", Time: 1111}
	addResult(newResult)
	addResult(newResult2)
	var listOfResult []poller.Response
	response, err := http.Get("http://localhost:8384/results")
	if response.StatusCode != http.StatusOK {
		t.Error("Error, GetResults should not return a error", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Error("Error, GetResults should not return a error", err)
	}
	err = json.Unmarshal(body, &listOfResult)
	if err != nil {
		t.Error("Error, GetResults should not return a error", err)
	}
	if listOfResult[0].Destination != "AddValidTarget" {
		t.Error("Error, GetResults should not return a error", err)
	}
	if listOfResult[1].Destination != "AddValidTarge2" {
		t.Error("Error, GetResults should not return a error", err)
	}
	emptyDatabase(db)
}

func TestGetResultWithValidIdShouldNotTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarge2", Status: "good", Time: 1111}
	res, _ := addResult(newResult)
	response, err := http.Get("http://localhost:8384/results/" + strconv.Itoa(res.Id))
	if response.StatusCode != http.StatusOK {
		t.Error("Error, ShowResult should not return a error", err)
	}
	var getResult poller.Response
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	err = json.Unmarshal(body, &getResult)
	if getResult.Destination != "AddValidTarge2" {
		t.Error("Error, GetResult should not return a error", err)
	}
	emptyDatabase(db)
}

func TestGetResultWithBadIdsShouldTriggerError(t *testing.T) {
	ids := [3]string{"-1", "0", "14235581485"}
	for i := 0; i < 3; i++ {
		response, err := http.Get("http://localhost:8384/results/" + ids[i])
		if response.StatusCode == http.StatusOK {
			t.Error("Error, ShowResult should return a error", err)
		}
	}
}

func TestCreateResultWithValidDataShouldNotTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarge2", Status: "good", Time: 1111}
	data, _ := json.Marshal(newResult)
	req, error := http.NewRequest("POST", "http://localhost:8384/results", bytes.NewBuffer(data))
	client := &http.Client{}
	resp, err := client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	var actualResult poller.Response
	errDecode := decoder.Decode(&actualResult)
	if errDecode != nil {
		t.Error("Error : ", errDecode)
	}
	if actualResult.Destination != newResult.Destination {
		t.Error("Error, result updated is different from expected ")
	}
	if resp.StatusCode != 200 {
		t.Error("Error, CreateResult should not return a error", resp.StatusCode)
	}
	if err != nil {
		t.Error("Error, CreateResult should not return a error", err)
	}
	if error != nil {
		t.Error("Error, CreateResult should not return a error", error)
	}
	defer resp.Body.Close()
	emptyDatabase(db)
}

func TestUpdateResultWithValidIdShouldTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	res, _ := addResult(newResult)
	newResult2 := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarge2", Status: "good", Time: 1111}
	data, _ := json.Marshal(newResult2)
	req, error := http.NewRequest("PUT", "http://localhost:8384/results/"+strconv.Itoa(res.Id), bytes.NewBuffer(data))
	client := &http.Client{}
	resp, err := client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	var result poller.Response
	errDecode := decoder.Decode(&result)
	if errDecode != nil {
		t.Error("Error : ", errDecode)
	}
	if result.Destination != newResult2.Destination {
		t.Error("Error, result updated is different from exepected ")
	}
	if resp.StatusCode != 200 {
		t.Error("Error, UpdateResult should not return a error", resp.StatusCode)
	}
	if err != nil {
		t.Error("Error, UpdateResult should not return a error", err)
	}
	if error != nil {
		t.Error("Error, UpdateResult should not return a error", error)
	}
	emptyDatabase(db)
}

func TestDeleteResultWithValidIdShouldTriggerError(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarget", Status: "good", Time: 1111}
	res, _ := addResult(newResult)
	req, error := http.NewRequest("DELETE", "http://localhost:8384/results/"+strconv.Itoa(res.Id), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	var result poller.Response
	errDecode := decoder.Decode(&result)
	if errDecode != nil {
		t.Error("Error : ", errDecode)
	}
	if result.Destination != "AddValidTarget" {
		t.Error("Error, target updated is different from exepected ")
	}
	if resp.StatusCode != 200 {
		t.Error("Error, UpdateResult should not return a error", resp.StatusCode)
	}
	if error != nil {
		t.Error("Error, DeleteResult should not return a error", err)
	}
	emptyDatabase(db)
}

func TestSetContentTypeWorksForGetRequest(t *testing.T) {
	targetAdded1, db := addTarget("testAdd")
	newResult := poller.Response{Target_id: targetAdded1.Id, Destination: "AddValidTarge2", Status: "good", Time: 1111}
	res, _ := addResult(newResult)
	response, err := http.Get("http://localhost:8384/results/" + strconv.Itoa(res.Id))
	if response.StatusCode != http.StatusOK {
		t.Error("Error, ShowResult should not return a error", err)
	}
	if response.Header.Get("Content-Type") != "application/json" {
		t.Error("SetContentType didn't set Content-Type of the response")
	}
	emptyDatabase(db)
}
