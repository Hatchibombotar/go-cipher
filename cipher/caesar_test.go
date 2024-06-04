package cipher_test

import (
	"testing"

	"github.com/Hatchibombotar/go-cipher/cipher"
)

func TestDecodeCaesar(t *testing.T) {
	encoded := "olssv dvysk"
	expected := "hello world"
	actual := cipher.DecodeCaesarCipher(encoded, -7)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
