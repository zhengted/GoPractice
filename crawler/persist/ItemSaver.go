package persist

import (
	"GoPractice/crawler/engine"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver got item "+
				"#%d:%v", itemCount, item)
			itemCount++

			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %s\n", err.Error())
				continue
			}

		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}
	IndexService := client.Index().Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id != "" {
		IndexService.Id(item.Id)
	}

	_, err := IndexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
