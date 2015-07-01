package test

import (
	"net/http"
	"testing"
	"encoding/json"
	"../target"
	"io/ioutil"
)

func TestPaginationWithValidValuesShouldNotTriggerError(t *testing.T) {
	_, db := addTarget("google.fr")
	addTarget("youtube2.fr")
	addTarget("youtube3.fr")
	addTarget("youtube4.fr")
	addTarget("youtube5.fr")
	addTarget("youtube6.fr")
	addTarget("youtube7.fr")

	var listOfTarget []target.Target_data
	response, err := http.Get("http://localhost:8384/targets?page=0&perPage=5")
	if response.StatusCode != http.StatusOK {
		t.Error("Error, GetTargets should return 200 as code", err)
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
	if len(listOfTarget) != 5 {
		t.Error("Error GetTargets pagination should return the good numbers of targets, got :", len(listOfTarget))
	}
	emptyDatabase(db)
}

func TestPaginationWithValidValuesButDifferentPerPageShouldNotTriggerError(t *testing.T) {
	_, db := addTarget("google.fr")
	addTarget("youtube2.fr")
	addTarget("youtube3.fr")
	addTarget("youtube4.fr")
	addTarget("youtube5.fr")
	addTarget("youtube6.fr")
	addTarget("youtube7.fr")
	addTarget("youtube8.fr")
	addTarget("youtube9.fr")
	addTarget("youtube11.fr")
	addTarget("youtube12.fr")
	addTarget("youtube13.fr")
	addTarget("youtube14.fr")

	var listOfTarget []target.Target_data
	response, err := http.Get("http://localhost:8384/targets?page=0&perPage=10")
	if response.StatusCode != http.StatusOK {
		t.Error("Error, GetTargets should return 200 as code", response.StatusCode)
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
	if len(listOfTarget) != 10 {
		t.Error("Error GetTargets pagination should return the good numbers of targets, got :", len(listOfTarget))
	}
	emptyDatabase(db)
}

func TestPaginationWithValidValuesButWithDifferentPageShouldNotTriggerError(t *testing.T) {
	_, db := addTarget("google.fr")
	addTarget("youtube2.fr")
	addTarget("youtube3.fr")
	addTarget("youtube4.fr")
	addTarget("youtube5.fr")
	addTarget("youtube6.fr")
	addTarget("youtube7.fr")
	addTarget("youtube8.fr")
	addTarget("youtube9.fr")
	addTarget("youtube11.fr")
	addTarget("youtube12.fr")
	addTarget("youtube13.fr")
	addTarget("youtube14.fr")

	var listOfTarget []target.Target_data
	response, err := http.Get("http://localhost:8384/targets?page=4&perPage=10")
	if response.StatusCode != http.StatusOK {
		t.Error("Error, GetTargets should return 200 as code", response.StatusCode)
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
	if len(listOfTarget) != 3 {
		t.Error("Error GetTargets pagination should return the good numbers of targets, got :", len(listOfTarget))
	}
	emptyDatabase(db)
}

func TestPaginationWithInValidValuesShouldNotTriggerError(t *testing.T) {
	_, db := addTarget("google.fr")
	addTarget("youtube2.fr")
	addTarget("youtube3.fr")
	addTarget("youtube4.fr")

	var listOfTarget []target.Target_data
	response, err := http.Get("http://localhost:8384/targets?page=-3&perPage=-4")
	if response.StatusCode != http.StatusOK {
		t.Error("Error, GetTargets should return 200 as code", response.StatusCode)
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
	emptyDatabase(db)
}
