package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/mrbrianhobo/sliver.today/common"
)

// TODO: update menu fields and update error
func ScrapeURL(url string, querySelector string, index int) *common.Menu {
	menu := &common.Menu{}

	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback
	c.OnHTML(querySelector, func(e *colly.HTMLElement) {
		if e.Index == index {
			menu.Pizza = common.TrimStr(e.Text)
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	c.Visit(url)
	menu.SetTime()

	return menu
}
