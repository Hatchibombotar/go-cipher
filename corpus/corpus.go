package main

import (
	"encoding/json"
	"fmt"
	"go-cipher/analysis"
	"go-cipher/utils"
)

func main() {
	fmt.Println("analysing corpus...")

	corpus := utils.ReadFile("./corpus.txt")
	analysis_data := analysis.FullAnalysis(corpus)

	jsondata, err := json.Marshal(analysis_data)
	if err != nil {
		panic(err)
	}
	utils.WriteFile("./corpus.json", jsondata)
}
