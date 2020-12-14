package main

import (
	"GoPractice/crawler/engine"
	"GoPractice/crawler/model"
	"GoPractice/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

const gHost = ":1234"
const gIndex = "test1"

func TestItemServer(t *testing.T) {
	// Start ItemSaverServer
	go serveRpc(gHost, gIndex)
	time.Sleep(time.Second)
	// Start ItemSaverClient
	client, err := rpcsupport.NewClient(gHost)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		"http://localhost:8080/mock/album.zhenai.com/u/7143522202848495805",
		"zhenai",
		"7143522202848495805",
		model.Profile{
			Name:       "厌与深情记得笑i",
			Gender:     "女",
			Age:        10,
			Height:     1,
			Weight:     116,
			Income:     "2001-3000元",
			Marriage:   "未婚",
			Education:  "硕士",
			Occupation: "销售",
			Hokou:      "东莞市",
			Xinzuo:     "天秤座",
			House:      "有房",
			Car:        "有豪车",
		},
	}
	// Call save
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
