package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routers
	r.HandleFunc("/", Marshal).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server and Port 8080 Start")
}

func Marshal(w http.ResponseWriter, r *http.Request) {

	urls := []string{
		"https://vl.ru",
		"https://digital.gov.ru/ru/",
		"http://rubycode.ru/",
	}

	var wg sync.WaitGroup

	for _, u := range urls {

		wg.Add(1)
		go func(url string) {

			defer wg.Done()
			w.Header().Set("Content-Type", "application/json")
			content := doReq(url)
			title := getTitle(content)
			mapVar1 := map[string]string{"url": url, "title": title}
			mapVar2, _ := json.Marshal(mapVar1)

			fmt.Sprintln(string(mapVar2))

		}(u)
	}

	wg.Wait()
}

func doReq(url string) (content string) {

	resp, err := http.Get(url)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func getTitle(content string) (title string) {

	re := regexp.MustCompile("<title>(.*)</title>")

	parts := re.FindStringSubmatch(content)

	if len(parts) > 0 {
		return parts[1]
	} else {
		return "no title"
	}
}
