package examples

import (
	"fmt"
	"net/http"

	"github.com/smirnoffmg/telegraph"
)

func main() {
	// Create a new Telegraph account
	client := telegraph.NewClient(http.DefaultClient)
	account, err := client.CreateAccount("ShortName", "AuthorName", "AuthorURL")
	if err != nil {
		fmt.Println("Failed to create account:", err)
		return
	}

	fmt.Printf("Created account: %v\n", account)

	// Get account information
	accountInfo, err := client.GetAccountInfo(account.AccessToken, []string{"short_name", "author_name"})

	if err != nil {
		fmt.Println("Failed to get account info:", err)
		return
	}

	fmt.Printf("Account info: %v\n", accountInfo)

	// Create a new telegraph page from HTML
	page, err := client.CreatePageFromHTML(account.AccessToken, "Title", "<p>Hello, world!</p>", account.AuthorName, account.AuthorURL)

	if err != nil {
		fmt.Println("Failed to create page:", err)
		return
	}

	fmt.Printf("Created page with URL: %v\n", page.URL)
}
