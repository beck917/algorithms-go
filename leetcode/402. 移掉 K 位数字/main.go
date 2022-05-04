package main

import "strings"

func removeKdigits(num string, k int) string {
	var stack []byte

	for k, _ := range num {
		for k > 0 && len(stack) > 0 && num[k] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, num[k])
	}

	res := strings.TrimLeft(string(stack), "0")

	return res
}
