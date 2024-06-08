# Telegraph Go client

## Example

```go
package main

import (
 "fmt"
 "log"
 "net/http"
 "time"

 "github.com/smirnoffmg/telegraph"
)

func main() {
 // Create a new Telegraph client with a custom HTTP client (optional)
 httpClient := &http.Client{
  Timeout: time.Second * 10,
 }
 client := telegraph.NewClient(httpClient)

 // Set the debug mode to true to enable detailed logging
 client.SetDebug(true)

 // Set the base URL (optional, only needed for testing)
 // client.SetBaseURL("https://api.telegra.ph/")

 // Create an account
 account, err := client.CreateAccount("Test", "Tester", "https://example.com")
 if err != nil {
  log.Fatalf("Failed to create account: %v", err)
 }
 fmt.Printf("Created account: %+v\n", account)

 // Get account information
 accountInfo, err := client.GetAccountInfo(account.AccessToken, []string{"short_name", "author_name", "author_url", "page_count"})
 if err != nil {
  log.Fatalf("Failed to get account info: %v", err)
 }
 fmt.Printf("Account info: %+v\n", accountInfo)

 // Edit account information
 editedAccount, err := client.EditAccountInfo(account.AccessToken, "NewTest", "NewTester", "https://newexample.com")
 if err != nil {
  log.Fatalf("Failed to edit account info: %v", err)
 }
 fmt.Printf("Edited account: %+v\n", editedAccount)

 // Revoke access token
 newAccount, err := client.RevokeAccessToken(account.AccessToken)
 if err != nil {
  log.Fatalf("Failed to revoke access token: %v", err)
 }
 fmt.Printf("New account with revoked access token: %+v\n", newAccount)

 // Create a page
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
 page, err := client.CreatePage(newAccount.AccessToken, "Test Page", content, "Tester", "https://example.com")
 if err != nil {
  log.Fatalf("Failed to create page: %v", err)
 }
 fmt.Printf("Created page: %+v\n", page)

 // Get a page
 retrievedPage, err := client.GetPage("test-path", true)
 if err != nil {
  log.Fatalf("Failed to get page: %v", err)
 }
 fmt.Printf("Retrieved page: %+v\n", retrievedPage)

 // Edit a page
 editedPage, err := client.EditPage(newAccount.AccessToken, "test-path", "New Test Page", content, "NewTester", "https://newexample.com")
 if err != nil {
  log.Fatalf("Failed to edit page: %v", err)
 }
 fmt.Printf("Edited page: %+v\n", editedPage)

 // Get page list
 pageList, err := client.GetPageList(newAccount.AccessToken, 0, 10)
 if err != nil {
  log.Fatalf("Failed to get page list: %v", err)
 }
 fmt.Printf("Page list: %+v\n", pageList)

 // Get views
 views, err := client.GetViews("test-path", 2023, 1, 1)
 if err != nil {
  log.Fatalf("Failed to get views: %v", err)
 }
 fmt.Printf("Page views: %+v\n", views)
}


```
