package main

import (
	"GoPractice/CatchError/fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	//defer fmt.Println(1)		// return 之前执行
	//defer fmt.Println(2)		// 栈打印 先打印2 再打印1
	//fmt.Println(3)

	// 语句在运行时计算
	for i:=0; i < 100; i++ {
		defer fmt.Println(i) // 输出 30 29 28 27 26
		if i == 30 {
			panic("printed too much")
		}
	}
}

func writeFile(filename string) {
	// file,err := os.Create(filename)

	//8-2 错误处理
	file,err := os.OpenFile(
		filename, os.O_EXCL | os.O_CREATE,0666)
	// err = errors.New("this is a custom error")	//自建error
	if err != nil {
		//fmt.Println("Error:",err)
		if pathError, ok := err.(*os.PathError);!ok {
			panic(err)
		}else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
		// panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fib()
	for i := 0;i < 20; i++ {
		fmt.Fprintln(writer,f())
	}
}

func main() {
	writeFile("fib.txt")
	//tryDefer()
}
