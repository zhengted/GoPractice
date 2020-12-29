package main

import "fmt"

func minPatches(nums []int, n int) (patches int) {
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			patches++
		}
		fmt.Printf("x = %d\t patches = %d\t nums = %v i = %d\n", x, patches, nums, i)
	}
	return
}
