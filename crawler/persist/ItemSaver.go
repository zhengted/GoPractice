package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver got item "+
				"#%d:%v", itemCount, item)
			itemCount++

			id, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v %s\n", id, err.Error())
				continue
			}

		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", nil
	}
	return resp.Id, nil
}
