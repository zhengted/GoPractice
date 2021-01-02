package main

import (
	"fmt"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	var window []int
	var ret []int
	curIndex := 0
	window = nums[curIndex : curIndex+k]
	lastMax, lastMaxIndex := max239(window, curIndex)
	ret = append(ret, lastMax)
	curIndex++
	for ; curIndex+k-1 < len(nums); curIndex++ {
		window = nums[curIndex : curIndex+k]
		fmt.Println("window", window)
		if lastMaxIndex < curIndex {
			lastMax, lastMaxIndex = max239(window, curIndex)
			fmt.Println("1", lastMax)
			ret = append(ret, lastMax)
		} else if window[len(window)-1] > lastMax {
			lastMax = window[len(window)-1]
			lastMaxIndex = curIndex + k - 1
			fmt.Println("2", lastMax)
			ret = append(ret, lastMax)
		} else {
			fmt.Println("3", lastMax)
			ret = append(ret, lastMax)
		}
		fmt.Println("lastMax", lastMax)
	}
	return ret
}

func max239(window []int, curIndex int) (int, int) {
	ret := -1 * math.MaxInt32
	retIndex := 0
	for i, v := range window {
		if v > ret {
			ret = v
			retIndex = i
		}
	}
	return ret, retIndex + curIndex
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/sliding-window-maximum/solution/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// q队列的变化如下 以 nums = [1 3 1 2 0 5], k = 3 为例
// [1,2]
// [1,3]
// [1,3,4]
// [5]

func maxSlidingWindowEx(nums []int, k int) []int {
	q := []int{} // q队列存储下标
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1] // 和队尾元素比较 如果队尾值较小则先将队尾弹出
		}
		q = append(q, i)
	}

	// 先将前k个元素插入队列中 并且保证插入队列后的元素都是**较大**的
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
