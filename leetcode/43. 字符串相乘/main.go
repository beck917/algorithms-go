package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(multiply("9", "9"))
}

func multiply(num1 string, num2 string) string {
	sum := 0
	for i := len(num2) - 1; i >= 0; i-- {
		tmp := 0
		carry := 0
		multi := 1
		for j := len(num1) - 1; j >= 0; j-- {
			if j == len(num1)-1 {
				multi = 1
			} else {
				multi = multi * 10
			}
			num := int(num2[i]-'0')*int(num1[j]-'0') + carry
			tmp += num % 10 * multi

			carry = num / 10
		}

		if carry != 0 {
			tmp += carry * multi * 10
		}

		if i == len(num2)-1 {
			sum += tmp
		} else {
			sum += (tmp) * int(math.Pow(float64(10), float64(len(num2)-i-1)))
		}
	}

	return strconv.Itoa(sum)
}
