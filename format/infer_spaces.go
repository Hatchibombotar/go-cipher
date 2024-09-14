package format

import (
	_ "embed"
	"math"
	"slices"
	"strings"
)

//go:embed words_by_frequency.txt
var wordsByFrequency string

var populatedWordData bool = false
var maxWord int = 0
var wordCost map[string]float64

// Function to infer spaces in a string without spaces.
func InferSpaces(s string) string {
	splitlist := []string{}
	wordStart := 0

	for i, char := range s {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {

		} else {
			splitlist = append(splitlist, s[wordStart:i])
			splitlist = append(splitlist, s[i:i+1])
			wordStart = i + 1
		}
	}

	splitlist = append(splitlist, s[wordStart:])

	resultstring := []string{}

	for _, substring := range splitlist {
		if len(substring) <= 1 {
			resultstring = append(resultstring, substring)
		} else {
			resultstring = append(resultstring, inferSpaces(substring))

		}
	}

	return strings.Join(resultstring, "")
}

// Adapted from https://stackoverflow.com/a/11642687;
// used internally after spaces and punctuation are split
func inferSpaces(s string) string {
	if !populatedWordData {
		words := strings.Split(wordsByFrequency, "\n")
		wordCost = make(map[string]float64)
		maxWord = 0

		for i, word := range words {
			if len(word) > maxWord {
				maxWord = len(word)
			}
			wordCost[word] = math.Log(float64(i+1) * math.Log(float64(len(words))))
		}
		populatedWordData = true
	}

	cost := make([]float64, len(s)+1)

	// Function to find the best match for the first i characters.
	bestMatch := func(i int) (float64, int) {
		bestCost := math.MaxFloat64
		bestLength := 0

		for k := 1; k <= maxWord && i-k >= 0; k++ {
			c := cost[i-k]
			word := strings.ToLower(s[i-k : i])
			wordCost, ok := wordCost[word]
			if !ok {
				wordCost = math.MaxFloat64 // Large cost if the word is not found.
			}
			if c+wordCost < bestCost {
				bestCost = c + wordCost
				bestLength = k
			}
		}
		return bestCost, bestLength
	}

	// Build the cost array.
	cost[0] = 0
	for i := 1; i <= len(s); i++ {
		c, _ := bestMatch(i)
		cost[i] = c
	}

	// Backtrack to recover the minimal-cost string.
	var out []string
	i := len(s)
	for i > 0 {
		_, k := bestMatch(i)
		out = append(out, s[i-k:i])
		i -= k
	}

	// Return the final string with spaces.
	slices.Reverse(out)
	return strings.Join(out, " ")
}
