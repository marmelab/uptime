<<<<<<< HEAD
<<<<<<< HEAD
package main

import (
=======



package main

import (

>>>>>>> add PingPoller_test.go
=======
package main

import (
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
			var ips [2]model.Ip
			ips[0].Destination = "google.fr"
			ips[1].Destination = "failfailfail.fail"
			listIp := ips
=======
			listIp := model.Ips{
				model.Ip{Destination: "google.fr"},
				model.Ip{Destination: "youtube.fr"},
				model.Ip{Destination: "bing.fr"},
				model.Ip{Destination: "szszdzdadafdff.fr"},
			}
>>>>>>> add PingPoller_test.go
			json.NewEncoder(w).Encode(listIp)
		}

	})
	http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
=======
		if r.Method == "GET" {
		}
>>>>>>> add PingPoller_test.go
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
=======

>>>>>>> add PingPoller_test.go
=======
>>>>>>> test failed again
