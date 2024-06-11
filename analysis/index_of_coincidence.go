package analysis

import (
	"github.com/Hatchibombotar/go-cipher/format"
)

func MonogramIndexOfCoincidence(text_raw string) float64 {
	text := format.FormatString(text_raw, &format.FormatOptions{
		CaseMode:      format.LowerCaseFormatting,
		RemoveUnknown: true,
	})
	N := float64(len(text))

	ioc := 0.0
	frequency_analysis := CountMonograms(text)
	for i := range 26 {
		letter := 'a' + i
		letter_str := string(rune(letter))
		n := float64(frequency_analysis[letter_str])

		ioc += (n * (n - 1)) / (N * (N - 1))
	}

	return ioc * 26
}
