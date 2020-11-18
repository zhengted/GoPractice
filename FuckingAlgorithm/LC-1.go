package main

import (
	"fmt"
)

func twoSum(nums []int,target int) []int {
	var i,j int
	i,j = 0,len(nums)-1
	for ;nums[i]+nums[j] != target; {
		if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	var res []int
	res = append(res,i,j)
	return res
}

func main(){
	arr := []int{2,7,11,15}
	target := 9
	res := twoSum(arr,target)
	fmt.Println(res)
}
