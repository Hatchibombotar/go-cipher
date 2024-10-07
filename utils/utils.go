package utils

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func ReadFile(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func WriteFile(path string, content []byte) {
	file, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintln("Error creating file:", err))
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		panic(fmt.Sprintln("Error writing to file:", err))
	}
}

var Alphabet string = "abcdefghijklmnopqrstuvwxyz"

func ExportExcel() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func IsAlpha(char rune) bool {
	includeCaps := true
	if includeCaps {
		return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
	} else {
		return 'a' <= char && char <= 'z'
	}
}
