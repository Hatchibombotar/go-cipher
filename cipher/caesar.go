package cipher

import (
	"go-cipher/utils"
	"strings"
)

func DecodeCaesarCipher(text string, steps int) string {
	decoded := ""
	for _, char := range text {
		alpha_index := strings.IndexRune(utils.Alphabet, char)
		if alpha_index == -1 {
			decoded += string(char)
		} else {
			new_index := (alpha_index + steps) % len(utils.Alphabet)
			if new_index < 0 {
				new_index += len(utils.Alphabet)
			}
			decoded += string(utils.Alphabet[new_index])
		}
	}
	return decoded
}
