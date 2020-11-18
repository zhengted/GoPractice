package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := arr[:]
	fmt.Println("After Update slice s1")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("After Update slice s2")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
	/*
		切片可以改变数组的值
	*/


	s3 := s1[3:5]	// [5,6]  按理来说取得是s1[3] s1[4] 但是硬取是取不出来的
	// 理解 Array的View视图
	fmt.Printf("s1=%v\tlen(s1)=%d\tcap(s1)=%d\n",s1,len(s1),cap(s1))
	fmt.Printf("s3=%v\tlen(s3)=%d\tcap(s3)=%d\n",s3,len(s3),cap(s3))
	//s1=[100 3 4 5]	len(s1)=4	cap(s1)=6
	//s3=[5 6]	len(s3)=2	cap(s3)=3
	// cap的大小是基于原数组的

}
