package handlers

import (
	"../repositories"
	"../target"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


func GetTargets(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	perPage, _ := strconv.Atoi(vars["perPage"])
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		error500(w, errorGetDb, "errorGetDb")
	}
	rows, errGetTargets := repositories.GetTargetsWithLastResult(db,page,perPage)
	if errGetTargets != nil {
		error500(w, errGetTargets, " error GetTargets")
		return
	}
	targets := make([]target.Target_data, 0)
	for rows.Next() {
		var newTarget target.Target_data
		error := rows.Scan(&newTarget.Id, &newTarget.Destination, &newTarget.Status)
		if error != nil {
			error500(w, error, " error scan GetTargets")
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
		error500(w, errorGetDb, " errorGetDb")
		return
	}
	newTarget, errorGetTarget := repositories.GetTarget(db, id)
	if errorGetTarget != nil {
		error500(w, errorGetTarget, "error GetTarget")
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
		error500(w, error, "")
	}
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		error500(w, errorGetDb, "error GetDb")
	}
	targetAdded, errAddTarget := repositories.AddTarget(db, newTarget)
	if errAddTarget != nil {
		error500(w, errAddTarget, " error AddTarget")
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
		error500(w, error, " error decode PutTarget")
	}
	db, errorGetDb := repositories.GetDb()
	if errorGetDb != nil {
		error500(w, errorGetDb, " errorGetDb")
	}
	newTarget, errorUpdateTarget := repositories.UpdateTarget(db, newTarget, id)
	if errorUpdateTarget != nil {
		error500(w, errorUpdateTarget, " errorUpdateTarget")
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
		error500(w, errorGetDb, " errorGetDb")
	}
	targetDeleted, errorDeleteTarget := repositories.DeleteTarget(db, id)
	if errorDeleteTarget != nil {
		error500(w, errorDeleteTarget, " errorDeleteTarget")
	}
	json.NewEncoder(w).Encode(targetDeleted)
}
