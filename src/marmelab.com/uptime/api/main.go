package main

import (
	"log"
	"net/http"
	"./router"
	"../poller"
)

func main() {
	config := poller.RetrieveConfDbFromJsonFile("../conf.json")
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+config["port"].(string), router))
}
