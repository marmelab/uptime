package main

import (
	"../config"
	"log"
	"net/http"
	"./router"
)

func main() {
	config, err := config.GetConfig("../conf.json")
	if (err != nil) {
		log.Fatal(err)
	}

	router := router.NewRouter()

	port := config["port"].(string)
	log.Println("API listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
