package main

import (
	"bufio"
	"fmt"
	"gocrawl/crawl"
	"gocrawl/datastore"
	"gocrawl/report"
	"os"
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
	sb.WriteString("Load data, or run crawler to populate datastore\n")
	return sb.String()
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("$ ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	ds := datastore.NewDataStore()

	baseUrl := "https://quotes.toscrape.com"

	fmt.Println(header())
	fmt.Println(helpPrompt())

	for input := getInput(); input != ExitCommand; input = getInput() {
		switch input {
		case HelpCommand:
			fmt.Println(help())

		case StatusCommand:
			fmt.Println("baseUrl:   ", baseUrl)
			fmt.Println("datastore: ", ds.State())

		case RunCommand:
			start := time.Now()
			index, urlsVisited := crawl.Crawl(baseUrl, 1000)
			fmt.Printf("Crawled %q in %v\n", baseUrl, time.Since(start))
			ds.Load(index, urlsVisited)

		case LoadCommand:
			err := ds.LoadJSON()
			if err != nil {
				fmt.Println(err.Error())
			}

		case SaveCommand:
			err := ds.SaveJSON()
			if err != nil {
				fmt.Println(err.Error())
			}

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
			report.ReportVisitedUrls(urls, true)

		default:
			fmt.Printf("Invalid command: %q\n", input)
			fmt.Println(helpPrompt())
		}
	}
}
