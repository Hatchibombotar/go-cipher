package cipher

import (
	"errors"
	"math"
)

func ColumnarTransposition(str string, columns int, key []int) (string, error) {
	if len(key) != columns {
		return "", errors.New("the key and column length must be the same")
	}

	result := ""
	rows := int(math.Ceil(float64(len(str)) / float64(columns)))

	for row := range rows {
		for _, col := range key {
			char := str[(row*columns)+col]
			result = result + string(char)
		}
	}

	return result, nil
}
