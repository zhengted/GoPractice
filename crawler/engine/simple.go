package engine

import (
	"GoPractice/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Run Fetching %s", r.Url)
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,
			parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got Item %v\n", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	//log.Printf("worker Fetching %v",r)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v",
			r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
