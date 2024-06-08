package telegraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for the Telegraph API
	DefaultBaseURL = "https://api.telegra.ph/"
	// RequestTimeout is the timeout duration for HTTP requests
	RequestTimeout = 10 * time.Second
)

// Client is a HTTP client for the Telegraph API
type Client struct {
	httpClient *http.Client
	baseURL    string
	debug      bool
}

// NewClient creates a new Telegraph API client with an optional custom http.Client and base URL
func NewClient(customClient ...*http.Client) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: RequestTimeout,
		},
		baseURL: DefaultBaseURL,
		debug:   false,
	}
	if len(customClient) > 0 && customClient[0] != nil {
		client.httpClient = customClient[0]
	}
	return client
}

// SetBaseURL sets a custom base URL for the client
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// SetDebug sets the debug mode for the client
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// doRequest sends a HTTP request to the Telegraph API
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
