package main

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	m := make(map[byte]int)
	for _, b := range tasks {
		m[b]++
	}
	nMax := 0
	nMaxCount := 0
	for _, n := range m {
		if n > nMax {
			nMax = n
			nMaxCount = 1
		} else if n == nMax {
			nMaxCount++
		}
	}
	res := nMaxCount * nMax * (n - 1)
	return res
}
