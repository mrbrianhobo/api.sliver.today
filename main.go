package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

const (
	// Sliver URLs to crawl
	telegraphURL = "https://www.sliverpizzeria.com/pizza-telegraph/"
	shattuckURL  = "https://www.sliverpizzeria.com/pizza-shattuck/"
	broadwayURL  = "https://www.sliverpizzeria.com/pizza-broadway/"

	// Today's Pizza is stored at index 2 of the query for "div .summary-excerpt p"
	todaysPizzaIndex = 2
)

func trimStr(in string) string {
	space := regexp.MustCompile(`\s+`)
	out := space.ReplaceAllString(in, " ")
	return out
}

func main() {
	var todaysPizza string

	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback
	c.OnHTML("div .summary-excerpt p", func(e *colly.HTMLElement) {
		// fmt.Println(e)
		if e.Index == todaysPizzaIndex {
			todaysPizza = e.Text
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	c.Visit(telegraphURL)

	fmt.Printf("today's sliver pizza is: %s\n", trimStr(todaysPizza))
}
