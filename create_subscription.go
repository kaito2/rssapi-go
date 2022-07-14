package rssapi

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

type CreateSubscriptionResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Status         string `json:"status"`          // e.g. "subscribed"
		SubscriptionID string `json:"subscription_id"` // e.g. "22946
		FeedType       string `json:"feed_type"`       // e.g. "rss"
		WebhookURL     string `json:"webhook_url"`     // e.g. "https://example.org/webhook"
		URL            string `json:"url"`             // e.g. "https://www.nasa.gov/rss/dyn/breaking_news.rss"
		Info           string `json:"info"`
	} `json:"result"`
}

// CreateSubscription Create a Subscription.
// TODO: Make `info` to optional.
func (c *Client) CreateSubscription(ctx context.Context, url, info string) (*CreateSubscriptionResponse, error) {
	// TODO: Set header for authentication.
	spath := fmt.Sprintf("/v1/header/subscribe")
	q := map[string]string{
		"url":  url,
		"info": info,
	}
	req, err := c.newGETRequest(ctx, spath, q)
	if err != nil {
		// TODO: Define custom error type.
		return nil, errors.Wrap(err, "failed to create GET request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Do http request")
	}

	// TODO: Check status here...

	var createSubscriptionResponse CreateSubscriptionResponse
	if err := decodeBody(res, &createSubscriptionResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode response body")
	}

	return &createSubscriptionResponse, nil
}
