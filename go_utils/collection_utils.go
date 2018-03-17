package go_utils

import (
	"strconv"
)

func MapStringToInt(input []string) []int64 {
	var output = make([]int64, len(input))
	for index, value := range input {
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			output[index] = parsed
		}
	}

	return output
}

func MaxValue(input []int64) int64 {
	var max int64 = 0
	for _, value := range input {
		if value > max {
			max = value
		}
	}

	return max
}