package main

import (
	"fmt"
	"math"
)

func main() {
	s := "-000000000000000000000000000000000000000000000000001"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	var tmpNums []int
	numFlag := false
	var marks [2]int
	for _, v := range s {
		if v == ' ' {
			continue
		}

		if len(tmpNums) == 0 {
			if v == '-' {
				numFlag = true
				marks[0] = 1
				continue
			}

			if v == '+' {
				numFlag = false
				marks[1] = 1
				continue
			}

			if v < '0' || v > '9' {
				break
			}
		} else {
			if v < '0' || v > '9' {
				break
			}
		}

		//for i := 0; i < len(tmpNums); i++ {
		//	tmpNums[i] *= 10
		//}

		tmpNums = append(tmpNums, int(v-'0'))
	}

	result := 0

	if marks[0] == 1 && marks[1] == 1 {
		return result
	}

	n := len(tmpNums)
	for i := 0; i < n; i++ {
		num := pow(10, (n - i - 1))

		if result > math.MaxInt64/num {
			result = math.MaxInt32
		} else {
			result += tmpNums[i] * num
		}
	}

	if numFlag == true {
		result *= -1
	}
	fmt.Println(result)

	if result > math.MaxInt32 {
		result = math.MaxInt32
	}

	if result < math.MinInt32 {
		result = math.MinInt32
	}

	return result
}

func pow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}

	return result
}
