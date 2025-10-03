package main

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
