package telegraph_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smirnoffmg/telegraph"
)

const (
	testAccountResponse  = `{"ok":true,"result":{"short_name":"Test","author_name":"Tester","author_url":"https://example.com","access_token":"123456","page_count":0}}`
	testPageResponse     = `{"ok":true,"result":{"path":"test-path","url":"https://example.com/test-path","title":"Test Page","content":[{"tag":"p","children":[{"tag":"Hello, world!"}]}],"author_name":"Tester","author_url":"https://example.com","views":0,"can_edit":true}}`
	testPageListResponse = `{"ok":true,"result":{"total_count":1,"pages":[{"path":"test-path","url":"https://example.com/test-path","title":"Test Page","content":[{"tag":"p","children":[{"tag":"Hello, world!"}]}],"author_name":"Tester","author_url":"https://example.com","views":0,"can_edit":true}]}}`
	testViewsResponse    = `{"ok":true,"result":{"path":"test-path","views":100}}`
	shortName            = "Test"
	title                = "Test Page"
)

func mockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/createAccount", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testAccountResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/getAccountInfo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testAccountResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/editAccountInfo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testAccountResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/revokeAccessToken", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testAccountResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/createPage", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testPageResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/editPage/test-path", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testPageResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/getPage/test-path", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testPageResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/getPageList", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testPageListResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/getViews", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(testViewsResponse)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
	return httptest.NewServer(mux)
}

func TestCreateAccount(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	account, err := client.CreateAccount(shortName, "Tester", "https://example.com")
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}
}

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	account, err := client.GetAccountInfo("123456", []string{"short_name", "author_name", "author_url", "page_count"})
	if err != nil {
		t.Fatalf("Failed to get account info: %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}
}

func TestEditAccountInfo(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	account, err := client.EditAccountInfo("123456", shortName, "Tester", "https://example.com")
	if err != nil {
		t.Fatalf("Failed to edit account info: %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}
}

func TestRevokeAccessToken(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	account, err := client.RevokeAccessToken("123456")
	if err != nil {
		t.Fatalf("Failed to revoke access token: %v", err)
	}
	if account.ShortName != shortName {
		t.Errorf("Expected ShortName to be '%s', got '%s'", shortName, account.ShortName)
	}
}

func TestCreatePage(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

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

	page, err := client.CreatePage("123456", title, content, "Tester", "https://example.com")
	if err != nil {
		t.Fatalf("Failed to create page: %v", err)
	}
	if page.Title != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, page.Title)
	}
}

func TestEditPage(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

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

	page, err := client.EditPage("123456", "test-path", title, content, "Tester", "https://example.com")
	if err != nil {
		t.Fatalf("Failed to edit page: %v", err)
	}
	if page.Title != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, page.Title)
	}
}

func TestGetPage(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	page, err := client.GetPage("test-path", true)
	if err != nil {
		t.Fatalf("Failed to get page: %v", err)
	}
	if page.Title != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, page.Title)
	}
}

func TestGetPageList(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	pageList, err := client.GetPageList("123456", 0, 10)
	if err != nil {
		t.Fatalf("Failed to get page list: %v", err)
	}
	if pageList.TotalCount != 1 {
		t.Errorf("Expected TotalCount to be 1, got %d", pageList.TotalCount)
	}
}

func TestGetViews(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	client := telegraph.NewClient(server.Client())
	client.SetBaseURL(server.URL + "/")
	client.SetDebug(true)

	views, err := client.GetViews("test-path", 2023, 1, 1)
	if err != nil {
		t.Fatalf("Failed to get views: %v", err)
	}
	if views.Views != 100 {
		t.Errorf("Expected Views to be 100, got %d", views.Views)
	}
}
