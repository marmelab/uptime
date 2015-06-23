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

func GetTargets(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		log.Print("ERROR GetDb", errorGetDb)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rows, errGetTargets := repositories.GetTargetsWithLastResult(db)
	if errGetTargets != nil {
		log.Print("ERROR GetTargets", errGetTargets)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	targets := make([]target.Target_data, 0)
	for rows.Next() {
		var newTarget target.Target_data
		error := rows.Scan(&newTarget.Id, &newTarget.Destination, &newTarget.Status)
		if error != nil {
			log.Print("ERROR Scan GetTargets ", error)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		targets = append(targets, newTarget)
	}
	json.NewEncoder(w).Encode(targets)
}

func GetTarget(w http.ResponseWriter, r *http.Request) {
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
	newTarget, errorGetTarget := repositories.GetTarget(db, id)
	if errorGetTarget != nil {
		log.Print("ERROR GetTarget", errorGetTarget)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(newTarget)
}

func PostTarget(w http.ResponseWriter, r *http.Request) {
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
	targetAdded, errAddTarget := repositories.AddTarget(db, newTarget)
	if errAddTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(targetAdded)
}

func PutTarget(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	decoder := json.NewDecoder(r.Body)
	var newTarget target.Target_data
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	error := decoder.Decode(&newTarget)
	if error != nil {
		log.Print(error)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		log.Print("ERROR GetDb", errorGetDb)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	newTarget, errorUpdateTarget := repositories.UpdateTarget(db, newTarget, id)
	if errorUpdateTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(newTarget)
}

func DeleteTarget(w http.ResponseWriter, r *http.Request) {
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
	targetDeleted, errorDeleteTarget := repositories.DeleteTarget(db, id)
	if errorDeleteTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(targetDeleted)
}
