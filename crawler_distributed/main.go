package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler/scheduler"
	"github.com/guopuke/learngo/crawler/zhenai/parser"
	"github.com/guopuke/learngo/crawler_distributed/config"
	itemsaver "github.com/guopuke/learngo/crawler_distributed/persist/client"
	"github.com/guopuke/learngo/crawler_distributed/rpcsupport"
	worker "github.com/guopuke/learngo/crawler_distributed/worker/client"
)

func main() {
	port := fmt.Sprintf(":%d", config.ItemSaverPort)
	itemChan, err := itemsaver.ItemSaver(port)
	if err != nil {
		panic(err)
	}

	poll := createClientPoll()

	processor := worker.CreateProcessor(poll)

	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueuedScheduler{},
		WorkerCount:    100,
		ItemChan:       itemChan,
		RequestProcess: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPoll(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
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
