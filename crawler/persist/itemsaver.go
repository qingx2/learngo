package persist

import (
	"context"
	"log"

	"github.com/guopuke/learngo/crawler/engine"
	"github.com/pkg/errors"
)
import "github.com/olivere/elastic"

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			e := Save(client, index, item)
			if e != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, e)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, e := indexService.Do(context.Background())
	if e != nil {
		return e
	}

	return nil
}
