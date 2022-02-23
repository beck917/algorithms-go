package main

import "fmt"

func main() {
	fmt.Println(addStrings("0", "10"))
}

/**
输入：num1 = "11", num2 = "123"
输出："134"
*/
func addStrings(num1 string, num2 string) string {
	/**
	var new byte
	for i := 0; i < len(num1); i++ {
		new = num1[i] + num2[i] - '0'

		if new > '9' {
			new -= 10
		}
		fmt.Println(new, '9')
	}
	*/

	n1 := len(num1)
	n2 := len(num2)
	n := n1
	if n2 > n1 {
		n = n2
		for i := 0; i < n2-n1; i++ {
			num1 = "0" + num1
		}
	} else {
		for i := 0; i < n1-n2; i++ {
			num2 = "0" + num2
		}
	}

	var tmp byte
	var str []byte
	for i := n - 1; i >= 0; i-- {
		new := num1[i] + num2[i] - '0' + tmp
		if new > '9' {
			new -= 10
			tmp = 1
		} else {
			tmp = 0
		}
		str = append(str, new)
	}
	if tmp == 1 {
		str = append(str, '1')
	}

	return Reverse(string(str))
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
