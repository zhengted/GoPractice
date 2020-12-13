package client

import (
	"GoPractice/crawler/engine"
	"GoPractice/crawler_distributed/config"
	"GoPractice/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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

			// call rpc save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("Item Saver: error saving item %s\n", err.Error())
				continue
			}

		}
	}()
	return out, nil
}
