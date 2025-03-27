package leetcode

import (
	"math"
	"strings"
)

// #################### string to integer atoi ####################
func MyAtoi(s string) int {

	trimStr := strings.Trim(s, " ")
	length := len(trimStr)

	if length == 0 {
		return 0
	}

	var checkNavigates bool = string(trimStr[0]) == "-"
	var checkPositive bool = string(trimStr[0]) == "+"

	var numStr string
	for i := 0; i < length; i++ {
		if checkNavigates && i == 0 {
			continue
		}

		if checkPositive && i == 0 {
			continue
		}

		if trimStr[i] == '0' && len(numStr) == 0 {
			continue
		}

		if trimStr[i] >= '0' && trimStr[i] <= '9' {
			numStr += string(trimStr[i])
		} else if len(numStr) > 0 {
			break
		} else {
			return 0
		}
	}

	var pow = math.Pow10(len(numStr) - 1)

	var maxInt = math.Pow(2, 31)

	if pow >= maxInt {
		if checkNavigates {
			return int(maxInt) * -1
		}
		return int(maxInt) - 1
	}

	var result int
	var index = int(pow)
	for i := 0; i < len(numStr); i++ {

		var num int
		switch numStr[i] {
		case '1':
			num = 1
		case '2':
			num = 2
		case '3':
			num = 3
		case '4':
			num = 4
		case '5':
			num = 5
		case '6':
			num = 6
		case '7':
			num = 7
		case '8':
			num = 8
		case '9':
			num = 9
		default:
			num = 0
		}
		result += num * index

		index /= 10
	}

	if checkNavigates {
		result *= -1

		if result <= int(maxInt*-1) {
			return int(maxInt) * -1
		}
	}

	if result >= int(maxInt) {
		return int(maxInt) - 1
	}

	return result

}
