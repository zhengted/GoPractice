package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v\tlen:%d\tcap:%d\t\n",s,len(s),cap(s))
}

func main() {
	fmt.Println("Define slice\n")
	// method 1
	a := [6]int{1,2,3,4,5,6}
	s0 := a[:3]
	printSlice(s0)

	// method 2
	var s1 []int
	printSlice(s1)
	for i := 1;i < 10;i++ {
		s1 = append(s1,2*i)
	}
	printSlice(s1)

	// method 3
	s2 := make([]int,10)
	printSlice(s2)
	s3 := make([]int,8,16)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlice(s2)
	// 注意S3复制S1之后 s3的len发生了改变
	copy(s3,s1)
	printSlice(s3)

	fmt.Println("Delete element from slice")
	s2 = append(s2[:5],s2[6:]...)
	printSlice(s2)
}
