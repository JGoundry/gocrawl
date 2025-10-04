package report

import (
	"fmt"
	"strings"
)

const (
	outterDivider = "+--------------------------------------------------------------"
	totalVisited  = "| Total URLs Visited: "
	innerDivider  = "|+-------------------------------------------------------------"
	innerUrl      = "|| -> "
)

func ReportVisitedUrls(urls []string) {
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
