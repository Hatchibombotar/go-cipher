package format

import (
	"strings"

	"github.com/Hatchibombotar/go-cipher/utils"
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
	}
	new_str := ""
	if op.RemoveUnknown {
		for _, char := range str {
			if strings.ContainsRune(utils.Alphabet, char) {
				new_str += string(char)
			}
		}
		str = new_str
	}
	return str
}
