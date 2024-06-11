package analysis_test

import (
	"math"
	"testing"

	"github.com/Hatchibombotar/go-cipher/analysis"
	"github.com/Hatchibombotar/go-cipher/utils"
)

func TestDecodeCaesar(t *testing.T) {
	text := utils.ReadFile("../corpus.txt")
	ioc := analysis.MonogramIndexOfCoincidence(text)
	expected := 1.75
	if math.Abs(ioc-expected) > 0.1 {
		t.Fatalf("Index of Coincidence too far from %f, recieved %f", expected, ioc)
	}
}
