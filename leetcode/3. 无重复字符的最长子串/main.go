package main

import "fmt"

/**
这个题目是最长的不重复的连续子串,比如abcabcbb,最长的就是abc.
核心点是需要一个strmap存储已经遍历过的字符,但是需要注意的是,从第二次循环开始,就需要
删除strmap的前一个字符,因为这个时候,从i-1为头的已经遍历完了
第二层循环,判断条件,j<n而且j不在strmap里,这里要注意j只在最外层循环外初始化一次,原因是
abcabcbb 比如第一次循环后,在第二个a这里停止了,也就是j=3的位置,第二次循环从b开始,
bc这来字符,签名已经算过是不重复的了,所以不会在进去循环了,已经在strmap里面了.
这里会继续在a(bca)b也就是j=4的位置停下来,后面也是如此,由此得出j-i就是每次最大的数值了
*/
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
