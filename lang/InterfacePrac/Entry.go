package main

import (
	"GoPractice/lang/InterfacePrac/ByteCounter"
	"GoPractice/lang/InterfacePrac/LineCounter"
	"GoPractice/lang/InterfacePrac/Tree"
	"GoPractice/lang/InterfacePrac/WordCounter"
	"fmt"
	"io"
	"log"
)

type CountWriter struct {
	Writer io.Writer
	Count  int64
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	n, err := cw.Writer.Write(p)
	if err != nil {
		log.Printf("An error occured:%s", err.Error())
		return 0, err
	}
	cw.Count += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer, strType string) (io.Writer, *int64) {
	cw := CountWriter{
		Writer: w,
	}
	return &cw, &(cw.Count)
}

func testByteCount() {

	var (
		c ByteCounter.ByteCounter
		b LineCounter.LCount
		w WordCounter.Count
	)
	// Golang 圣经 习题7.1
	fmt.Println("Ch.7 7.1")
	// Byte Counter
	fmt.Println("Byte counter test")
	c.Write([]byte("Hello"))
	fmt.Println(c)
	c = 0
	name := "Dolly"
	// 这里Fprintf的第一个参数是指针接收者(是指针作为this重载io.Writer)
	// 因此传参的时候要写取地址
	fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(c)

	// Line Counter
	fmt.Println("Line counter test")
	b.Write([]byte("I am the king\nShe is my queen"))
	fmt.Println(b)

	// Word Counter
	fmt.Println("Word counter test")
	w.Write([]byte("I am the king"))
	fmt.Println(w)

	// 习题7.2  没做出来 上面是别人的答案

}

func BuildTree() *Tree.Node {
	root := Tree.Node{Val: 1}
	root.Left = &Tree.Node{2, nil, nil}
	root.Right = &Tree.Node{3, nil, nil}
	root.Left.Right = &Tree.Node{4, nil, nil}
	root.Right.Left = &Tree.Node{5, nil, nil}
	return &root
}

func main() {
	//testByteCount()
	var rw io.ReadWriter

	fmt.Println("WriteData")
	rw.Write([]byte("hello"))
	fmt.Println("Read Data")
	var temp []byte
	rw.Read(temp)
	fmt.Println("result:", string(temp))
}
