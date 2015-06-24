package handlers

import (
	"../../poller"
	"../repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

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
	results := make([]poller.Response, 0)
	for rows.Next() {
		var newResult poller.Response
		error := rows.Scan(&newResult.Id, &newResult.Target_id, &newResult.Destination, &newResult.Status, &newResult.Time, &newResult.Created_at)
		if error != nil {
			log.Print("ERROR Scan GetResults ", error)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		results = append(results, newResult)
	}
	json.NewEncoder(w).Encode(results)
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
	newResult, errorGetTarget := repositories.GetResult(db, id)
	if errorGetTarget != nil {
		log.Print("ERROR GetResult", errorGetTarget)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(newResult)
}

func PostResult(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	decoder := json.NewDecoder(r.Body)
	var newResult poller.Response
	error := decoder.Decode(&newResult)
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
	resultAdded, errAddTarget := repositories.AddResult(db, newResult)
	if errAddTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(resultAdded)
}

func PutResult(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	decoder := json.NewDecoder(r.Body)
	var newResult poller.Response
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	error := decoder.Decode(&newResult)
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
	resultUpdated, errorUpdateTarget := repositories.UpdateResult(db, newResult, id)
	if errorUpdateTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(resultUpdated)
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
	resultDeleted, errorDeleteTarget := repositories.DeleteResult(db, id)
	if errorDeleteTarget != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(resultDeleted)
}
