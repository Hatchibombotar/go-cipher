package analysis

import (
	"go-cipher/format"
	"go-cipher/utils"
)

func CountMonograms(text string) map[string]int {
	text = format.FormatString(text, &format.FormatOptions{MakeLowercase: true, RemoveUnknown: true})
	monograms := make(map[string]int)
	for _, char := range utils.Alphabet {
		monograms[string(char)] = 0
	}
	for _, char := range text {
		monograms[string(char)] += 1
	}
	return monograms
}
