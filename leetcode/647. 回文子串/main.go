package main

import "fmt"

func main() {
	s := "aaaa"
	fmt.Println(countSubstrings(s))
}

/**
输入：s = "aaa" //babad
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
*/
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		tmpone := findSubNum(s, i, i)
		tmptwo := findSubNum(s, i, i+1)

		res += tmpone + tmptwo
	}

	return res
}

func findSubNum(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			res++
			left--
			right++
		} else {
			break
		}
	}
	return res
}
