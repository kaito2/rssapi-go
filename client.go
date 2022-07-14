package rssapi

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

type Client struct {
	url        *url.URL
	apiKey     string
	httpClient *http.Client

	Logger *log.Logger
}

// NewClient returns new RSS API Client.
func NewClient(urlStr, apiKey string, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	// TODO: Configure HTTP Client.
	httpClient := &http.Client{}

	var discardLogger = log.New(io.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	return &Client{
		url:        parsedURL,
		apiKey:     apiKey,
		httpClient: httpClient,
		Logger:     logger,
	}, nil
}

func (c *Client) newGETRequest(ctx context.Context, spath string, query map[string]string) (*http.Request, error) {
	u := *c.url
	u.Path = path.Join(c.url.Path, spath)

	q := u.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new request")
	}

	req.Header.Add("X-API-KEY", c.apiKey)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
