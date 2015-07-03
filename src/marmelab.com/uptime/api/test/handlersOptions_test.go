package test

import (
	"testing"
	"log"
	"net/http"
)



func TestOptionsHandlerShouldNotRaiseErrorAndSetStatusCode(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "http://localhost:8384/results", nil)
	log.Print("request error", err)
	client := &http.Client{}
	resp, error := client.Do(req)
	log.Print("reponse error", error)
	if resp.StatusCode != http.StatusOK {
		t.Error("Error,  OptionsHandler should return 200", resp.StatusCode)
	}
}
