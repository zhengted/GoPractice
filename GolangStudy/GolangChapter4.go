package main

import (
	"fmt"
	"unicode"
)

// 练习 4.3： 重写reverse函数，使用数组指针代替slice。
func reverse(array *[]int) {
	for i,j := 0,len(*array) - 1;i < j ; i,j = i+1,j-1 {
		(*array)[i],(*array)[j] = (*array)[j],(*array)[i]
	}
}

// 练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
func rotate(array []int, nCount uint) []int {
	var res,temp []int
	res = array[:nCount]
	temp = array[nCount:]
	for _,nIns := range res {
		temp = append(temp,nIns)
	}
	return temp
}

// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
func clearSameString(arrStr []string)[]string {
	k := 0
	for _,w := range arrStr {
		if arrStr[k] == w {
			continue
		}
		k++
		arrStr[k] = w
	}
	return arrStr[:k+1]
}

// 练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func clearSameSpace(arrByte []byte) []byte {
	num := len(arrByte)
	for i := 0; i < num; i++ {
		num = len(arrByte)
		if i+1 >= num {
			break
		}
		if unicode.IsSpace(rune(arrByte[i])) && unicode.IsSpace(rune(arrByte[i+1])) {
			copy(arrByte[i:],arrByte[i+1:])
			arrByte = arrByte[:len(arrByte)-1]
			i--
		}
	}
	return arrByte
}

// 练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
func byteArrReserve(byteArr []byte) {
	for i,j := 0,len(byteArr) - 1; i < j; i,j = i+1,j-1 {
		byteArr[i], byteArr[j] = byteArr[j], byteArr[i]
	}
}

func main() {
	arr := &[]int{1,2,3,4,5,6}
	reverse(arr)
	fmt.Printf("After Reverse %v\n",*arr)
	res := rotate(*arr,3)
	fmt.Printf("After Rotate %v\n",res)
	arrStr := []string{"Hello","Hello","Me","World","World","Me","World","A","A"}
	resStr := clearSameString(arrStr)
	fmt.Println(resStr)
	byteArr := []byte{'H',' ',' ','e',' ','l','l',' ',' ',' ','o'}
	resByte := clearSameSpace(byteArr)
	fmt.Println(resByte)
	byteArrReserve(resByte)
	fmt.Println(resByte)
}
