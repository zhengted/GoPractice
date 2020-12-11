package persist

import (
	"GoPractice/crawler/engine"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver got item "+
				"#%d:%v", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %s\n", err.Error())
				continue
			}

		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("must supply type")
	}
	IndexService := client.Index().Index("dating_profile").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id != "" {
		IndexService.Id(item.Id)
	}

	_, err = IndexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
