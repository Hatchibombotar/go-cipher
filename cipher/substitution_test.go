package cipher_test

import (
	"fmt"
	"testing"

	"github.com/Hatchibombotar/go-cipher/cipher"
)

func TestSubstitutionCipher(t *testing.T) {
	encoded := "qwerty"
	expected := "asdfgh"
	subsitution_map := make(map[rune]rune)
	subsitution_map['q'] = 'a'
	subsitution_map['w'] = 's'
	subsitution_map['e'] = 'd'
	subsitution_map['r'] = 'f'
	subsitution_map['t'] = 'g'
	subsitution_map['y'] = 'h'

	actual := cipher.SubstitutionCipher(encoded, subsitution_map)
	fmt.Println(actual)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
