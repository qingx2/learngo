package client

import (
	"log"

	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, e := rpcsupport.NewClient(host)
	if e != nil {
		return nil, e
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			e := client.Call("ItemSaverService.Save", item, &result)
			if e != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, e)
			}
		}
	}()

	return out, nil
}
