package main

import (
	"fmt"
	"regexp"
)

const text = "My email is zhengted@163.com"
const text2 = `My email is zhengted@163.com
email is abc@def.org
email2 is kkk@qq.com
email3 is kkk@qq.com.cn
`

// 正则表达式
func main() {
	re := regexp.MustCompile(
		`([a-zA-z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text2,-1)
	match := re.FindAllStringSubmatch(text2, -1)
	fmt.Println(match)
}
