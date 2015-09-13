package main

import (
	"log"
	"net/http"
)

func main() {
	SetupLogging()
	LoadConfig()
	GenerateRandomPageState()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
