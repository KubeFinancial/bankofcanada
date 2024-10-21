package valet

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Client is a custom HTTP client that logs request and response details.
type Client struct {
	logger *log.Logger
}

// RoundTrip implements the http.RoundTripper interface.
func (c Client) RoundTrip(r *http.Request) (*http.Response, error) {
	// log the request
	c.logger.Printf("Request: %s %s", r.Method, r.URL)

	// make the request
	response, err := http.DefaultTransport.RoundTrip(r)

	// log the response
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

	return response, err
}

// API makes a GET request to the Bank of Canada Valet API and returns the unmarshalled JSON response.
func API(URL string) (APIResponse, error) {
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &Client{logger: log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)},
	}

	response, err := client.Get(URL)
	if err != nil {
		return APIResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return APIResponse{}, err
	}

	var apiResponse APIResponse
	if err = json.Unmarshal(body, &apiResponse); err != nil {
		return APIResponse{}, err
	}

	if response.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf(apiResponse.Message)
	}

	return apiResponse, nil
}
