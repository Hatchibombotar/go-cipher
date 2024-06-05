package format_test

import (
	"testing"

	"github.com/Hatchibombotar/go-cipher/format"
)

func TestFormatToUppercase(t *testing.T) {
	str := "Hello, World!"
	expected := "HELLO, WORLD!"
	op := format.FormatOptions{
		RemoveUnknown: false,
		CaseMode:      format.UpperCaseFormatting,
	}
	actual := format.FormatString(str, &op)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestFormatToLowercase(t *testing.T) {
	str := "Hello, World!"
	expected := "hello, world!"
	op := format.FormatOptions{
		RemoveUnknown: false,
		CaseMode:      format.LowerCaseFormatting,
	}
	actual := format.FormatString(str, &op)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestFormatToSentenceCase(t *testing.T) {
	str := "hi there WHAT IS YOUR NAME? this is a sentence. NICE day huh! bye then."
	expected := "Hi there what is your name? This is a sentence. Nice day huh! Bye then."
	op := format.FormatOptions{
		RemoveUnknown: false,
		CaseMode:      format.SentenceCaseFormatting,
	}
	actual := format.FormatString(str, &op)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestFormatUnchangedCase(t *testing.T) {
	str := "Hello, World!"
	op := format.FormatOptions{
		RemoveUnknown: false,
		CaseMode:      format.UnchangedCaseFormatting,
	}
	actual := format.FormatString(str, &op)
	if str != actual {
		t.Fatalf("Expected: %s, Actual: %s", str, actual)
	}
}

func TestRemoveUnknown(t *testing.T) {
	str := "Hello, World!"
	expected := "hello world"
	op := format.FormatOptions{
		RemoveUnknown: true,
		CaseMode:      format.UnchangedCaseFormatting,
	}
	actual := format.FormatString(str, &op)
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
