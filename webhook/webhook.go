package webhook

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

// Body RSS API webhook body.
// for more details: https://docs.rssapi.net/webhooks?id=example-webhook
type Body struct {
	WebhookFrom         string `json:"webhook_from"`
	WebhookReason       string `json:"webhook_reason"`
	NewEntriesCount     int    `json:"new_entries_count"`
	SubscriptionID      int    `json:"subscription_id"`
	LastWebhookResponse int    `json:"last_webhook_response"`
	IsRetry             bool   `json:"is_retry"`
	Info                string `json:"info"`
	Feed                struct {
		Title             string `json:"title"`
		Description       string `json:"description"`
		Homepage          string `json:"homepage"`
		FeedURL           string `json:"feed_url"`
		AdditionalDetails struct {
			Language string `json:"language"`
			// NOTE: The type is not specified in the documentation.
			// Categories []any `json:"categories"`
			// NOTE: The type is not specified in the documentation.
			// Authors []any `json:"authors"`
			// NOTE: The type is not specified in the documentation.
			Copyright *string `json:"copyright"`
		} `json:"additional_details"`
	} `json:"feed"`
	NewEntries []struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		GUID        string `json:"guid"`
		Time        string `json:"time"`
		Timestamp   int    `json:"timestamp"`
	} `json:"new_entries"`
}

func ParseBody(body io.ReadCloser) (*Body, error) {
	defer body.Close()
	decoder := json.NewDecoder(body)
	var parsedBody Body
	if err := decoder.Decode(&parsedBody); err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}
	return &parsedBody, nil
}
