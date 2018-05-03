package main

import (
	"fmt"
	"github.com/guopuke/learngo/retriever/mock"
	"github.com/guopuke/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"Hello retriever"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
