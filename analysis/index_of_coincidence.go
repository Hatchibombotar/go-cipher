package analysis

import (
	"errors"

	"github.com/Hatchibombotar/go-cipher/format"
)

func MonogramIndexOfCoincidence(text_raw string) (float64, error) {
	text := format.FormatString(text_raw, &format.FormatOptions{
		CaseMode:      format.LowerCaseFormatting,
		RemoveUnknown: true,
	})
	n := float64(len(text))

	if n <= 1.0 {
		return 0, errors.New("cannot calculate index of coincidence for text of length 1")
	}

	ioc := 0.0
	frequency_analysis := CountMonograms(text)
	for i := range 26 {
		letter := 'a' + i
		letter_str := rune(letter)
		N := float64(frequency_analysis[letter_str])

		ioc += (N * (N - 1)) / (n * (n - 1))
	}

	return ioc, nil
}
