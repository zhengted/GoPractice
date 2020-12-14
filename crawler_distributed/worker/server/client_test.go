package main

import (
	"GoPractice/crawler_distributed/config"
	"GoPractice/crawler_distributed/rpcsupport"
	"GoPractice/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(5 * time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	request := worker.Request{
		"http://localhost:8080/mock/album.zhenai.com/u/7143522202848495805",
		worker.SerializedParser{
			config.ParseProfile,
			"厌与深情记得笑i",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, request, &result)
	if err != nil {
		t.Errorf(err.Error())
	}
}
