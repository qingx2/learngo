package main

import (
	"log"

	"github.com/guopuke/learngo/crawler_distributed/persist"
	"github.com/guopuke/learngo/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic"
)

func main() {
	log.Fatal(serveRpc(":1234", "crawler_dating_profile"))
	// client, err := elastic.NewClient(elastic.SetSniff(false))
	// if err != nil {
	// 	panic(err)
	// }
	// rpcsupport.ServeRpc(":1234", persist.ItemSaverService{
	// 	Client: client,
	// 	Index:  "crawler_dating_profile",
	// })
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
