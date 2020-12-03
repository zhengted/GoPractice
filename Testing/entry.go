package main

import (
	"fmt"
	"math"
)
var lastOccured = make([]int, 0xffff)
// 最长不重复子串 长度
func lengthOfLongestSubstring(s string) int {

	for i := range lastOccured {
		lastOccured[i] = -1
	}
	start := 0
	maxLength := 0

	for i,ch := range []rune(s) {
		if lastI := lastOccured[ch];lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}

func triangle() {
	var a, b int = 3,4
	fmt.Println(calcTiangle(a,b))
}

func calcTiangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func main() {
	triangle()
}
