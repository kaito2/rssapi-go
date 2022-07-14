package main

import (
	"context"
	"github.com/kaito2/rssapi-go"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("RSS_API_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the environment variable named RSS_API_API_KEY.")
	}

	client, err := rssapi.NewClient(
		"https://api.rssapi.net",
		apiKey,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	res, err := client.CreateSubscription(
		ctx, "https://www.nasa.gov/rss/dyn/breaking_news.rss", "test-info")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(*res)
}
