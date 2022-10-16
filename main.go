package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routers
	r.HandleFunc("/", meta).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server and Port 8080 Start")
}
