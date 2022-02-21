package main

import "strconv"

func compressString(S string) string {
	newStr := ""
	tmpStr := ""
	num := 0

	if S == "" {
		return S
	}

	for _, s := range S {
		str := string(s)
		if tmpStr == "" {
			newStr += str
			tmpStr = str
		}

		if tmpStr != str {
			newStr += strconv.Itoa(num)
			newStr += str
			tmpStr = str
			num = 0
		}
		num++
	}
	newStr += strconv.Itoa(num)

	if len(S) <= len(newStr) {
		return S
	}

	return newStr
}
