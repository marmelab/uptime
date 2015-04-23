<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>> test failed again
package main

import (
=======



package main

import (

>>>>>>> add PingPoller_test.go
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> test failed again
=======
package main

import (
>>>>>>> test failed again
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>> test failed again
	"../poller"
	"./model"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
>>>>>>> add PingPoller_test.go
=======
=======
>>>>>>>  code and test fixed
>>>>>>>  code and test fixed
			var ips [2]model.Ip
			ips[0].Destination = "google.fr"
			ips[1].Destination = "failfailfail.fail"
			listIp := ips
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>>  code and test fixed
=======
			listIp := model.Ips{
				model.Ip{Destination: "google.fr"},
				model.Ip{Destination: "youtube.fr"},
				model.Ip{Destination: "bing.fr"},
				model.Ip{Destination: "szszdzdadafdff.fr"},
			}
>>>>>>> add PingPoller_test.go
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
>>>>>>> add PingPoller_test.go
=======
=======
>>>>>>>  code and test fixed
>>>>>>>  code and test fixed
			json.NewEncoder(w).Encode(listIp)
		}

	})
	http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> add PingPoller_test.go
=======
>>>>>>>  code and test fixed
=======
		if r.Method == "GET" {
		}
>>>>>>> add PingPoller_test.go
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
>>>>>>> add PingPoller_test.go
=======
=======
>>>>>>>  code and test fixed
>>>>>>>  code and test fixed
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			newResultat := poller.Response{}
			error := decoder.Decode(&newResultat)
			if error != nil {
				log.Fatal(error)
			}
			log.Print(newResultat)
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======

>>>>>>> add PingPoller_test.go
=======
>>>>>>> test failed again
=======
=======

>>>>>>> add PingPoller_test.go
>>>>>>> add PingPoller_test.go
=======
=======

>>>>>>> add PingPoller_test.go
=======
>>>>>>> test failed again
>>>>>>> test failed again
