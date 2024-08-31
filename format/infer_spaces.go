package format

import (
	_ "embed"
	"math"
	"slices"
	"strings"
)

//go:embed words_by_frequency.txt
var wordsByFrequency string

// Function to infer spaces in a string without spaces.
// Adapted from https://stackoverflow.com/a/11642687
func InferSpaces(s string) string {
	words := strings.Split(wordsByFrequency, "\n")
	wordCost := make(map[string]float64)
	maxWord := 0

	for i, word := range words {
		if len(word) > maxWord {
			maxWord = len(word)
		}
		wordCost[word] = math.Log(float64(i+1) * math.Log(float64(len(words))))
	}

	cost := make([]float64, len(s)+1)

	// Function to find the best match for the first i characters.
	bestMatch := func(i int) (float64, int) {
		bestCost := math.MaxFloat64
		bestLength := 0

		for k := 1; k <= maxWord && i-k >= 0; k++ {
			c := cost[i-k]
			word := s[i-k : i]
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
