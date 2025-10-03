package main

/*

    TODO List

	[ ] Scrape web pages and create inverted index
	[ ] Crawler.SaveIndex() JSON to disk
	[ ] Crawler.LoadIndex() JSON from disk
	[ ] CLI to run, save, load, print

*/

import (
	"fmt"
	"webcrawler/crawler"
)

func main() {
	c := crawler.Crawler{
		BaseUrl:     "https://quotes.toscrape.com",
		WorkerCount: 1000,
	}

	c.Start()

	fmt.Println("Total visited:", c.TotalVisited())

}
