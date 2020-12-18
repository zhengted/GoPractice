package main

import "fmt"

func main() {
	price := []int{1, 4, 6, 2, 8, 3, 10, 14}
	fee := 3
	fmt.Println(maxProfit(price, fee))

}

func groupAnagramsDemo() {
	param := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsEx(param))
}
