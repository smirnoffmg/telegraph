package telegraph

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testAccountResponse = `{"ok":true,"result":{"short_name":"Test","author_name":"Tester","author_url":"https://example.com","access_token":"123456","page_count":0}}`
	testErrorResponse   = `{"ok":false,"error":"test error"}`
	shortName           = "Test"
	authorName          = "Tester"
	authorURL           = "https://example.com"
	accessToken         = "123456"
)

func mockServer(response string, statusCode int) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		if _, err := w.Write([]byte(response)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	return httptest.NewServer(mux)
}

func TestDoRequest(t *testing.T) {
	server := mockServer(testAccountResponse, http.StatusOK)
	defer server.Close()

	client := NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	var result struct {
		Ok     bool    `json:"ok"`
		Result Account `json:"result"`
	}

	// Test successful request
	err := client.doRequest(http.MethodPost, "createAccount", map[string]string{
		"short_name":  shortName,
		"author_name": authorName,
		"author_url":  authorURL,
	}, &result)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result.Ok {
		t.Errorf("Expected Ok to be true, got false")
	}

	// Test JSON marshal error
	invalidBody := make(chan int) // channels cannot be marshaled to JSON
	err = client.doRequest(http.MethodPost, "createAccount", invalidBody, &result)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	} else {
		fmt.Printf("Received expected error: %v\n", err)
	}

	// Test request creation error
	client.SetBaseURL(string([]byte{0x7f})) // invalid URL to cause error
	err = client.doRequest(http.MethodPost, "createAccount", map[string]string{
		"short_name":  shortName,
		"author_name": authorName,
		"author_url":  authorURL,
	}, &result)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	} else {
		fmt.Printf("Received expected error: %v\n", err)
	}

	// Test unexpected status code
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	err = client.doRequest(http.MethodPost, "createAccount", map[string]string{
		"short_name":  shortName,
		"author_name": authorName,
		"author_url":  authorURL,
	}, &result)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	} else {
		fmt.Printf("Received expected error: %v\n", err)
	}

	// Test JSON decode error
	server = mockServer(`{"ok":true,"result":`, http.StatusOK) // malformed JSON response
	defer server.Close()

	client = NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	err = client.doRequest(http.MethodPost, "createAccount", map[string]string{
		"short_name":  shortName,
		"author_name": authorName,
		"author_url":  authorURL,
	}, &result)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	} else {
		fmt.Printf("Received expected error: %v\n", err)
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient(&http.Client{})

	if client == nil {
		t.Fatalf("Expected client to be non-nil")
	}
}

func TestSetBaseURL(t *testing.T) {
	client := NewClient(&http.Client{})
	client.SetBaseURL("https://example.com")

	if client.baseURL != "https://example.com" {
		t.Fatalf("Expected base URL to be 'https://example.com', got '%s'", client.baseURL)
	}
}

func TestSetDebug(t *testing.T) {
	client := NewClient(&http.Client{})
	client.SetDebug(true)

	if !client.debug {
		t.Fatalf("Expected debug to be true, got false")
	}
}
