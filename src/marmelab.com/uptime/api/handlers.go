package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)


func RetrieveTargets(w http.ResponseWriter, r *http.Request) {
	db := getDb()
	rows, err := getTargets(db)
	if err == nil {
		targets := make([]target.Target_data, 0)
		for rows.Next() {
			var id int
			var dest string
			var status bool
			error := rows.Scan(&id, &dest, &status)
			if error != nil {
				log.Print(error)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			targets = append(targets, target.Target_data{Id: id, Destination: dest, Status: status})
		}
		json.NewEncoder(w).Encode(targets)
	} else {
		log.Print(error)
		http.Error(w, http.StatusText(500), 500)
		return		
	}
}

func ShowTarget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := getDb()
	row, err := getTargets(db, vars)
	if err == nil {
		ips := make([]target.Target_data, 0)
		for rows.Next() {
			var id int
			var dest string
			var status bool
			error := rows.Scan(&id, &dest, &status)
			if error != nil {
				log.Print(error)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			ips = append(ips, target.Target_data{Id: id, Destination: dest, Status: status})
		}
		json.NewEncoder(w).Encode(ips)		
	} else {
		log.Print(error)
		http.Error(w, http.StatusText(500), 500)
		return		
	}
}


func CreateTarget(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newTarget string
	error := decoder.Decode(&newTarget)
	if error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	db := getDb()
	err := addTarget(db, newTarget)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return		
	}
}

func UpdateTarget(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newTarget string
	var oldTarget string
	error := decoder.Decode(&newTarget, &oldTarget)
	if error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	db := getDb()
	err := UpdateTarget(db, newTarget, oldTarget)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return		
	}
}

func DeleteTarget(w http.ResponseWriter, r *http.Request) {
			if r.Method == "DELETE" {
			decoder := json.NewDecoder(r.Body)
			var target string
			error := decoder.Decode(&target)
			if error != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}	
}
