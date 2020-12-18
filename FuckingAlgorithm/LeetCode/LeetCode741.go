package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	var dp [][]int
	for i := 0; i < len(prices); i++ {
		temp := make([]int, len(prices))
		dp = append(dp, temp)
	}
	for sell := 0; sell < len(prices); sell++ {
		for buy := 0; buy < len(prices); buy++ {
			if sell > buy {
				dp[sell][buy] = prices[sell] - prices[buy] - fee
			}
		}
	}
	//PrindTwoDSlice(dp)
	ret := 0
	for i := 1; i < len(prices); i++ {
		nUpTriangleMax := FindMaxInTriangle(0, i, dp)
		nDownTriangleMax := FindMaxInTriangle(i+1, len(prices)-1, dp)
		curVal := nUpTriangleMax + nDownTriangleMax
		if curVal > ret {
			fmt.Println(i, nUpTriangleMax, nDownTriangleMax)
			ret = curVal
		}
	}
	return ret
}

// 在三角形中找出最大值
func FindMaxInTriangle(start int, end int, dp [][]int) int {
	if start == end {
		return dp[start][end]
	}
	if start > end {
		return 0
	}
	ret := 0
	for s := start; s <= end; s++ {
		for b := start; b <= end; b++ {
			if s > b && dp[s][b] > ret {
				ret = dp[s][b]
			}
		}
	}
	if ret < 0 {
		ret = 0
	}
	return ret
}

func PrindTwoDSlice(dp [][]int) {
	for _, s := range dp {
		for _, val := range s {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

// TODO 动态规划
// sell 表示第i天交易完后手里没有股票的最大利润
// buy  表示第i天交易完后手里有股票
//		sell的状态转移：①前一天也没有股票（sell不改变） ②卖出手里持有的股票（buy+prices[i]-fee）
//		buy 的状态转移：①前一天已经有股票了（buy不改变）②买下当前的股票（sell-prices[i]）
func maxProfitEx(prices []int, fee int) int {
	n := len(prices)
	sell, buy := 0, -prices[0]
	for i := 1; i < n; i++ {
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, sell-prices[i])
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
