package main

import (
	"github.com/Hatchibombotar/go-cipher/analysis"
	"github.com/Hatchibombotar/go-cipher/cipher"
	"github.com/Hatchibombotar/go-cipher/format"
	"github.com/Hatchibombotar/go-cipher/utils"
)

func main() {
	corpus := utils.ReadFile("./corpus.txt")
	analysis_corpus := analysis.FullAnalysis(corpus)
	challenge_1a := utils.ReadFile("./challenges/1a.txt")

	decoded := cipher.DecodeCaesarCipher(
		format.FormatString(
			challenge_1a, &format.FormatOptions{MakeLowercase: true, RemoveUnknown: false},
		), 4,
	)
	analysis_challenge := analysis.FullAnalysis(decoded)

	utils.WriteFile("result.txt", []byte(decoded))

	analysis.ComparitiveExcelAnalysis(analysis_challenge, analysis_corpus)
}
