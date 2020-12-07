package parser

import (
	"GoPractice/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User"+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParserProfile(bytes, name)
				}, // 函数式编程
			})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
