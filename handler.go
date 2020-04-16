package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/mrbrianhobo/sliver.today/common"
	scraper "github.com/mrbrianhobo/sliver.today/web-scraper"
)

var (
	// Sliver URLs to crawl
	// Note: uses Squarespace URL queries (https://developers.squarespace.com/url-queries)
	// since normal website caching is weird
	telegraphURL = "https://www.sliverpizzeria.com/pizza-telegraph/?format=main-content"
	shattuckURL  = "https://www.sliverpizzeria.com/pizza-shattuck/?format=main-content"
	broadwayURL  = "https://www.sliverpizzeria.com/pizza-broadway/?format=main-content"

	// Today's Pizza is stored at index 2 of the query
	querySelector    = ".row .sqs-row p"
	todaysPizzaIndex = 2

	telegraphKeywords = []string{"telegraph", "southside"}
	shattuckKeywords  = []string{"shattuck", "downtown", "downtown-berkeley"}
	broadwayKeywords  = []string{"broadway", "oakland"}
)

var decoder = schema.NewDecoder()

// Handler is the HTTP Cloud Function entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	pizzas := make([]*common.Menu, 0)
	values := r.URL.Query()

	switch loc := values.Get("location"); {
	case common.MatchKeywords(loc, telegraphKeywords):
		menu := scraper.ScrapeURL(telegraphURL, querySelector, todaysPizzaIndex)
		menu.SetLocation(common.Telegraph)
		pizzas = append(pizzas, menu)
	case common.MatchKeywords(loc, shattuckKeywords):
		menu := scraper.ScrapeURL(shattuckURL, querySelector, todaysPizzaIndex)
		menu.SetLocation(common.Shattuck)
		pizzas = append(pizzas, menu)
	case common.MatchKeywords(loc, broadwayKeywords):
		menu := scraper.ScrapeURL(broadwayURL, querySelector, todaysPizzaIndex)
		menu.SetLocation(common.Broadway)
		pizzas = append(pizzas, menu)
	default:
		telegraph := scraper.ScrapeURL(telegraphURL, querySelector, todaysPizzaIndex)
		telegraph.SetLocation(common.Telegraph)

		shattuck := scraper.ScrapeURL(shattuckURL, querySelector, todaysPizzaIndex)
		shattuck.SetLocation(common.Shattuck)

		broadway := scraper.ScrapeURL(broadwayURL, querySelector, todaysPizzaIndex)
		broadway.SetLocation(common.Broadway)

		pizzas = append(pizzas, []*common.Menu{telegraph, shattuck, broadway}...)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizzas)
}
