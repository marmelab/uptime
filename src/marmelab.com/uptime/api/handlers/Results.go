package handlers

import (
	"../../poller"
	"../repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetResults(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		error500(w, errorGetDb, " errorGetDb")
	}
	rows, errGetResults := repositories.GetResults(db)
	if errGetResults != nil {
		error500(w, errGetResults, "errGetResults")
	}
	results := make([]poller.Response, 0)
	for rows.Next() {
		var newResult poller.Response
		error := rows.Scan(&newResult.Id, &newResult.Target_id, &newResult.Destination, &newResult.Status, &newResult.Time, &newResult.Created_at)
		if error != nil {
			error500(w, error, " error San GetResults")
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
		error500(w, errorGetDb, "errorGetDb")
	}
	newResult, errorGetTarget := repositories.GetResult(db, id)
	if errorGetTarget != nil {
		error500(w, errorGetTarget, " errorGetTarget")
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
		error500(w, error, "error decode PostResult")
	}
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		error500(w, errorGetDb, " errorGetDb")
	}
	resultAdded, errAddTarget := repositories.AddResult(db, newResult)
	if errAddTarget != nil {
		error500(w, errAddTarget, " error AddResult")
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
		error500(w, error, " error decode PutResult")
	}
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		error500(w, errorGetDb, " errorGetDb")
	}
	resultUpdated, errorUpdateTarget := repositories.UpdateResult(db, newResult, id)
	if errorUpdateTarget != nil {
		error500(w, errorUpdateTarget, " errorUpdateResult")
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
		error500(w, errorGetDb, "errorGetDb")
	}
	resultDeleted, errorDeleteTarget := repositories.DeleteResult(db, id)
	if errorDeleteTarget != nil {
		error500(w, errorDeleteTarget, "errorDeleteResult")
	}
	json.NewEncoder(w).Encode(resultDeleted)
}
