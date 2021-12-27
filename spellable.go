package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var (
	spell = flag.String("spell", "", "Word to look up spellables from")
)

// dictionary reads the contents of the dictionary, minus any blank lines.
func dictionary(file string) ([]string, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file %s", file)
	}

	words := strings.Split(string(contents), "\n")

	// Strip trailing blank lines
	for words[len(words)-1] == "" {
		words = words[:len(words)-1]
	}

	sort.Strings(words)

	return words, nil
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

// spellables returns each word that can be spelled
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

	sort.Strings(matches)

	return matches
}

// matchless finds all words that do not spell anything
func matchless(s map[string][]string) []string {
	var m []string

	for word, words := range s {
		if len(words) == 0 {
			m = append(m, word)
		}
	}

	sort.Strings(m)

	return m
}

// longestMatchless returns the longest words that do not spell anything
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

// lookup returns what the word spells and whether the word is in the dictionary
func lookup(l string, dict []string) ([]string, bool) {
	known := false

	for _, word := range dict {
		if word == l {
			known = true
			break
		}
	}

	results := spellables(*spell, dict)

	return results, known
}

func main() {
	flag.Parse()

	// Load the dictionary.
	dict, err := dictionary("dict")
	if err != nil {
		fmt.Println(err)
		return
	}

	// If there is a word to look up do that, else print matchless words.
	if *spell != "" {
		for _, d := range dict {
			if *spell == d {
				fmt.Println("This is a known word!!!")
				break
			}
		}

		s := spellables(*spell, dict)
		fmt.Println(s, len(s))
	} else {
		// This can be sped up greatly. Short circuit after any match is
		// found. We don't need to find all matches; any match is an
		// automatic elimination from consideration.

		// Find all words that are spellable from a given word.
		s := make(map[string][]string)
		for _, word := range dict {
			s[word] = spellables(word, dict)
		}

		fmt.Println(longestMatchless(s))
	}
}
