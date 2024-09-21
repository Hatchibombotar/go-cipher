package cipher

import "errors"

func SubstitutionCipher(text string, substitutions map[rune]rune) (string, error) {
	capital_substitutions := make(map[rune]rune)
	for k, v := range substitutions {
		if !('a' <= k && k <= 'z') || !('a' <= v && v <= 'z') {
			return "", errors.New("substitution key or value not a lowercase letter")
		}
		capital_substitutions[k-'a'+'A'] = v - 'a' + 'A'
	}
	decoded := ""
	for _, char := range text {
		substitution, sub_exists := substitutions[char]
		capital_substitution, capital_sub_exists := capital_substitutions[char]
		if sub_exists {
			decoded += string(substitution)
		} else if capital_sub_exists {
			decoded += string(capital_substitution)
		} else {
			decoded += string(char)
		}
	}

	return decoded, nil
}
