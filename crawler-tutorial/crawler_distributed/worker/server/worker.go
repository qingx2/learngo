package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/guopuke/learngo/crawler_distributed/config"
	"github.com/guopuke/learngo/crawler_distributed/rpcsupport"
	"github.com/guopuke/learngo/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}))
}
