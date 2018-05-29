package persist

import (
	"context"
	"log"
)
import "github.com/olivere/elastic"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()

	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return "", err
	}

	response, e := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if e != nil {
		return "", err
	}

	return response.Id, nil
}
