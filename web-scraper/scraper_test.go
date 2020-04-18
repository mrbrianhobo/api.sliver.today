package scraper

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	// Sliver URLs to crawl
	telegraphURL = "https://www.sliverpizzeria.com/pizza-telegraph/"
	shattuckURL  = "https://www.sliverpizzeria.com/pizza-shattuck/"
	broadwayURL  = "https://www.sliverpizzeria.com/pizza-broadway/"

	// // query selector for finding the menu items
	querySelector = "div .summary-excerpt p"
)

func TestRun(t *testing.T) {
	menu := ScrapeURL(telegraphURL, querySelector)
	fmt.Printf("today's sliver pizza is: %s\n", menu.Pizza)

	stringified, _ := json.Marshal(menu)
	fmt.Println(string(stringified))
}
