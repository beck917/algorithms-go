package main

import "fmt"

func main() {
	prices := []int{7, 2, 5, 3, 10, 1, 7}

	fmt.Println(maxProfit((prices)))
}

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = prices[0]

	max := 0
	for i := 1; i < len(prices); i++ {

		if prices[i]-dp[i-1] > max {
			max = prices[i] - dp[i-1]
		}

		if prices[i] <= dp[i-1] {
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return max
}
