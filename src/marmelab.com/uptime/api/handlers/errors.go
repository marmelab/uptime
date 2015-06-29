package handlers

import (
	"log"
	"net/http"
)

func SetCors(w *http.Header) {
	w.Set("Access-Control-Allow-Origin", "*")
	w.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func error500(w http.ResponseWriter, err error, message string) {
	log.Print(message, err)
	http.Error(w, http.StatusText(500), 500)
	return
}

func SetContentType(w *http.Header) {
	w.Set("Content-Type", "application/json")
}
