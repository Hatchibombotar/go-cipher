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

func gcdExtended(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	gcd, x1, y1 := gcdExtended(b%a, a)
	x := y1 - (b/a)*x1
	y := x1

	return gcd, x, y
}

// Function to find the modular inverse of 'base' mod 'number'
func ModularInverse(base int, mod int) (int, bool) {
	gcd, x, _ := gcdExtended(base, mod)

	// If gcd(base, mod) is not 1, inverse doesn't exist
	if gcd != 1 {
		return -1, false // No inverse exists
	}

	// x might be negative, so we normalize it
	inverse := (x%mod + mod) % mod
	return inverse, true
}
