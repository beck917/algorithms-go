package main

func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))

	dp[0][0] = 0          //卖出股票时的利润, 是否卖出股票,取决于上次买入股票后的利润+当前股票的价格+手续费>上次卖出股票的利润
	dp[0][1] = -prices[0] //持有股票时的利润,是否买入当前股票.取决于上次卖出股票后的利润-当前股票的价格>上次持有股票的价格

	// 7 1 5 3 6 4
	// 0 0 4 4 7 7
	// -7 -1 -1 1 1 3
	// 可以看到持有(和未持有的)股票的利润,这样计算后,是包含前一次卖出的数据的
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
