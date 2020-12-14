package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	var mTotalMapSlice []map[byte]int // 记录每个单词的map
	var ret [][]string                // 结果切片
	if len(strs) <= 0 {
		return ret
	}
out:
	for _, str := range strs {
		curMap := WordToMap(str)
		if len(mTotalMapSlice) > 0 {
			for i, tempMap := range mTotalMapSlice {
				if equal(curMap, tempMap) {
					ret[i] = append(ret[i], str)
					continue out
				}
			}
			mTotalMapSlice = append(mTotalMapSlice, curMap)
			newSliceString := []string{str}
			ret = append(ret, newSliceString)
			continue
		} else {
			mTotalMapSlice = append(mTotalMapSlice, curMap)
			newSliceString := []string{str}
			ret = append(ret, newSliceString)
			continue
		}
	}
	return ret
}

// 将单词转换为map
func WordToMap(str string) map[byte]int {
	ret := make(map[byte]int)
	for _, b := range str {
		ret[byte(b)]++
	}
	return ret
}

// 判断两个Map是否相同
func equal(a map[byte]int, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// 化劲儿
func groupAnagramsEx(strs []string) [][]string {
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	tempMap := make(map[int]int)
	var ret [][]string
	for _, s := range strs {
		key := Keyize(s, &prime)
		fmt.Println(s, key)
		if _, ok := tempMap[key]; ok {
			for i, val := range ret {
				if Keyize(val[0], &prime) == key {
					ret[i] = append(ret[i], s)
				}
			}
		} else {
			tempMap[key]++
			newSliceString := []string{s}
			ret = append(ret, newSliceString)
		}
	}
	return ret
}

func Keyize(s string, prime *[]int) int {
	ret := 1
	for _, b := range s {
		ret = ret * (*prime)[b-97]
	}
	return ret
}
