package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	URL    string
	Yandex int
	Google int
}

func URL(w http.ResponseWriter, r *http.Request) {
	u1 := User{"site.ru", 50, 60}
	json_data, err := json.Marshal(u1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
}

func main() {
	r := mux.NewRouter()

	// Routers
	r.HandleFunc("/", URL).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server and Port 8080 Start")

}
