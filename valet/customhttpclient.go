package valet

import (
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// Client is a custom HTTP client that logs request and response details.
type Client struct {
	logger *zap.Logger
}

// RoundTrip implements the http.RoundTripper interface.
func (c Client) RoundTrip(r *http.Request) (*http.Response, error) {
	// log the request
	c.logger.Debug(fmt.Sprintf("Request: %s %s", r.Method, r.URL))

	// make the request
	response, err := http.DefaultTransport.RoundTrip(r)

	// log the response
	if response != nil {
		if response.StatusCode == http.StatusOK {
			c.logger.Debug(fmt.Sprintf(
				"Response: %s, Filename: %s, Generated: %s UTC",
				response.Status,
				strings.TrimPrefix(response.Header.Get("Content-Disposition"), "attachment; filename="),
				response.Header.Get("X-Generated")))
		} else {
			c.logger.Debug(fmt.Sprintf("Response: %s", response.Status))
		}
	}

	return response, err
}
