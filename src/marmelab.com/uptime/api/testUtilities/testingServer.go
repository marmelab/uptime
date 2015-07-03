package main

import (
	"net/http"
	"log"
	"../router"
)


func main() {
	config := map[string]string{"port": "8384"}
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+config["port"], router))
}
