package main

/*

    TODO List

	[ ] Create inverted index on crawler
	[ ] DataStore.SaveIndex() JSON to disk
	[ ] DataStore.LoadIndex() JSON from disk
	[ ] ReportIndex()

*/

import (
	"fmt"
	"gocrawl/crawl"
	"gocrawl/datastore"
	"gocrawl/report"
	"strings"
	"time"
)

type Command = string

const (
	HelpCommand   Command = "help"
	StatusCommand Command = "status"
	RunCommand    Command = "run"
	LoadCommand   Command = "load"
	SaveCommand   Command = "save"
	UrlsCommand   Command = "urls"
	IndexCommand  Command = "index"
	ExitCommand   Command = "exit"
)

func header() string {
	return "Gocrawl @ Josh Goundry"
}

func helpPrompt() string {
	return "Use command `help` for usage"
}

func help() string {
	var sb strings.Builder
	sb.WriteString("Gocrawl is a webcrawler and inverted index reporter\n")
	sb.WriteString("\n")
	sb.WriteString("Commands:\n")
	sb.WriteString("    " + HelpCommand + "\n")
	sb.WriteString("    " + StatusCommand + "\n")
	sb.WriteString("    " + RunCommand + "\n")
	sb.WriteString("    " + LoadCommand + "\n")
	sb.WriteString("    " + SaveCommand + "\n")
	sb.WriteString("    " + UrlsCommand + "\n")
	sb.WriteString("    " + IndexCommand + "\n")
	sb.WriteString("    " + ExitCommand + "\n")
	sb.WriteString("\n")
	sb.WriteString("Load data, or run crawler to display index information\n")
	return sb.String()
}

func getInput() string {
	fmt.Printf("$ ")
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	ds := datastore.NewDataStore()

	baseUrl := "https://quotes.toscrape.com"

	// Print
	fmt.Println(header())
	fmt.Println(helpPrompt())

	for input := getInput(); input != ExitCommand; input = getInput() {
		switch input {
		case HelpCommand:
			fmt.Println(help())

		case StatusCommand:
			fmt.Println(ds.State())

		case RunCommand:
			start := time.Now()
			index, urlsVisited := crawl.Crawl(baseUrl, 1000)
			fmt.Printf("Crawled %q in %v\n", baseUrl, time.Since(start))
			ds.Load(index, urlsVisited)

		case LoadCommand:
			fmt.Println("<TODO>")

		case SaveCommand:
			fmt.Println("<TODO>")

		case IndexCommand:
			index, err := ds.Index()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			report.ReportIndex(index)

		case UrlsCommand:
			urls, err := ds.UrlsVisited()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			report.ReportVisitedUrls(urls)

		case ExitCommand:
			return

		default:
			fmt.Printf("Invalid command: %q\n", input)
			fmt.Println(helpPrompt())
		}
	}
}
