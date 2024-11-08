package cipher_test

import (
	"testing"

	"github.com/Hatchibombotar/go-cipher/cipher"
)

func TestEncodeVigenere(t *testing.T) {
	actual := cipher.EncodeVigenereCipher("attackatdawn", "now")
	expected := "nhpnqgnhznkj"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecodeVigenere(t *testing.T) {
	actual := cipher.DecodeVigenereCipher("nhpnqgnhznkj", "now")
	expected := "attackatdawn"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
