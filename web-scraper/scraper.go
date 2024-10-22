package scraper

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/mrbrianhobo/sliver.today/common"
)

func ScrapeURL(url string, querySelector string) *common.Menu {
	menu := &common.Menu{}
	now := time.Now().In(time.Local)
	index := -1

	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback
	c.OnHTML(querySelector, func(e *colly.HTMLElement) {
		// Match day of the week and ensure first instance
		if strings.EqualFold(strings.TrimSpace(e.Text), now.Weekday().String()) && index == -1 {
			index = e.Index + 2
		}

		if e.Index == index {
			menu.Pizza = common.SanitizeStr(e.Text)
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with response: %v\n", r.Request.URL, r)
		log.Printf("Error: %v\n", err)
	})

	// Start scraping
	c.Visit(url)
	menu.SetTime(now)

	return menu
}
