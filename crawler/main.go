package main

import (
	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler/scheduler"
	"github.com/guopuke/learngo/crawler/zhenai/parser"
)

func main() {
	// 1. Single Task Edition
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// 2. Concurrent Edition
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
