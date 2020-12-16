package main

import "fmt"

func main() {
	var x int
	inc := func() int {
		x++
		return x
	}
	fmt.Println(func() (a, b int) {
		return inc(), inc()
	}())
}

func groupAnagramsDemo() {
	param := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsEx(param))
}
