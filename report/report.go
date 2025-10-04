package report

import (
	"fmt"
	"slices"
	"strings"
)

const (
	outterDivider = "+--------------------------------------------------------------"
	totalVisited  = "| Total URLs Visited: "
	totalWords    = "| Total Words: "
	innerDivider  = "|+-------------------------------------------------------------"
	innerUrl      = "|| -> "
	innerWord     = "|| Word: "
	innerEmpty    = "||"
)

// Word -> URLs -> Occurence count
type InvertedIndex = map[string]map[string]uint

func ReportVisitedUrls(urls []string, sort bool) {
	if sort {
		slices.Sort(urls)
	}

	var sb strings.Builder

	sb.WriteString(outterDivider + "\n")

	if len(urls) > 0 {
		sb.WriteString(innerDivider + "\n")
		for _, url := range urls {
			sb.WriteString(innerUrl + url + "\n")
		}
		sb.WriteString(innerDivider + "\n")
	}

	sb.WriteString(fmt.Sprintf("%v%v\n", totalVisited, len(urls)))
	sb.WriteString(outterDivider)

	fmt.Println(sb.String())
}

func ReportIndex(index InvertedIndex) {
	var sb strings.Builder

	sb.WriteString(outterDivider + "\n")

	if len(index) > 0 {
		for word, urlCount := range index {
			sb.WriteString(innerDivider + "\n")
			sb.WriteString(innerWord + word + "\n")
			sb.WriteString(innerEmpty + "\n")
			for url, count := range urlCount {
				sb.WriteString(fmt.Sprintf("%v    %q: %v\n", innerEmpty, url, count))
			}
			sb.WriteString(innerDivider + "\n")
		}
	}

	sb.WriteString(fmt.Sprintf("%v%v\n", totalWords, len(index)))
	sb.WriteString(outterDivider)

	fmt.Println(sb.String())
}
