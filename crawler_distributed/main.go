package main

import (
	"fmt"

	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler/scheduler"
	"github.com/guopuke/learngo/crawler/zhenai/parser"
	"github.com/guopuke/learngo/crawler_distributed/config"
	"github.com/guopuke/learngo/crawler_distributed/persist/client"
)

func main() {
	port := fmt.Sprintf(":%d", config.ItemSaverPort)
	itemChan, err := client.ItemSaver(port)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
