package main

import "fmt"

// 题目：寻找最长不含有重复字符的子串
/*

	解法：对于每一个字母x
		1. lastOccured[x] 不存在或者 < start ===》无需操作
		2. lastOccured[x] >=start ===> 更新start
		3. 更新lastOccured
*/

func lengthOfNonRepeatSubStr(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI,ok := lastOccured[ch]; ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatSubStr("abcabcaa"))
	fmt.Println(lengthOfNonRepeatSubStr("abcabdcfaaa"))
	fmt.Println(lengthOfNonRepeatSubStr("aaaaaa"))
	fmt.Println(lengthOfNonRepeatSubStr(""))
	fmt.Println(lengthOfNonRepeatSubStr("一二三二一"))
}
