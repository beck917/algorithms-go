package main

func reverseWords(s string) string {
	b := []byte(s)
	prev := 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			reverse(b, prev, i-1)
			prev = i + 1
		}
	}

	reverse(b, prev, len(s)-1)

	return string(b)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
