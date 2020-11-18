package main

import (
	"fmt"
	"sort"
)

func main() {
	// 一、map定义法
	// 1.
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	// 2.
	ages1:= map[string]int{
		"alice":31,
		"charlie":34,
		"charli":34,
		"charl":34,
		"char":34,
		"cha":34,
		"ch":34,
	}
	// 3. 这种情况要注意  由于ages2为nil值 是一个空的map所以直接赋值的操作会panic
	//	查找 删除 range都没问题
	// var ages2 map[string]int
	// ages2["alice"] = 20

	// 二、可以对不存在的元素加值
	ages1["dad"]+=1

	// 三、遍历
	for name,age := range ages1 {
		fmt.Printf("%s\t%d\n",name,age)
	}
	/*
	char	34
	cha	34
	ch	34
	dad	1
	alice	31
	charlie	34
	charli	34
	charl	34
	*/
	/*
		上面打印可以看出遍历的顺序是随机的 如果需要按序要引入sort包
		以下是按字母序排列的结果
	*/
	var names []string
	for name := range ages1 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages1[name])
	}


}
// 四、图
var graph = make(map[string]map[string]bool)
func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

// 五、习题
// 练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
//
//练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。