package cipher

func DecodeCaesarCipher(text string, steps int) string {
	decoded := ""
	for _, char := range text {
		if 'a' <= char && char <= 'z' {
			alpha_index := int(char) - 'a'
			new_index := (alpha_index + steps) % 26
			for new_index < 0 {
				new_index += 26
			}
			decoded += string(rune(new_index + 'a'))
		} else if 'A' <= char && char <= 'Z' {
			alpha_index := int(char) - 'A'
			new_index := (alpha_index + steps) % 26
			for new_index < 0 {
				new_index += 26
			}
			decoded += string(rune(new_index + 'A'))
		} else {
			decoded += string(char)
		}
	}
	return decoded
}
