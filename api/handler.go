package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mrbrianhobo/sliver.today/common"
	scraper "github.com/mrbrianhobo/sliver.today/web-scraper"
)

var (
	// Sliver URLs to crawl
	telegraphURL = "https://www.sliverpizzeria.com/pizza-telegraph/"
	shattuckURL  = "https://www.sliverpizzeria.com/pizza-shattuck/"
	broadwayURL  = "https://www.sliverpizzeria.com/pizza-broadway/"

	// query selector for finding the menu items
	querySelector = "div .summary-excerpt p"

	telegraphKeywords = []string{"telegraph", "southside"}
	shattuckKeywords  = []string{"shattuck", "downtown", "downtown-berkeley"}
	broadwayKeywords  = []string{"broadway", "oakland"}
)

// Handler is the HTTP Cloud Function entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	pizzas := make([]*common.Menu, 0)
	values := r.URL.Query()

	log.Printf("API Request URL: %s\n", r.URL.String())
	switch loc := values.Get("location"); {
	case loc == "", strings.EqualFold(loc, "all"):
		telegraph := scraper.ScrapeURL(telegraphURL, querySelector)
		telegraph.SetLocation(common.Telegraph)

		shattuck := scraper.ScrapeURL(shattuckURL, querySelector)
		shattuck.SetLocation(common.Shattuck)

		broadway := scraper.ScrapeURL(broadwayURL, querySelector)
		broadway.SetLocation(common.Broadway)

		pizzas = append(pizzas, []*common.Menu{telegraph, shattuck, broadway}...)

		log.Printf("%s pizza: %s\n", common.Telegraph, telegraph.Pizza)
		log.Printf("%s pizza: %s\n", common.Shattuck, shattuck.Pizza)
		log.Printf("%s pizza: %s\n", common.Broadway, broadway.Pizza)
	case common.MatchKeywords(loc, telegraphKeywords):
		menu := scraper.ScrapeURL(telegraphURL, querySelector)
		menu.SetLocation(common.Telegraph)
		pizzas = append(pizzas, menu)

		log.Printf("%s pizza: %s\n", common.Telegraph, menu.Pizza)
	case common.MatchKeywords(loc, shattuckKeywords):
		menu := scraper.ScrapeURL(shattuckURL, querySelector)
		menu.SetLocation(common.Shattuck)
		pizzas = append(pizzas, menu)

		log.Printf("%s pizza: %s\n", common.Shattuck, menu.Pizza)
	case common.MatchKeywords(loc, broadwayKeywords):
		menu := scraper.ScrapeURL(broadwayURL, querySelector)
		menu.SetLocation(common.Broadway)
		pizzas = append(pizzas, menu)

		log.Printf("%s pizza: %s\n", common.Broadway, menu.Pizza)
	}

	menus := common.Menus{
		Pizzas: pizzas,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}
