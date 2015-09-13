package main

import (
	"log"
	"net/http"
)

func main() {
	LoadConfig()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
