package telegraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client represents a client for the Telegraph API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	debug      bool
}

// NewClient creates a new Telegraph API client.
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    "https://api.telegra.ph/",
	}
}

// SetBaseURL sets the base URL for API requests.
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// BaseURL returns the base URL for API requests.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// SetDebug enables or disables debug mode.
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// Debug returns whether debug mode is enabled.
func (c *Client) Debug() bool {
	return c.debug
}

// doRequest sends a HTTP request to the Telegraph API.
func (c *Client) doRequest(method, endpoint string, body interface{}, result interface{}) error {
	url := c.baseURL + endpoint

	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	if c.debug {
		fmt.Printf("Sending request to %s with method %s and body %s\n", url, method, string(reqBody))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if c.debug {
		fmt.Printf("Received response: %+v\n", result)
	}
	return nil
}
