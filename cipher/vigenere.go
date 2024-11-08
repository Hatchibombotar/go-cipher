package cipher

func EncodeVigenereCipher(text string, keyword string) string {
	var grid map[rune][]rune = make(map[rune][]rune)
	for row := range 26 {
		grid_row := make([]rune, 26)
		for col := range 26 {
			grid_row[col] = rune('a' + (row+col)%26)
		}
		grid[rune('a'+row)] = grid_row
	}
	result := ""
	for i, char := range text {
		column := int(rune(keyword[i%len(keyword)]) - 'a')
		result += string(grid[char][column])
	}
	return result
}

func DecodeVigenereCipher(text string, keyword string) string {
	var grid map[rune][]rune = make(map[rune][]rune)
	for row := range 26 {
		grid_row := make([]rune, 26)
		for col := range 26 {
			grid_row[col] = rune('a' + (row+col)%26)
		}
		grid[rune('a'+row)] = grid_row
	}
	result := ""
	for i, char := range text {
		keywordchar := keyword[i%len(keyword)]
		for alpha := range 26 {
			letter := rune('a' + alpha)
			row := grid[letter][keywordchar-'a']
			if char == row {
				result += string(letter)
				break
			}
		}
	}
	return result
}
