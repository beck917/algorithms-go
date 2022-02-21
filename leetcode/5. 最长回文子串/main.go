package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	longestStr := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)

		if len(s1) > len(longestStr) || len(s2) > len(longestStr) {
			if len(s1) >= len(s2) {
				longestStr = s1
			} else {
				longestStr = s2
			}
		}
	}
	return longestStr
}

func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}

	return s[left+1 : right]
}
