package valet

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	defaultTimeout = 10 * time.Second
)

// Client is a custom HTTP client that logs request and response details.
type Client struct {
	logger *log.Logger
}

// RoundTrip implements the http.RoundTripper interface.
func (c *Client) RoundTrip(r *http.Request) (*http.Response, error) {
	c.logger.Printf("Request: %s %s", r.Method, r.URL)

	response, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if response != nil {
		if response.StatusCode == http.StatusOK {
			c.logger.Printf(
				"Response: %s, Filename: %s, Generated: %s UTC",
				response.Status,
				strings.TrimPrefix(
					response.Header.Get("Content-Disposition"),
					"attachment; filename=",
				),
				response.Header.Get("X-Generated"),
			)
		} else {
			c.logger.Printf("Response: %s", response.Status)
		}
	}

	return response, nil
}

// API makes a GET request to the Bank of Canada Valet API and returns the unmarshalled JSON response.
func API(url string) (APIResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	client := &http.Client{
		Transport:     &Client{logger: log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       defaultTimeout,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return APIResponse{}, fmt.Errorf("create request: %w", err)
	}

	response, err := client.Do(req)
	if err != nil {
		return APIResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("read response body: %w", err)
	}

	var apiResponse APIResponse
	if err = json.Unmarshal(body, &apiResponse); err != nil {
		return APIResponse{}, fmt.Errorf("unmarshal JSON: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf("API error: %s", apiResponse.Message)
	}

	return apiResponse, nil
}
