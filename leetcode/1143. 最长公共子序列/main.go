package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}

// 输入：text1 = "abcde", text2 = "ace" z
// "bsbininm" "jmjkbkjkv"
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))

	tmpi := len(text1)
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
		if text1[i] == text2[0] || i > tmpi {
			dp[i][0] = 1
			tmpi = i
		}
	}

	tmpi = len(text2)
	for i := 0; i < len(text2); i++ {

		if text1[0] == text2[i] || i > tmpi {
			dp[0][i] = 1
			tmpi = i
		}
	}

	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			tmp := dp[i-1][j]
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				fmt.Println(i, j, dp[i][j], string(text1[i]), string(text2[j]))
			} else {
				if dp[i][j-1] > tmp {
					tmp = dp[i][j-1]
				}
				dp[i][j] = tmp
			}
			//fmt.Println(i, j, dp[i][j])
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}
