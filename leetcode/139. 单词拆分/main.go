package main

/**
这个题目很绕,dp+hash
注意dp[i],正好是字符串的单词的起点

比如leetcode
遍历到leet中的t的时候,也就是s[0:4),取dp[j]==dp[0],然后复制dp[4]=true
然后遍历到code的时候,s[4:8)也就是dp[j]==dp[4],正好取到true
*/
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, 0)

	for _, word := range wordDict {
		wordMap[word] = true
	}

	n := len(s)

	dp := make([]bool, n+1)

	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[n]
}
