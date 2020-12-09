package main

import (
	"GoPractice/crawler/Scheduler"
	"GoPractice/crawler/engine"
	"GoPractice/crawler/persist"
	"GoPractice/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &Scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//engine.SimpleEngine{}.Run(engine.Request{
	//		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	})
}

// 城市列表解析器Seed

// 城市解析器
// 输入：UTF-8文本
// 输出：request列表{URL和对应的parser} item列表

// 用户解析器
