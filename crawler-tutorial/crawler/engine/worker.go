package engine

import "log"
import "github.com/guopuke/learngo/crawler/fetcher"

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: err fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parser(body, r.Url), nil
}
