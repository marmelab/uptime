package handlers

import (
	"../repositories"
	"../target"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func SetCors(w *http.Header) {
	w.Set("Access-Control-Allow-Origin", "*")
	w.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		log.Print("ERROR GetDb", errorGetDb)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rows, errGetResults := repositories.GetResults(db)
	if errGetResults != nil {
		log.Print("ERROR GetResults", errGetResults)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	targets := make([]target.Target_data, 0)
	for rows.Next() {
		var newTarget target.Target_data
		error := rows.Scan(&newTarget.Id, &newTarget.Destination, &newTarget.Status)
		if error != nil {
			log.Print("ERROR Scan GetResults ", error)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		targets = append(targets, newTarget)
	}
	json.NewEncoder(w).Encode(targets)
}

func GetResult(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		log.Print("ERROR GetDb", errorGetDb)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	newTarget, errorGetTarget := repositories.GetResult(db, id)
	if errorGetTarget != nil {
		log.Print("ERROR GetResult", errorGetTarget)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(newTarget)
}

func PostResult(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	decoder := json.NewDecoder(r.Body)
	var newTarget target.Target_data
	error := decoder.Decode(&newTarget)
	if error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		log.Print("ERROR GetDb", errorGetDb)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	_, errAddTarget := repositories.AddTarget(db, newTarget)
	if errAddTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func PutResult(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	decoder := json.NewDecoder(r.Body)
	var newTarget target.Target_data
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	error := decoder.Decode(&newTarget.Destination)
	if error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		log.Print("ERROR GetDb", errorGetDb)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	errorUpdateTarget := repositories.UpdateTarget(db, newTarget, id)
	if errorUpdateTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func DeleteResult(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		log.Print("ERROR GetDb", errorGetDb)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	_, errorDeleteTarget := repositories.DeleteResult(db, id)
	if errorDeleteTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
