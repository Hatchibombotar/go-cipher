package cipher

func SubstitutionCipher(text string, substitutions map[rune]rune) string {
	decoded := ""
	for _, char := range text {
		substitution, sub_exists := substitutions[char]
		if sub_exists {
			decoded += string(substitution)
		} else {
			decoded += string(char)
		}
	}

	return decoded
}
