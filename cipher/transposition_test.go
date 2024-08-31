package cipher_test

import (
	"testing"

	"github.com/Hatchibombotar/go-cipher/cipher"
)

func TestColumnarTranspositionNoChange(t *testing.T) {
	encoded := "abcdefghi"
	key := []int{0, 1, 2}
	expected := "abcdefghi"

	actual, _ := cipher.ColumnarTransposition(encoded, 3, key)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestColumnarTranspositionWithChange(t *testing.T) {
	encoded := "abcdefghi"
	key := []int{2, 1, 0}
	expected := "cbafedihg"

	actual, err := cipher.ColumnarTransposition(encoded, 3, key)
	if err != nil {
		panic(err)
	}
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
