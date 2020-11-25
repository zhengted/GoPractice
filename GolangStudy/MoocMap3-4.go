package main

import "fmt"

func main() {
	var m1 = map[string]string{
		"king":"queen",
		"knight":"princess",
	}

	m2 := make(map[string]int)		// m2 == empty map
	var m3 map[string]int			// m3 == nil(Go中的nil是安全的)

	fmt.Println(m1,m2,m3)
	fmt.Println("Tracesing map")
	// 遍历时无序的
	for k,v := range m1 {
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	knightName := m1["knight"]
	fmt.Println(knightName)
	knightName1 := m1["kn"]
	fmt.Println(knightName1) 	// 空串
	// 判空操作
	if kingName, ok := m1["king1"]; ok {
		fmt.Println(kingName)
	}else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Delete element")
	delete(m1,"king")
	fmt.Println(m1)
	delete(m1,"king")	// 不会panic

	/*
		map注意： key -- 使用哈希表，必须可以比较相等  除了slice map function都能作为key
			Struct也可作为key
	*/
}
