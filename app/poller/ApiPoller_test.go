<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>> test failed again
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
package poller

=======

package poller


>>>>>>> add PingPoller_test.go
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> test failed again
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
package poller

>>>>>>> test failed again
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>> test failed again
=======
=======
package poller

>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoPostOnWithoutUrl(t *testing.T) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>>  code and test fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	response := Response{"localhost", "good", 50, nil}
=======
	response := Response{}
	response.Time = 50
	response.Destination = "localhost"
	response.Status = "good"
	response.Error = nil
>>>>>>> add PingPoller_test.go
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
	response := Response{"localhost", "good", 50, nil}
>>>>>>>  code and test fixed
=======
<<<<<<< HEAD
>>>>>>> add PingPoller_test.go
=======
=======
	response := Response{"localhost", "good", 50, nil}
>>>>>>>  code and test fixed
>>>>>>>  code and test fixed
=======
	response := Response{"localhost", "good", 50, nil}
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
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
<<<<<<< HEAD
<<<<<<< HEAD
=======
	response := Response{"localhost", "good", 50, nil}
=======
	response := Response{}
	response.Time = 50
	response.Destination = "localhost"
	response.Status = "good"
	response.Error = nil
>>>>>>> add PingPoller_test.go
>>>>>>> add PingPoller_test.go
=======
>>>>>>>  code and test fixed
=======
=======
	response := Response{"localhost", "good", 50, nil}
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	err := DoPostOn(&response, ts.URL)
	if err != nil {
		t.Error("Expected no error, got", err)
	}
}
