package cipher

import (
	"errors"
)

func DecodePolybiusCipher(text string, key string) (string, error) {
	if len(key) != 25 {
		return "", errors.New("polybius cipher only works with keys of length 25")
	}
	if len(text)%2 == 1 {
		return "", errors.New("polybius cipher input must be of even length")
	}
	decoded := ""
	for i := range len(text) / 2 {
		row_char := text[2*i]
		column_char := text[2*i+1]

		if row_char < '1' || row_char > '5' {
			return "", errors.New("found row number outside of range 1-5")
		}
		if column_char < '1' || column_char > '5' {
			return "", errors.New("found column number outside of range 1-5")
		}
		row := row_char - 49
		column := column_char - 49

		char := key[row*5+column]

		decoded += string(char)
	}
	return decoded, nil
}
