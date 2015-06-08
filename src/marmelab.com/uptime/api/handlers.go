package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func RetrieveTargets(w http.ResponseWriter, r *http.Request) {
	db := initDb()
	rows := getTargets(db)
	ips := make([]target.Target_data, 0)
	for rows.Next() {
		var id int
		var dest string
		var status bool
		error := rows.Scan(&id, &dest, &status)
		log.Print(error)
		if error != nil {
			log.Print(error)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		ips = append(ips, target.Target_data{Id: id, Destination: dest, Status: status})
	}
	json.NewEncoder(w).Encode(ips)
}

func ShowTarget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
}


func CreateTarget(w http.ResponseWriter, r *http.Request) {

}

func UpdateTarget(w http.ResponseWriter, r *http.Request) {
	
}

func DeleteTarget(w http.ResponseWriter, r *http.Request) {
	
}
