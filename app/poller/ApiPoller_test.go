<<<<<<< HEAD
<<<<<<< HEAD
package poller

=======

package poller


>>>>>>> add PingPoller_test.go
=======
package poller

>>>>>>> test failed again
import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoPostOnWithoutUrl(t *testing.T) {
<<<<<<< HEAD
<<<<<<< HEAD
	response := Response{"localhost", "good", 50, nil}
=======
	response := Response{}
	response.Time = 50
	response.Destination = "localhost"
	response.Status = "good"
	response.Error = nil
>>>>>>> add PingPoller_test.go
=======
	response := Response{"localhost", "good", 50, nil}
>>>>>>>  code and test fixed
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
<<<<<<< HEAD
<<<<<<< HEAD
	response := Response{"localhost", "good", 50, nil}
=======
	response := Response{}
	response.Time = 50
	response.Destination = "localhost"
	response.Status = "good"
	response.Error = nil
>>>>>>> add PingPoller_test.go
=======
	response := Response{"localhost", "good", 50, nil}
>>>>>>>  code and test fixed
	err := DoPostOn(&response, ts.URL)
	if err != nil {
		t.Error("Expected no error, got", err)
	}
}
