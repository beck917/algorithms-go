package main

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))

	dp[0] = prices[0]

	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			// 5 > 1
			if prices[i] < dp[i-1] {
				dp[i] = prices[i]
			} else {
				dp[i] = dp[i-1]
			}

			if i == len(prices)-1 {
				max = max + prices[i] - dp[i-1]
			}

		} else {
			// 3 < 5
			//fmt.Println(max,  prices[i-1], dp[i-1],i )
			max = max + prices[i-1] - dp[i-1]
			dp[i] = prices[i]
		}
	}

	return max
}
