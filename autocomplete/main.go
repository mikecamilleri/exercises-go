package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

// TODO: ideally these values should be read from a config file
const (
	limit  = 25
	port   = 9000
	source = "shakespeare-complete.txt"
)

// a global variable is a cheap and easy way to make the trie accessable to an
// httpHandler function
var db *Trie

func main() {
	log.Println("Hello World!")
	log.Println("extracting words from file and building database in memory ...")

	// initialize the db
	db = NewTrie()

	// extract the words from the file and insert into the db
	wordChan := make(chan string, 124)
	go ExtractWordsFromFile(source, wordChan)
	for word := range wordChan {
		if err := db.Insert(word); err != nil {
			log.Println(err)
		}
	}

	// start http server
	log.Printf("listening on port %d ...", port)
	http.HandleFunc("/autocomplete", autocomplete)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

	os.Exit(0)
}

func autocomplete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	term := query.Get("term")
	results := db.Autocomplete(term, limit)

	type response struct {
		Completions []string
	}
	respBytes, err := json.Marshal(response{Completions: results})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
