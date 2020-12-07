package main

import "fmt"

func main() {
	input := []string{
		"A", "A", "A", "A", "A", "A", "B", "C", "D", "E", "F", "G", "H", "B", "B",
	}
	dist := 2
	fmt.Println(leastIntervalEx(input, dist))
}

func leastIntervalEx(tasks []string, n int) int {
	if n == 0 {
		return len(tasks)
	}
	m := make(map[string]int)
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

	res := nMaxCount*(nMax-1)*(n+1) + 1
	return res
}
