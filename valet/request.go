package valet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// API makes a GET request to the Bank of Canada Valet API and returns the unmarshalled JSON response.
func API(URL string) (APIResponse, error) {
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &Client{logger: logger},
	}

	response, err := client.Get(URL)
	if err != nil {
		logger.Error("Failed to send request", zap.Error(err))
		return APIResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Failed to read response body", zap.Error(err))
		return APIResponse{}, err
	}

	var apiResponse APIResponse
	if err = json.Unmarshal(body, &apiResponse); err != nil {
		logger.Error("Failed to unmarshal response", zap.Error(err))
		return APIResponse{}, err
	}

	if response.StatusCode != http.StatusOK {
		logger.Warn(fmt.Sprintf("Non-OK status code: %d, message: %s", response.StatusCode, apiResponse.Message))
		return APIResponse{}, fmt.Errorf(apiResponse.Message)
	}

	return apiResponse, nil
}
