package main

import(
	"fmt"
)

func printArray(arr [5]int) {
	arr[0] = 100 //强行修改0下标元素的值
	for i,v := range arr {
		fmt.Println(i,v)
	}
}

func main3_1() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}

	var grid [4][5]int

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	fmt.Println("############Print Array Test############")
	printArray(arr1)
	printArray(arr3)
	fmt.Println("############After Print Array#############")
	fmt.Println(arr1,arr3)
	// 方法调用不会改变数组原有的值  Go中只有值传递 没有引用传递
	// 1.可以选择用指针传递
	// 2.可以选择切片


}
