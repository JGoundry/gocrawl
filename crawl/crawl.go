package crawl

import (
	"fmt"
	"gocrawl/debug"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Word -> URLs -> Occurence count
type InvertedIndex = map[string]map[string]uint

func Crawl(baseUrl string, workerCount uint) (InvertedIndex, []string) {
	c := crawler{
		BaseUrl:     "https://quotes.toscrape.com",
		WorkerCount: workerCount,
	}
	c.run()

	return c.index, c.urlsVisited()

}

type crawler struct {
	BaseUrl     string
	WorkerCount uint

	mu          sync.Mutex          // Mutex for visited set, url overflow queue
	visited     map[string]struct{} // Visited URL set
	index       InvertedIndex       //
	urlQueue    chan string         //
	urlOverflow []string            // URL queue to prevent blocking when chan full
	workTracker sync.WaitGroup      // Blocks on main thread until all work is done
}

func (c *crawler) crawlWorker() {
	for {
		select {
		case url, ok := <-c.urlQueue:
			if !ok {
				debug.Println("Worker exiting: URL queue closed")
				return
			}
			debug.Println("Crawling url from chan", url)
			c.crawl(url)

		default:
			c.mu.Lock()
			if len(c.urlOverflow) == 0 {
				c.mu.Unlock()
			} else {
				url := c.urlOverflow[0]
				c.urlOverflow = c.urlOverflow[1:]
				c.mu.Unlock()
				debug.Println("Crawling url from overflow", url)
				c.crawl(url)
			}
		}
	}
}

func (c *crawler) crawl(url string) {
	defer c.workTracker.Done()

	c.mu.Lock()
	if _, visited := c.visited[url]; visited {
		c.mu.Unlock()
		return // already visited
	}
	c.visited[url] = struct{}{}
	c.mu.Unlock()

	node, err := getHtml(url)
	if err != nil {
		debug.Println(err)
		return
	}

	// A recursive function (or a simple visitor pattern) to walk the HTML nodes
	var findLinks func(*html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// search an.Attr child nodes? already in children?
			urls := findUrls(n.Attr, c.BaseUrl)

			c.workTracker.Add(len(urls)) // CRITICAL: increment work tracker first otherwise data race

			for _, url := range urls {
				select {
				case c.urlQueue <- url:
				default:
					debug.Println("Warning: URL queue is full, adding to overflow")
					c.mu.Lock()
					c.urlOverflow = append(c.urlOverflow, url)
					c.mu.Unlock()
				}
			}
		} else if n.Type == html.TextNode {
			words := strings.Fields(n.Data)
			c.mu.Lock()
			for _, word := range words {
				if _, ok := c.index[word]; !ok {
					c.index[word] = make(map[string]uint)
				}
				c.index[word][url]++
			}
			c.mu.Unlock()
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLinks(c)
		}
	}

	findLinks(node)
}

func (c *crawler) run() {
	var wg sync.WaitGroup

	c.visited = make(map[string]struct{})
	c.urlQueue = make(chan string, 1000)
	c.urlOverflow = make([]string, 0, 1000)
	c.index = make(InvertedIndex)

	// Increment work tracker and send base url to chan
	c.workTracker.Add(1)
	c.urlQueue <- c.BaseUrl

	workerCount := max(c.WorkerCount, 1)
	wg.Add(int(workerCount))
	for range workerCount {
		go func() {
			defer wg.Done()
			c.crawlWorker()
		}()
	}

	// Wait until all work is done
	c.workTracker.Wait()

	// Send done signal to workers
	close(c.urlQueue)

	// Wait for workers to complete
	wg.Wait()
}

func (c *crawler) urlsVisited() []string {
	urlsVisited := make([]string, 0, len(c.visited))
	for key := range c.visited {
		urlsVisited = append(urlsVisited, key)
	}
	return urlsVisited
}

func getHtml(url string) (*html.Node, error) {
	resp, err := http.Get(url) // send GET
	if err != nil {
		return nil, err
	}

	// Check status is ok
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-OK HTTP Status: %v", resp.Status)
	}

	// Ensure content is html
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return nil, fmt.Errorf("Content-Type is not html")
	}

	node, err := html.Parse(resp.Body) // parse HTML
	resp.Body.Close()                  // close http reader
	if err != nil {
		return nil, err
	}
	return node, nil
}

func findUrls(attr []html.Attribute, baseUrl string) []string {
	var urls []string
	for _, a := range attr {
		if a.Key == "href" {
			if len(a.Val) == 0 || a.Val[0] != '/' { // skip other websites for now
				continue
			}

			urls = append(urls, fmt.Sprintf("%s%s", baseUrl, a.Val))
		}
	}
	return urls
}
