package main

import (
	"MoocPractice/retriever/mock"
	"MoocPractice/retriever/real"
	"fmt"
	"time"
)
const URL = "https://imooc.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster)  {
	poster.Post("https://www.imooc.com",
		map[string]string{
		"name":"ccmouse",
		"course":"golang",
		},
	)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(URL, map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(URL)
}

func main() {
	var r Retriever
	r = mock.Retriever{"This is a fake url"}
	inspect(r)
	r = &real.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}
	inspect(r)

	// Type Assertion
	if mockRetriever,ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println("Not a mock Retriever")
	}

	// 接口的组合 参考io.ReadWriter
	retriever := mock.Retriever{Contents: "This is a mock Retrieber"}
	fmt.Println("Try a Session")
	fmt.Println(session(&retriever))

}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("This is a mock Retriever",v.Contents)
	case *real.Retriever:
		fmt.Println("Real retriever",v.UserAgent)
	}
}
