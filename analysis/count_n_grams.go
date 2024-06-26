package analysis

import (
	"github.com/Hatchibombotar/go-cipher/format"
)

func CountNGrams(text string, nGramSize int) map[string]int {
	text = format.FormatString(text, &format.FormatOptions{CaseMode: format.LowerCaseFormatting, RemoveUnknown: true})
	ngrams := make(map[string]int)
	for i := range text {
		if i >= len(text)-nGramSize {
			break
		}
		ngram := text[i : i+nGramSize]
		ngrams[ngram] += 1
	}
	return ngrams
}
