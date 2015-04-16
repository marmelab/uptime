package main

import (

    "log"
    "net/http"
    "encoding/json"
    "fmt"


)

type Ip struct{
	Destination string   `json:"destination"`
}

type Ips []Ip


func main() {
    http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {

        listIp := Ips{
        	Ip{Destination: "google.fr"},
        	Ip{Destination: "youtube.fr"},
        	Ip{Destination: "bing.fr"},
        	Ip{Destination: "szszdzdadafdff.fr"},
        }
        json.NewEncoder(w).Encode(listIp)
    })
    log.Fatal(http.ListenAndServe(":8000", nil))
}
