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

// Client is a custom HTTP client that logs request and response details.
type Client struct {
	debug *log.Logger
}

// RoundTrip implements the http.RoundTripper interface.
func (c Client) RoundTrip(r *http.Request) (*http.Response, error) {
	isDebug := os.Getenv("LOGLEVEL") == "" || os.Getenv("LOGLEVEL") == "DEBUG"
	if isDebug {
		c.debug.Println(r.Method, r.URL, r.Proto)
	}
	response, err := http.DefaultTransport.RoundTrip(r)
	if response != nil && isDebug {
		c.debug.Println(response.Status)
	}
	return response, err
}

// Api makes a GET request to the Bank of Canada Valet API and returns the unmarshalled JSON response.
func Api(endpoint string) (ApiResponse, error) {
	transport := Client{
		debug: log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Lmicroseconds),
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &transport,
	}

	url := fmt.Sprintf("https://www.bankofcanada.ca/valet%s", endpoint)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error sending request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error reading response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var errResp ApiErrorResponse
		if err = json.Unmarshal(body, &errResp); err != nil {
			return ApiResponse{}, fmt.Errorf("error unmarshalling error response (status %d): %s", response.StatusCode, string(body))
		}
		return ApiResponse{}, fmt.Errorf("API error: %s", errResp.Message)
	}

	var apiResponse ApiResponse
	if err = json.Unmarshal(body, &apiResponse); err != nil {
		return ApiResponse{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return apiResponse, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for custom unmarshalling of Observation.
func (o *Observation) UnmarshalJSON(data []byte) error {
	raw := make(map[string]json.RawMessage)

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if val, ok := raw["d"]; ok {
		if err := json.Unmarshal(val, &o.Date); err != nil {
			return err
		}
	}

	if val, ok := raw["q"]; ok {
		if err := json.Unmarshal(val, &o.Quarter); err != nil {
			return err
		}
	}

	o.Series = make(map[string]SeriesObservation)
	for key, value := range raw {
		if key == "d" || key == "q" {
			continue
		}
		var fxRate SeriesObservation
		if err := json.Unmarshal(value, &fxRate); err != nil {
			return err
		}
		o.Series[key] = fxRate
	}

	return nil
}
