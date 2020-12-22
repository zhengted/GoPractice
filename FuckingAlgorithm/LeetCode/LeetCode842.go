package main

import "math"

func splitIntoFibonacci(s string) (F []int) {
	n := len(s)
	var backtrack func(index, sum, prev int) bool
	backtrack = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}

		cur := 0
		for i := index; i < n; i++ {
			// 每个块的数字一定不要以零开头，除非这个块是数字 0 本身
			if i > index && s[index] == '0' {
				break
			}

			cur = cur*10 + int(s[i]-'0')
			// 拆出的整数要符合 32 位有符号整数类型
			if cur > math.MaxInt32 {
				break
			}

			// F[i] + F[i+1] = F[i+2]
			if len(F) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}

			// cur 符合要求，加入序列 F
			F = append(F, cur)
			if backtrack(i+1, prev+cur, cur) {
				return true
			}
			F = F[:len(F)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return
}
