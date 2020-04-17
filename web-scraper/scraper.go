package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/mrbrianhobo/sliver.today/common"
)

func ScrapeURL(url string, querySelector string) *common.Menu {
	menu := &common.Menu{}
	now := time.Now().In(common.Timezone)
	index := -1

	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback
	c.OnHTML(querySelector, func(e *colly.HTMLElement) {
		if strings.EqualFold(e.Text, now.Weekday().String()) {
			index = e.Index + 2
		}

		if e.Index == index {
			menu.Pizza = common.SanitizeStr(e.Text)
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
	})

	// Start scraping
	c.Visit(url)
	menu.SetTime(now)

	return menu
}
