package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps" // https://godoc.org/googlemaps.github.io/maps
)

const (
	Origin      = "647 Central Ave, Alameda 94501"
	Destination = "61 Oak Grove St, San Francisco 94107"
)

func main() {
	apiKey := os.Getenv("GCP_API_KEY")

	if apiKey == "" {
		log.Fatalf("GCP API key not set")

	}

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("API Client not created: %s", err)
	}

	driving := getTravelTimes(maps.TravelModeDriving, client)
	transit := getTravelTimes(maps.TravelModeTransit, client)

	outputResponse(driving)
	outputResponse(transit)

	quickestMode := "Driving"
	if driving.Rows[0].Elements[0].Duration > transit.Rows[0].Elements[0].Duration {
		quickestMode = "Transit"
	}

	fmt.Println("")
	fmt.Printf("Hello boobaby! Here is your commute: %s!\n", quickestMode)
}

func getTravelTimes(mode maps.Mode, c *maps.Client) *maps.DistanceMatrixResponse {
	drivingReq := &maps.DistanceMatrixRequest{
		Origins:       []string{Origin},
		Destinations:  []string{Destination},
		DepartureTime: `now`,
		Mode:          mode,
	}

	resp, err := c.DistanceMatrix(context.Background(), drivingReq)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return resp
}

func outputResponse(resp *maps.DistanceMatrixResponse) {
	for _, row := range resp.Rows {
		for _, element := range row.Elements {
			fmt.Printf("Distance:%s Duration:%s InTraffic:%s\n", element.Distance.HumanReadable, element.Duration, element.DurationInTraffic)
		}
	}
}
