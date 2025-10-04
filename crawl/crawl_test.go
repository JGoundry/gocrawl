package crawl

import (
	"testing"
)

func BenchmarkCrawler(b *testing.B) {

	baseUrl := "https://quotes.toscrape.com"
	workerCount := uint(1000)

	for i := 0; i < b.N; i++ {
		Crawl(baseUrl, workerCount)
	}
}
