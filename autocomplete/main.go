package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("extracting words from file and building trie ...")

	// initialize the db
	db := NewTrie()

	// extract the words from the file and insert into the db
	wordChan := make(chan string, 124)
	go ExtractWordsFromFile("shakespeare-complete.txt", wordChan)
	for word := range wordChan {
		if err := db.Insert(word); err != nil {
			fmt.Println(err)
		}
	}

	// test autocomplete
	fmt.Println(db.Autocomplete("th", 1))

	fmt.Println("Success! Goodbye World!")
	os.Exit(0)
}
