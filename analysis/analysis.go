package analysis

import (
	"fmt"
	"go-cipher/format"
	"go-cipher/utils"
	"slices"

	"github.com/xuri/excelize/v2"
)

type AnalysisData struct {
	Length     int            `json:"length"`
	Monograms  map[string]int `json:"monograms"`
	Tetragrams map[string]int `json:"tetragrams"`
}

func FullAnalysis(text string) AnalysisData {
	corpus := format.FormatString(text, &format.FormatOptions{MakeLowercase: true, RemoveUnknown: true})
	corpus_monograms := CountMonograms(corpus)
	corpus_tetragrams := CountNGrams(corpus, 4)

	analysis_data := AnalysisData{
		Length:     len(corpus),
		Monograms:  corpus_monograms,
		Tetragrams: corpus_tetragrams,
	}
	return analysis_data
}

func ComparitiveExcelAnalysis(analysis1 AnalysisData, analysis2 AnalysisData) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.DeleteSheet("Sheet1")
	f.NewSheet("FrequencyAnalysis")

	f.SetCellValue("FrequencyAnalysis", "A1", "Character")
	f.SetCellValue("FrequencyAnalysis", "B1", "Frequency 1")
	f.SetCellValue("FrequencyAnalysis", "C1", "Frequency 1 (%)")
	f.SetCellValue("FrequencyAnalysis", "D1", "Frequency 2")
	f.SetCellValue("FrequencyAnalysis", "E1", "Frequency 2 (%)")

	for i, char := range utils.Alphabet {
		f.SetCellValue("FrequencyAnalysis", fmt.Sprint("A", i+2), string(char))
		f.SetCellValue("FrequencyAnalysis", fmt.Sprint("B", i+2), analysis1.Monograms[string(char)])
		f.SetCellValue("FrequencyAnalysis", fmt.Sprint("C", i+2), float32(analysis1.Monograms[string(char)])/float32(analysis1.Length))
		f.SetCellValue("FrequencyAnalysis", fmt.Sprint("D", i+2), analysis2.Monograms[string(char)])
		f.SetCellValue("FrequencyAnalysis", fmt.Sprint("E", i+2), float32(analysis2.Monograms[string(char)])/float32(analysis2.Length))
	}

	f.AddChart("FrequencyAnalysis", "H3", &excelize.Chart{
		Type: excelize.Col3DClustered,
		Series: []excelize.ChartSeries{
			{
				Name:       "FrequencyAnalysis!$C$1",
				Categories: "FrequencyAnalysis!$A$2:$A$27",
				Values:     "FrequencyAnalysis!$C$2:$C$27",
			},
			{
				Name:       "FrequencyAnalysis!$E$1",
				Categories: "FrequencyAnalysis!$A$2:$A$27",
				Values:     "FrequencyAnalysis!$E$2:$E$27",
			},
		},
		Title: []excelize.RichTextRun{
			{
				Text: "Frequency",
			},
		},
	})

	_ = f.AddTable("FrequencyAnalysis", &excelize.Table{
		Range:             "A1:E27",
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   true,
		ShowLastColumn:    true,
		ShowColumnStripes: true,
	})

	f.NewSheet("TetragramAnalysis")

	f.SetCellValue("TetragramAnalysis", "A1", "Tetragram")
	f.SetCellValue("TetragramAnalysis", "B1", "Frequency 1")
	f.SetCellValue("TetragramAnalysis", "C1", "Frequency 1 (%)")
	f.SetCellValue("TetragramAnalysis", "D1", "Frequency 2")
	f.SetCellValue("TetragramAnalysis", "E1", "Frequency 2 (%)")

	tetragrams := []string{}

	for i := range analysis1.Tetragrams {
		tetragrams = append(tetragrams, i)
	}
	for i := range analysis2.Tetragrams {
		tetragrams = append(tetragrams, i)
	}

	slices.Sort(tetragrams)
	tetragrams = slices.Compact(tetragrams)

	for i, tetragram := range tetragrams {
		f.SetCellValue("TetragramAnalysis", fmt.Sprint("A", i+2), string(tetragram))
		f.SetCellValue("TetragramAnalysis", fmt.Sprint("B", i+2), analysis1.Tetragrams[string(tetragram)])
		f.SetCellValue("TetragramAnalysis", fmt.Sprint("C", i+2), float32(analysis1.Tetragrams[string(tetragram)])/float32(analysis1.Length))
		f.SetCellValue("TetragramAnalysis", fmt.Sprint("D", i+2), analysis2.Tetragrams[string(tetragram)])
		f.SetCellValue("TetragramAnalysis", fmt.Sprint("E", i+2), float32(analysis2.Tetragrams[string(tetragram)])/float32(analysis2.Length))
	}

	if err := f.SaveAs("Analysis.xlsx"); err != nil {
		fmt.Println(err)
	}
}
