package trie

import (
	"errors"
	"strings"
	"unicode"
)

var (
	// ErrInvalidWord is returned when a word value passed is invalid
	ErrInvalidWord = errors.New("invalid word")
)

// Trie holds a single trie data structure.
type Trie struct {
	root *Node
}

// Node represents a single node in our trie.
type Node struct {
	// children maps characters to the next node.
	children map[rune]*Node
	// wordcount contains a count of the number of times the word created by the
	// chain of characters from the root node through this node is found in the
	// source document.
	wordCount uint
}

// NewTrie creates a new empty Trie.
func NewTrie() *Trie {
	return &Trie{}
}

// Insert inserts a new word into the Trie.
func (t *Trie) Insert(word string) error {
	// validate and make lower case
	if !isValidWord(word) {
		return ErrInvalidWord
	}
	word = strings.ToLower(word)

	// walk and insert if necessary
	node := t.root
	for _, v := range word {
		// if no child node for this rune exists, create one
		if _, ok := node.children[v]; !ok {
			// avoid write to nil map
			if node.children == nil {
				node.children = map[rune]*Node{}
			}
			node.children[v] = &Node{}
		}

		// walk to the child node
		node = node.children[v]
	}

	// we are at the node representing the last rune in the word
	// increment the wordCount on that node
	node.wordCount++

	return nil
}

// Autocomplete returns a slice of autocomplete suggestions sorted from most
// to least frequent. An empty slice indicates that no suggestions were found.
func (t *Trie) Autocomplete(wordFragment string) []string {
	// validate and make lower case
	if !isValidWord(wordFragment) {
		// alternatively we could return an error here, but I think that not
		// finding `#^jkjk` isn't really an error. `#^jkjk` simply isn't in the
		// Trie.
		return []string{}
	}
	wordFragment = strings.ToLower(wordFragment)

	// walk to end of wordFragment
	node := t.root
	for _, v := range wordFragment {
		// if no child node for this rune exists, return empty result
		if _, ok := node.children[v]; !ok {
			return []string{}
		}
		// walk to the child node
		node = node.children[v]
	}

	// we are at the node representing the last rune in wordFragment

	// a slice of words and counts
	type result struct {
		word  string
		count uint
	}
	results := []result{}

	// recursively complete the word fragment using anonymous function
	// variable declared first so that it may be used in the anonymous function
	var recComplete func(wordFragment string, node *Node, results []result)
	recComplete = func(wordFragment string, node *Node, results []result) {
		// if we are at a word end, add the word to our results
		if node.wordCount >= 0 {
			results = append(results, result{word: wordFragment, count: node.wordCount})
		}

		// if there are any children, handle them recursively,\
		for k, v := range node.children {
			var nextWordFragment strings.Builder
			nextWordFragment.WriteString(wordFragment)
			nextWordFragment.WriteRune(k)
			recComplete(nextWordFragment.String(), v, results)
		}
	}
	recComplete(wordFragment, node, results)

	return nil
}

// isValidWord ...
func isValidWord(word string) bool {
	// word length is greatter than zero
	if len(word) == 0 {
		return false
	}

	// word contains only letters, hyphens, and apostrophies
	for _, v := range word {
		if !unicode.IsLetter(v) && (v != '-') && (v != '\'') && (v != '’') {
			return false
		}
	}

	// word contains no repeated hyphens or apostrophies
	if strings.Contains(word, "--") || strings.Contains(word, "''") || strings.Contains(word, "’’") {
		return false
	}

	return true
}
