package analysis

import (
	"github.com/Hatchibombotar/go-cipher/format"
	"github.com/Hatchibombotar/go-cipher/utils"
)

func CountNGrams(text string, nGramSize int) map[string]int {
	text = format.FormatString(text, &format.FormatOptions{CaseMode: format.LowerCaseFormatting, RemoveUnknown: false})
	ngrams := make(map[string]int)
	for i := range text {
		if i >= len(text)-nGramSize {
			break
		}

		skipNGram := false
		for position := range nGramSize {
			if !utils.IsAlpha(rune(text[i+position])) {
				skipNGram = true
				break
			}
		}
		if skipNGram {
			continue
		}

		ngram := text[i : i+nGramSize]
		ngrams[ngram] += 1
	}
	return ngrams
}
