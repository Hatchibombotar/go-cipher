package cipher

import (
	"errors"
)

func DecodePolybiusCipher(text string, key string) (string, error) {
	if len(key) != 25 {
		return "", errors.New("polybius cipher only works with keys of length 25")
	}
	decoded := ""
	for i := range len(text) {
		row := text[2*i]
		column := text[2*i+1]

		char := key[row*5+column]

		decoded += string(char)
	}
	return decoded, nil
}
