package main

import "fmt"

func candy(ratings []int) int {
	if len(ratings) <= 0 {
		return 0
	}
	res := len(ratings)
	left2right := make([]int, len(ratings))
	if ratings[0] > ratings[1] {
		left2right[0]++
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left2right[i] = left2right[i-1] + 1
		}
	}
	fmt.Println(left2right)
	right2left := make([]int, len(ratings))
	if right2left[len(ratings)-1] > right2left[len(ratings)-2] {
		right2left[len(ratings)-1]++
	}
	for j := len(ratings) - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			right2left[j] = right2left[j+1] + 1
		}
	}
	fmt.Println(right2left)
	for k := 0; k < len(ratings); k++ {
		res += max135(left2right[k], right2left[k])
	}
	return res
}

func max135(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
