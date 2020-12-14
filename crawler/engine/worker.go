package engine

import (
	"GoPractice/crawler/fetcher"
	"log"
)

// TODO
// use redis to abandon duplications
func Worker(r Request) (ParseResult, error) {
	//log.Printf("worker Fetching %v",r)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v",
			r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
