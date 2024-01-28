package ffo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FlagGetter interface {
	FlagList() ([]Flag, error)
}

type FliptClient struct {
	BaseURL string
}

func NewFliptClient(baseURL string) *FliptClient {
	return &FliptClient{BaseURL: baseURL}
}

func (f *FliptClient) FlagList() ([]Flag, error) {
	url := fmt.Sprintf("%s/api/flags", f.BaseURL)

	// Send GET request to Flipt API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feature flags: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch feature flags, status code: %d", resp.StatusCode)
	}

	// Decode response body
	var flags []Flag
	err = json.NewDecoder(resp.Body).Decode(&flags)
	if err != nil {
		return nil, fmt.Errorf("failed to decode feature flags response: %v", err)
	}

	return flags, nil
}
