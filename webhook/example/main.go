package main

import (
	"github.com/kaito2/rssapi-go/webhook"
	"io"
	"log"
	"strings"
)

// ref: https://docs.rssapi.net/webhooks?id=example-webhook
const exampleBody = `
{
  "webhook_from": "rssapi.net",
  "webhook_reason": "new entries in feed",
  "new_entries_count": 2,
  "subscription_id": 22798,
  "last_webhook_response": 200,
  "is_retry": false,
  "info": "Your subscription info",
  "feed": {
    "title": "NASA Breaking News",
    "description": "A RSS news feed containing the latest NASA news articles and press releases.",
    "homepage": "http://www.nasa.gov/",
    "feed_url": "https://www.nasa.gov/rss/dyn/breaking_news.rss",
    "additional_details": {
      "language": "en-us",
      "categories": [],
      "authors": [],
      "copyright": null
    }
  },
  "new_entries": [
    {
      "title": "NASA Astronauts to Call Indiana, Texas, Virginia Students from Space",
      "link": "http://www.nasa.gov/press-release/nasa-astronauts-to-call-indiana-texas-virginia-students-from-space",
      "description": "Students across the country will have three opportunities to hear from NASA astronauts aboard the International Space Station.",
      "guid": "http://www.nasa.gov/press-release/nasa-astronauts-to-call-indiana-texas-virginia-students-from-space",
      "time": "Fri, 04 Feb 2022 20:13:00 +0000",
      "timestamp": 1644005580
    },
    {
      "title": "NASA Invites Media to Learn About Mission Studying Snowstorms",
      "link": "http://www.nasa.gov/press-release/nasa-invites-media-to-learn-about-mission-studying-snowstorms",
      "description": "NASA will hold a briefing for members of the media to learn more about an airborne science campaign studying snowstorms at 11 a.m. EST Thursday, Feb. 10. The agency will stream audio of the call live online.",
      "guid": "http://www.nasa.gov/press-release/nasa-invites-media-to-learn-about-mission-studying-snowstorms",
      "time": "Fri, 04 Feb 2022 16:53:00 +0000",
      "timestamp": 1643993580
    }
  ]
}`

func main() {
	body := io.NopCloser(strings.NewReader(exampleBody))
	parsedBody, err := webhook.ParseBody(body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", parsedBody)
}
