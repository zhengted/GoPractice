package main

import (
	"GoPractice/crawler_distributed/config"
	"GoPractice/crawler_distributed/persist"
	"GoPractice/crawler_distributed/rpcsupport"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
	// 出错强制退出
}
func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})

}
