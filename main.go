package main

import (
	"fmt"

	"github.com/mrbrianhobo/sliver.today/common"
	scraper "github.com/mrbrianhobo/sliver.today/web-scraper"
)

const (
	// Sliver URLs to crawl
	// Note: uses Squarespace URL queries (https://developers.squarespace.com/url-queries)
	// since normal website caching is weird
	telegraphURL = "https://www.sliverpizzeria.com/pizza-telegraph/?format=main-content"
	shattuckURL  = "https://www.sliverpizzeria.com/pizza-shattuck/?format=main-content"
	broadwayURL  = "https://www.sliverpizzeria.com/pizza-broadway/?format=main-content"

	// Today's Pizza is stored at index 2 of the query
	querySelector    = ".row .sqs-row p"
	todaysPizzaIndex = 2
)

func main() {
	menu := scraper.ScrapeURL(telegraphURL, querySelector, todaysPizzaIndex)
	fmt.Printf("today's sliver pizza is: %s\n", common.TrimStr(menu.Pizza))
}
