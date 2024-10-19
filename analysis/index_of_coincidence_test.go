package analysis_test

import (
	"math"
	"testing"

	"github.com/Hatchibombotar/go-cipher/analysis"
	"github.com/Hatchibombotar/go-cipher/corpus"
)

func TestIoC(t *testing.T) {
	text := corpus.GetRawCorpus()
	ioc, err := analysis.MonogramIndexOfCoincidence(text)
	if err != nil {
		t.Fatalf("Error is not nil.")
	}
	expected := 1.75
	if math.Abs(ioc-expected) > 0.1 {
		t.Fatalf("Index of Coincidence too far from %f, recieved %f", expected, ioc)
	}
}

func TestIoCSingleLetter(t *testing.T) {
	text := "w"
	_, err := analysis.MonogramIndexOfCoincidence(text)
	if err == nil {
		t.Fatalf("Expected an error, but did not get one.")
	}
}
