package main

import "fmt"

func main() {
	s := []byte{1}
	reverseString(s)
	fmt.Println(s)
}

func reverseString2(s []byte) {
	slen := len(s)
	j := slen - 1
	for i := 0; i < slen/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func reverseString(s []byte) {
	slen := len(s)
	right := slen - 1
	left := 0
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
