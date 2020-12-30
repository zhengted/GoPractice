package main

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		// 1. 取出最大的
		max := 0
		index := 0
		for i := 0; i < len(stones); i++ {
			if stones[i] > max {
				max = stones[i]
				index = i
			}
		}
		stones = append(stones[:index], stones[index+1:]...)
		secMax := 0
		secIndex := 0
		for i := 0; i < len(stones); i++ {
			if stones[i] > secMax {
				secMax = stones[i]
				secIndex = i
			}
		}
		if secMax == max {
			stones = append(stones[:secIndex], stones[secIndex+1:]...)
		} else {
			stones[secIndex] = max - secMax
		}
	}
	if len(stones) < 1 {
		return 0
	}
	return stones[0]
}
