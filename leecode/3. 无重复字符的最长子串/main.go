package main

import "fmt"

func main() {
	s := "pwwkew" //pwwkew,abcabcbb
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	strMap := make(map[byte]bool)
	maxlen := 0
	n := len(s)
	j := 0

	for i := 0; i < n; i++ {
		// 删除
		if i != 0 {
			delete(strMap, s[i-1])
		}

		// 这里j从-1开始，每次+1对比，是为了预判，下一个字符是否已经存在
		for j < n && strMap[s[j]] != true {
			strMap[s[j]] = true
			j++
		}
		fmt.Println(i, j)

		if j-i > maxlen {
			maxlen = j - i
		}
	}

	return maxlen
}
