package format

import (
	"go-cipher/utils"
	"strings"
)

type FormatOptions struct {
	MakeLowercase bool
	RemoveUnknown bool
}

func FormatString(str string, op *FormatOptions) string {
	if op.MakeLowercase {
		str = strings.ToLower(str)
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
