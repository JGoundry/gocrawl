package crawler

import (
	"testing"
)

func BenchmarkCrawler(b *testing.B) {
	c := Crawler{
		BaseUrl:     "https://quotes.toscrape.com",
		WorkerCount: 1000,
	}

	for i := 0; i < b.N; i++ {
		c.Start()
	}
}
