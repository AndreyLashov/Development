package main1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Site struct {
	URL   string
	Title string
}

func URL(w http.ResponseWriter, r *http.Request) {
	Site := Site{"site.ru", Site}
	json_data, err := json.Marshal(Site)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
}

func main1() {
	r := mux.NewRouter()

	// Routers
	r.HandleFunc("/", URL).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server and Port 8080 Start")

}
