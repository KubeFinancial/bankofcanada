package valet

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Client is a custom HTTP client that logs request and response details
type Client struct {
	log *log.Logger
}

// RoundTrip implements the http.RoundTripper interface
func (c Client) RoundTrip(r *http.Request) (*http.Response, error) {
	c.log.Printf("%s %s %s", r.Method, r.URL, r.Proto)
	resp, err := http.DefaultTransport.RoundTrip(r)
	if resp != nil {
		c.log.Println(resp.Status)
	}
	return resp, err
}

// Api makes a GET request to the Bank of Canada Valet API and returns the unmarshalled JSON response
func Api(endpoint string) (interface{}, error) {
	transport := Client{
		log: log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds),
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &transport,
	}

	url := fmt.Sprintf("https://www.bankofcanada.ca/valet%s", endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "bankofcanada_go/0.0.1")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, ioErr := io.ReadAll(resp.Body)
		if ioErr != nil {
			return nil, fmt.Errorf("error reading response: %w", ioErr)
		}
		var errorResponse struct {
			Message string `json:"message"`
		}
		if err = json.Unmarshal(body, &errorResponse); err != nil {
			return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("API error: %s", errorResponse.Message)
	}

	var result interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
