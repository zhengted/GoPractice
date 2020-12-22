package main

import "math"

//硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
//
//示例1:
//
//输入: n = 5
//输出：2
//解释: 有两种方式可以凑成总金额:
//5=5
//5=1+1+1+1+1
//示例2:
//
//输入: n = 10
//输出：4
//解释: 有四种方式可以凑成总金额:
//10=10
//10=5+5
//10=5+1+1+1+1+1
//10=1+1+1+1+1+1+1+1+1+1

func waysToChange(n int) int {
	if n == 0 {
		return 0
	}
	preTwenty := []int{1, 2, 4}
	nLoc := int(math.Floor(float64(n / 5)))
	if n < 15 {
		return preTwenty[nLoc]
	}
	initialNum := 6
	nLoc = nLoc - 3
	j := 3
	for i := 0; i < nLoc; i++ {

		initialNum = int(math.Mod(float64(initialNum), 1000000007) + math.Mod(float64(j), 1000000007))
		j++
	}
	return initialNum
}
