package main

import (
	"log"
	"net/http"
	"./router"
)

func main() {

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8383", router))
}
