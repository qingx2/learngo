package main

import (
	"testing"
	"time"

	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler/model"
	"github.com/guopuke/learngo/crawler_distributed/config"
	"github.com/guopuke/learngo/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")

	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:       "冰之泪",
			Age:        47,
			Height:     170,
			Weight:     189,
			Marriage:   "离异",
			Income:     "8001-12000元",
			Education:  "大专",
			Occupation: "其他职业",
			Gender:     "男",
			House:      "有",
			Car:        "无",
			Hukou:      "无",
			Xingzuo:    "无",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
