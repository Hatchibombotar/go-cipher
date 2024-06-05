package analysis

import (
	"github.com/Hatchibombotar/go-cipher/format"
	"github.com/Hatchibombotar/go-cipher/utils"
)

func CountMonograms(text string) map[string]int {
	text = format.FormatString(text, &format.FormatOptions{CaseMode: format.LowerCaseFormatting, RemoveUnknown: true})
	monograms := make(map[string]int)
	for _, char := range utils.Alphabet {
		monograms[string(char)] = 0
	}
	for _, char := range text {
		monograms[string(char)] += 1
	}
	return monograms
}
