package telegraph_test

import (
	"net/http"
	"testing"

	"github.com/smirnoffmg/telegraph"
)

const (
	testPageResponse     = `{"ok":true,"result":{"path":"test-path","url":"https://example.com/test-path","title":"Test Page","content":[{"tag":"p","children":[{"tag":"Hello, world!"}]}],"author_name":"Tester","author_url":"https://example.com","views":0,"can_edit":true}}`
	testPageListResponse = `{"ok":true,"result":{"total_count":1,"pages":[{"path":"test-path","url":"https://example.com/test-path","title":"Test Page","content":[{"tag":"p","children":[{"tag":"Hello, world!"}]}],"author_name":"Tester","author_url":"https://example.com","views":0,"can_edit":true}]}}`
	testViewsResponse    = `{"ok":true,"result":{"path":"test-path","views":100}}`
	title                = "Test Page"
	path                 = "test-path"
)

func TestCreatePage(t *testing.T) {
	server := mockServer(testPageResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	content := []telegraph.Node{
		{
			Tag: "p",
			Children: []telegraph.Node{
				{
					Tag: "Hello, world!",
				},
			},
		},
	}

	page, err := client.CreatePage(accessToken, title, content, authorName, authorURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if page.Title != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, page.Title)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.CreatePage(accessToken, title, content, authorName, authorURL)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestEditPage(t *testing.T) {
	server := mockServer(testPageResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	content := []telegraph.Node{
		{
			Tag: "p",
			Children: []telegraph.Node{
				{
					Tag: "Hello, world!",
				},
			},
		},
	}

	page, err := client.EditPage(accessToken, path, title, content, authorName, authorURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if page.Title != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, page.Title)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.EditPage(accessToken, path, title, content, authorName, authorURL)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestGetPage(t *testing.T) {
	server := mockServer(testPageResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	page, err := client.GetPage(path, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if page.Title != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, page.Title)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.GetPage(path, true)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestGetPageList(t *testing.T) {
	server := mockServer(testPageListResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	pageList, err := client.GetPageList(accessToken, 0, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if pageList.TotalCount != 1 {
		t.Errorf("Expected TotalCount to be 1, got %d", pageList.TotalCount)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.GetPageList(accessToken, 0, 10)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestGetViews(t *testing.T) {
	server := mockServer(testViewsResponse, http.StatusOK)
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	views, err := client.GetViews(path, 2023, 1, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if views.Views != 100 {
		t.Errorf("Expected Views to be 100, got %d", views.Views)
	}

	// Test error response
	server = mockServer(testErrorResponse, http.StatusBadRequest)
	defer server.Close()

	client = telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")

	_, err = client.GetViews(path, 2023, 1, 1)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
