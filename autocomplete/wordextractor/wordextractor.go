// Package wordextractor ...
//
// An althernative way to implement this would be an iterator.
package wordextractor

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

// ExtractWordsFromFile ...
func ExtractWordsFromFile(path string, wordChan chan<- string) error {
	// open the file and defer close
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// create the scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// scan
	for scanner.Scan() {
		// get a text fragment
		fragment := scanner.Text()

		// here we split the text fragment on `--`. A better solution to the
		// em-dash problem would be to write a custome splitter function for
		// the scanner.
		words := strings.Split(fragment, "--")

		// clean and emit the words
		for _, v := range words {
			cleanedWord := cleanWord(v)
			if len(cleanedWord) > 0 {
				wordChan <- cleanedWord
			}
		}
	}
	if err := scanner.Err(); err != nil {
		// TODO: error handlig should be improved
		return err
	}

	return nil
}

// cleanWord removes non-alphabet characters from the beginning and and of the
// word and ensures that the word only contains alphabet letters, hyphens, and
// apostrophies internally.
func cleanWord(word string) string {
	// remove unacceptable characters from word
	var stageOneWord strings.Builder
	for _, v := range word {
		if unicode.IsLetter(v) || (v == '-') || (v == '\'') || (v == '’') {
			stageOneWord.WriteRune(v)
		}
	}

	// strip characters excepted in stage one from left and right
	stageTwoWord := strings.TrimRight(strings.TrimLeft(stageOneWord.String(), "-'’"), "-'’")

	// if the word contains any repeated special characters internally, it is
	// not a word so return empty
	if strings.Contains(stageTwoWord, "--") || strings.Contains(stageTwoWord, "''") || strings.Contains(stageTwoWord, "’’") {
		return ""
	}

	return stageTwoWord
}
