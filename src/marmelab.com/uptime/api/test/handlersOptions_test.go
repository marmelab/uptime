package test

import (
	"net/http"
	"testing"
	"log"
	"../router"
)

func newServer() {
	config := map[string]string{"port": "8384"}
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+config["port"], router))
}

func TestBeforeTest(t *testing.T) {
	go newServer()
}

func TestOptionsHandlerShouldNotRaiseErrorAndSetStatusCode(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "http://localhost:8384/results", nil)
	log.Print("ici", err)
	client := &http.Client{}
	resp, error := client.Do(req)
	log.Print("la", error)
	if resp.StatusCode != http.StatusOK {
		t.Error("Error,  OptionsHandler should return 200", resp.StatusCode)
	}
}
