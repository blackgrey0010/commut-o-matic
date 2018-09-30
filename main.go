package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps" // https://godoc.org/googlemaps.github.io/maps
)

func main() {
	apiKey := os.Getenv("GCP_API_KEY")

	if apiKey == "" {
		log.Fatalf("GCP API key not set")

	}

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("API Client not created: %s", err)
	}
	r := &maps.DistanceMatrixRequest{
		Origins:       []string{"647 Central Ave, Alameda 94501"},
		Destinations:  []string{"61 Oak Grove St, San Francisco 94107"},
		DepartureTime: `now`,
		Mode:          `ModeDriving`,
	}
	resp, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	fmt.Println(resp)
	fmt.Println("")
	fmt.Println("Hello boobaby! Here is your commute:")
}
