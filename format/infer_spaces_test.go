package format_test

import (
	"testing"

	"github.com/Hatchibombotar/go-cipher/format"
)

func TestInferSpaces(t *testing.T) {
	str := "hellotherehowareyoujodie"
	expected := "hello there how are you jodie"
	actual := format.InferSpaces(str)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
