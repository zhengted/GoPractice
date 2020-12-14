package persist

import (
	"GoPractice/crawler/engine"
	"GoPractice/crawler/persist"
	"github.com/olivere/elastic/v7"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (i *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(i.Client, item, i.Index)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v %v", item, err)
	}
	return err
}
