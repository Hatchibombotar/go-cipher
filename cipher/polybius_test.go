package cipher_test

import (
	"testing"

	"github.com/Hatchibombotar/go-cipher/cipher"
)

func TestPolybiusCipher(t *testing.T) {
	encoded := "1122334455"
	expected := "agntz"
	key := "abcdefghiklmnopqrstuvwxyz"

	actual, err := cipher.DecodePolybiusCipher(encoded, key)
	if err != nil {
		panic(err)
	}
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
