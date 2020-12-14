package main

import "fmt"

func main() {
	groupAnagramsDemo()
}

func groupAnagramsDemo() {
	param := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsEx(param))
}
