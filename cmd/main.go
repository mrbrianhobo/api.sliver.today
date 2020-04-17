package main

import (
	"encoding/json"
	"fmt"

	scraper "github.com/mrbrianhobo/sliver.today/web-scraper"
)

const (
	// Sliver URLs to crawl
	telegraphURL = "https://www.sliverpizzeria.com/pizza-telegraph/"
	shattuckURL  = "https://www.sliverpizzeria.com/pizza-shattuck/"
	broadwayURL  = "https://www.sliverpizzeria.com/pizza-broadway/"

	// Today's Pizza is stored at index 2 of the query
	querySelector = "div .summary-excerpt p"
)

func main() {

	menu := scraper.ScrapeURL(telegraphURL, querySelector)
	fmt.Printf("today's sliver pizza is: %s\n", menu.Pizza)

	stringified, _ := json.Marshal(menu)
	fmt.Println(string(stringified))
	// http.HandleFunc("/", api.Handler)
	// fmt.Println("Listening on localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
