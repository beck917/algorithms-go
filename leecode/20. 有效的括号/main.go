package main

import "fmt"

func main() {
	s := "("
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []byte
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		if _, ok := pairs[s[i]]; !ok {
			// 入栈
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			// 出栈
			tmpv := stack[len(stack)-1]
			if tmpv != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
