package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mrbrianhobo/sliver.today/api"
)

// Handler is the HTTP Cloud Function entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	api.Handler(w, r)
}

func main() {
	// manually set time zone
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Fatalf("error loading location '%s': %v\n", tz, err)
		}
	}

	http.HandleFunc("/", api.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
