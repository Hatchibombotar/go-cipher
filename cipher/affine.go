package cipher

import (
	"errors"

	"github.com/Hatchibombotar/go-cipher/analysis"
	"github.com/Hatchibombotar/go-cipher/utils"
)

func EncodeAffineCipher(text string, a int, b int) string {
	encoded := ""
	for _, char := range text {
		if 'a' <= char && char <= 'z' {
			alpha_index := int(char) - 'a'
			new_index := (a*alpha_index + b) % 26
			for new_index < 0 {
				new_index += 26
			}
			encoded += string(rune(new_index + 'a'))
		} else if 'A' <= char && char <= 'Z' {
			alpha_index := int(char) - 'A'
			new_index := (a*alpha_index + b) % 26
			for new_index < 0 {
				new_index += 26
			}
			encoded += string(rune(new_index + 'A'))
		} else {
			encoded += string(char)
		}
	}
	return encoded
}

func DecodeAffineCipher(text string, a int, b int) (string, error) {
	decoded := ""
	inverse, inverse_exists := utils.ModularInverse(a, 26)
	if !inverse_exists {
		return "", errors.New("there does not exist an inverse for the provided multiplier value")
	}
	for _, char := range text {
		if 'a' <= char && char <= 'z' {
			alpha_index := int(char) - 'a'
			new_index := inverse * (alpha_index - b) % 26
			for new_index < 0 {
				new_index += 26
			}
			decoded += string(rune(new_index + 'a'))
		} else if 'A' <= char && char <= 'Z' {
			alpha_index := int(char) - 'A'
			new_index := inverse * (alpha_index - b) % 26
			for new_index < 0 {
				new_index += 26
			}
			decoded += string(rune(new_index + 'A'))
		} else {
			decoded += string(char)
		}
	}
	return decoded, nil
}

func AttemptCrackAffineCipher(text string) (int, int, error) {
	frequency := analysis.CountMonograms(text)
	most_frequent := '.' // the result of the cipher applied to 'e'
	most_frequent_count := -1
	second_most_frequent := '.' // the result of the cipher applied to 't'
	second_most_frequent_count := -1

	// find the first and second most frequent letters
	for char, freq := range frequency {
		if freq > int(most_frequent_count) {
			second_most_frequent = most_frequent
			second_most_frequent_count = most_frequent_count
			most_frequent = char
			most_frequent_count = freq
		} else if freq > int(second_most_frequent_count) {
			second_most_frequent = char
			second_most_frequent_count = freq
		}
	}
	if most_frequent_count == -1 || second_most_frequent_count == -1 {
		return -1, -1, errors.New("not enough letters in the provided text to crack")
	}

	// two simultanious equations to be solved:
	// 't'*a + b = second_most_frequent (mod 26)
	// 'e'*a + b = most_frequent (mod 26)

	// subtract the left and right sides of the equations
	lhs := ('t' - 'a') - ('e' - 'a')
	rhs := (second_most_frequent - 'a') - (most_frequent - 'a')
	for rhs < 0 {
		rhs += 26
	}
	// (the inverse of lhs used for calculations)
	lhs_inverse, inverse_exists := utils.ModularInverse(int(lhs), 26)
	if !inverse_exists {
		return -1, -1, errors.New("no inverse exists for predicted values")
	}
	// and divide by lhs to find a
	a := lhs_inverse * int(rhs) % 26
	for a < 0 {
		a += 26
	}

	// put value of a back into equation to find b
	b := int(most_frequent-'a') - a*int('e'-'a')
	for b < 0 {
		b += 26
	}

	return a, b, nil
}
