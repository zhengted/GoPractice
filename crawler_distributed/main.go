package main

import (
	"GoPractice/crawler/Scheduler"
	"GoPractice/crawler/engine"
	"GoPractice/crawler/zhenai/parser"
	"GoPractice/crawler_distributed/config"
	itemSaver "GoPractice/crawler_distributed/persist/client"
	"GoPractice/crawler_distributed/rpcsupport"
	worker "GoPractice/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

// ：开头
var (
	itemSaverHost = flag.String(
		"ite,saver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(
		strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &Scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
	//engine.SimpleEngine{}.Run(engine.Request{
	//		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("error connecting to %s: %v", h, err)
		} else {
			clients = append(clients, client)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}

// 城市列表解析器Seed

// 城市解析器
// 输入：UTF-8文本
// 输出：request列表{URL和对应的parser} item列表

// 用户解析器
