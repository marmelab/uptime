
package poller


import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoPostOnWithoutUrl(t *testing.T) {
	response := Response{}
	response.Time = 50
	response.Destination = "localhost"
	response.Status = "good"
	response.Error = nil
	err := DoPostOn(&response, "")
	if err == nil {
		t.Error("Expected error, got", err)
	}
}

func TestDoPostOnWithUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			newResultat := Response{}
			error := decoder.Decode(&newResultat)
			if error != nil {
				t.Error("Expected no error, got ", error)
			}
		}
	}))
	defer ts.Close()
	response := Response{}
	response.Time = 50
	response.Destination = "localhost"
	response.Status = "good"
	response.Error = nil
	err := DoPostOn(&response, ts.URL)
	if err != nil {
		t.Error("Expected no error, got", err)
	}
}
