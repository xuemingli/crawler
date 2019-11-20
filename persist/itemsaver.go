package persist

import (
	"context"
	"learngo/crawler/engine"
	"log"

	"github.com/pkg/errors"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// TODO: Must turn off sniff if elasticsearch run in docker.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: Got item #%d: %v", itemCount, item)
			itemCount++
			err := save(client, item, index)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("Error: Must supply type.")
	}
	indexService := client.Index().Index(index).Type(item.Type).Id(item.Id).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
