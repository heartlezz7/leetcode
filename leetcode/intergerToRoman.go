package leetcode

import (
	"math"
	"strings"
)

// ############### integer to roman ###############
func sliceNumber(n int) []int {
	var result []int
	for n > 0 {
		digit := n % 10                          // Get the last digit
		result = append([]int{digit}, result...) // Prepend to maintain order
		n /= 10                                  // Remove the last digit
	}
	return result
}

func intToRoman(num int) string {
	var result string

	numSlice := sliceNumber(num)
	digit := int(math.Pow10(len(numSlice) - 1))
	for _, numValue := range numSlice {
		if digit >= 1000 {
			result += strings.Repeat("M", numValue)
		} else if digit >= 100 {
			if numValue == 9 {
				result += "CM"
			} else if numValue >= 5 {
				result += "D" + strings.Repeat("C", numValue-5)
			} else if numValue == 4 {
				result += "CD"
			} else {
				result += strings.Repeat("C", numValue)
			}
		} else if digit >= 10 {
			if numValue == 9 {
				result += "XC"
			} else if numValue >= 5 {
				result += "L" + strings.Repeat("X", numValue-5)
			} else if numValue == 4 {
				result += "XL"
			} else {
				result += strings.Repeat("X", numValue)
			}
		} else {
			if numValue == 9 {
				result += "IX"
			} else if numValue >= 5 {

				result += "V" + strings.Repeat("I", numValue-5)

			} else if numValue == 4 {
				result += "IV"
			} else {
				result += strings.Repeat("I", numValue)
			}
		}
		digit /= 10

	}

	return result
}
