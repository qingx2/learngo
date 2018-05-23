package main

import (
	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
