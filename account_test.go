package telegraph_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smirnoffmg/telegraph"
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

func TestCreateAccount(t *testing.T) {
	server := mockServer(testAccountResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	account, err := client.CreateAccount(shortName, authorName, authorURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.CreateAccount(shortName, authorName, authorURL)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestGetAccountInfo(t *testing.T) {
	server := mockServer(testAccountResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	account, err := client.GetAccountInfo(accessToken, []string{"short_name", "author_name", "author_url", "page_count"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.GetAccountInfo(accessToken, []string{"short_name", "author_name", "author_url", "page_count"})
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestEditAccountInfo(t *testing.T) {
	server := mockServer(testAccountResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	account, err := client.EditAccountInfo(accessToken, shortName, authorName, authorURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.EditAccountInfo(accessToken, shortName, authorName, authorURL)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestRevokeAccessToken(t *testing.T) {
	server := mockServer(testAccountResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	account, err := client.RevokeAccessToken(accessToken)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.RevokeAccessToken(accessToken)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
