package format

import (
	"strings"
)

type FormattingMode int

const (
	UnchangedCaseFormatting FormattingMode = iota
	UpperCaseFormatting
	LowerCaseFormatting
	SentenceCaseFormatting
)

type FormatOptions struct {
	CaseMode      FormattingMode
	RemoveUnknown bool
}

func FormatString(str string, op *FormatOptions) string {
	if op.CaseMode == LowerCaseFormatting {
		str = strings.ToLower(str)
	} else if op.CaseMode == UpperCaseFormatting {
		str = strings.ToUpper(str)
	} else if op.CaseMode == SentenceCaseFormatting {
		new_str := ""
		start_of_sentence := true
		for _, char := range str {
			if char == '.' || char == '!' || char == '?' {
				start_of_sentence = true
			}
			is_letter := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
			if start_of_sentence && is_letter {
				new_str += strings.ToUpper(string(char))
				start_of_sentence = false
			} else {
				new_str += strings.ToLower(string(char))
			}
		}
		str = new_str
	}
	new_str := ""
	if op.RemoveUnknown {
		for _, char := range str {
			if 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' {
				new_str += string(char)
			}
		}
		str = new_str
	}
	return str
}
