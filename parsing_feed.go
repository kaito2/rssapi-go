package rssapi

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
)

type ParsingFeedRequest struct {
	URL    string
	Limit  *int
	Search *string
	Sort   *string
}

type ParsingFeedResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		// NOTE: The type is not specified in the documentation.
		// Settings []any `json:"settings"`
		Info struct {
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
		} `json:"info"`
		Entries []struct {
			Title       string `json:"title"`
			Link        string `json:"link"`
			Description string `json:"description"`
			GUID        string `json:"guid"`
			Time        string `json:"time"`
			Timestamp   int    `json:"timestamp"`
		} `json:"entries"`
	} `json:"result"`
}

func (c *Client) ParsingFeed(
	ctx context.Context,
	req *ParsingFeedRequest,
) (*ParsingFeedResponse, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}

	q := map[string]string{
		"url": req.URL,
	}
	if req.Limit != nil {
		q["limit"] = strconv.Itoa(*req.Limit)
	}
	if req.Search != nil {
		q["search"] = *req.Search
	}
	if req.Sort != nil {
		q["sort"] = *req.Sort
	}

	spath := "/v1/header/get"
	httpReq, err := c.newGETRequest(ctx, spath, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create GET request")
	}

	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Do http request")
	}

	// TODO: Check status here...

	var parsingFeedResponse ParsingFeedResponse
	if err := decodeBody(res, &parsingFeedResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode response body")
	}

	return &parsingFeedResponse, nil
}
