package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// dict reads the contents of the dictionary, minus any blank lines.
func dict(file string) ([]string, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file %s", file)
	}

	lines := strings.Split(string(contents), "\n")

	// Strip trailing blank lines
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

// spellable returns whether w2 is spellable from the letters in w1.
func spellable(w1, w2 string) bool {
	for _, c2 := range w2 {
		found := false
		for i, c1 := range w1 {
			if c1 == c2 {
				w1 = w1[:i] + w1[i+1:]
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func spellables(word string, words []string) []string {
	var matches []string

	for _, candidate := range words {
		if word == candidate {
			continue
		}
		if spellable(word, candidate) {
			matches = append(matches, candidate)
		}
	}

	return matches
}

func matchless(s map[string][]string) []string {
	var matchless []string

	for word, words := range s {
		if len(words) == 0 {
			matchless = append(matchless, word)
		}
	}

	return matchless
}

func longestMatchless(s map[string][]string) []string {
	m := matchless(s)

	maxLen := 0

	for _, word := range m {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	var longest []string
	for _, word := range m {
		if len(word) < maxLen {
			continue
		}
		longest = append(longest, word)
	}

	return longest
}

func main() {
	// Load the dictionary.
	words, err := dict("dict")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find all words that are spellable from a given word.
	s := make(map[string][]string)
	for _, word := range words {
		s[word] = spellables(word, words)
	}

	fmt.Println(longestMatchless(s))
}
