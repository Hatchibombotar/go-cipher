package analysis_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Hatchibombotar/go-cipher/analysis"
)

func TestBigrams(t *testing.T) {
	text := "Hello there how are you?"
	expected := map[string]int{
		"he": 2,
		"el": 1,
		"ll": 1,
		"lo": 1,
		"th": 1,
		"er": 1,
		"re": 2,
		"ho": 1,
		"ow": 1,
		"ar": 1,
		"yo": 1,
		"ou": 1,
	}
	ngrams := analysis.CountNGrams(text, 2)

	if !reflect.DeepEqual(expected, ngrams) {
		fmt.Println("expected", expected)
		fmt.Println("recieved", ngrams)
		t.Fatal("Expected:", expected, "Actual:", ngrams)
	}
}
func TestBounds(t *testing.T) {
	text := "hihihi"
	expected := map[string]int{
		"hi": 3,
		"ih": 2,
	}
	ngrams := analysis.CountNGrams(text, 2)

	if !reflect.DeepEqual(expected, ngrams) {
		fmt.Println("expected", expected)
		fmt.Println("recieved", ngrams)
		t.Fatal("Expected:", expected, "Actual:", ngrams)
	}
}

func TestTrigrams(t *testing.T) {
	text := "Hello there how are you?"
	expected := map[string]int{
		"hel": 1,
		"ell": 1,
		"llo": 1,
		"the": 1,
		"her": 1,
		"ere": 1,
		"how": 1,
		"are": 1,
		"you": 1,
	}
	ngrams := analysis.CountNGrams(text, 3)

	if !reflect.DeepEqual(expected, ngrams) {
		fmt.Println("expected", expected)
		fmt.Println("recieved", ngrams)
		t.Fatal("Expected:", expected, "Actual:", ngrams)
	}
}
