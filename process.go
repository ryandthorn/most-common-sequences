package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type wordGroup struct {
	text  string
	count int
}

func (wg *wordGroup) Text() string {
	return wg.text
}

func (wg *wordGroup) Count() int {
	return wg.count
}

// processFile is the worker function for processing file contents.
// Returns a sorted list of the 100 most common three-word groups in descending order
func processFile(f *file) ([]wordGroup, error) {
	if f == nil {
		return nil, errors.New("cannot process nil file")
	}

	formatted := formatContents(f.ContentsToString())
	occurrences := countOccurrences(formatted)
	return processWordGroups(occurrences), nil
}

// formatContents trims external whitespace, converts text to lowercase, and trims leading/trailing punctuation on each word
func formatContents(s string) string {
	words := strings.Fields(strings.Trim(strings.ToLower(s), " "))
	for i, word := range words {
		for isRemovablePunctuation(word[0]) {
			word = word[1:]
			if len(word) == 0 {
				break
			}
		}
		if len(word) == 0 {
			continue
		}
		for isRemovablePunctuation(word[len(word)-1]) {
			word = word[:len(word)-1]
			if len(word) == 0 {
				break
			}
		}
		words[i] = word
	}
	return strings.Join(words, " ")
}

func isRemovablePunctuation(b byte) bool {
	for _, p := range []byte(".,'\"“”?!()[]:;-_") {
		if b == p {
			return true
		}
	}
	return false
}

func countOccurrences(s string) map[string]int {
	occurrences := make(map[string]int)

	words := strings.Fields(s)
	if len(words) < 3 {
		return occurrences
	}

	for i := 0; i+2 < len(words); i++ {
		triplet := fmt.Sprintf("%s %s %s", words[i], words[i+1], words[i+2])
		_, exists := occurrences[triplet]
		if !exists {
			occurrences[triplet] = 1
		} else {
			occurrences[triplet] += 1
		}
	}
	return occurrences
}

func processWordGroups(occurrences map[string]int) []wordGroup {
	wgs := make([]wordGroup, len(occurrences))
	i := 0
	for text, count := range occurrences {
		wgs[i] = wordGroup{text: text, count: count}
		i++
	}
	return sortWordGroups(wgs)
}

func sortWordGroups(wordGroups []wordGroup) []wordGroup {
	sort.Slice(wordGroups, func(i, j int) bool { return wordGroups[i].count > wordGroups[j].count })
	return wordGroups
}
