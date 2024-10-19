package analysis

import (
	"github.com/Hatchibombotar/go-cipher/format"
	"github.com/Hatchibombotar/go-cipher/utils"
)

// counts the frequency of monograms in text, converting all letters to lowercase.
func CountMonograms(text string) map[rune]int {
	text = format.FormatString(text, &format.FormatOptions{CaseMode: format.LowerCaseFormatting, RemoveUnknown: true})
	monograms := make(map[rune]int)
	for _, char := range utils.Alphabet {
		monograms[char] = 0
	}
	for _, char := range text {
		monograms[char] += 1
	}
	return monograms
}
