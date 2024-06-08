
# Telegraph API Client in Go

![Go](https://img.shields.io/badge/Go-1.16%2B-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Build Status](https://github.com/smirnoffmg/telegraph/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/smirnoffmg/telegraph/branch/main/graph/badge.svg)](https://codecov.io/gh/smirnoffmg/telegraph)

This repository contains a Go client library for the [Telegraph API](https://telegra.ph/api). Telegraph is a publishing tool that allows you to create formatted posts with rich media, and this client enables developers to interact programmatically with the Telegraph API.

## Features

- **Account Management**: Create and manage Telegraph accounts.
- **Page Management**: Create, edit, and retrieve Telegraph pages.
- **Statistics**: Retrieve view statistics for Telegraph pages.
- **Customizable**: Easy to extend and customize for specific needs.
- **Debug Mode**: Enable detailed logging for debugging purposes.

## Installation

To install the package, use the following command:

```sh
go get github.com/smirnoffmg/telegraph
```

## Usage

### Creating a New Account

Here's a simple example of how to create a new Telegraph account:

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
    httpClient := &http.Client{
        Timeout: time.Second * 10,
    }
    client := telegraph.NewClient(httpClient)
    client.SetDebug(true)

    account, err := client.CreateAccount("Test", "Tester", "https://example.com")
    if err != nil {
        log.Fatalf("Failed to create account: %v", err)
    }
    fmt.Printf("Created account: %+v\n", account)
}
```

### Other Operations

#### Get Account Information

```go
accountInfo, err := client.GetAccountInfo(account.AccessToken, []string{"short_name", "author_name", "author_url", "page_count"})
if err != nil {
    log.Fatalf("Failed to get account info: %v", err)
}
fmt.Printf("Account info: %+v\n", accountInfo)
```

#### Create a Page

```go
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
page, err := client.CreatePage(account.AccessToken, "Test Page", content, "Tester", "https://example.com")
if err != nil {
    log.Fatalf("Failed to create page: %v", err)
}
fmt.Printf("Created page: %+v\n", page)
```

### More Examples

For more examples on how to use this client, please refer to the [telegraph_example.go](telegraph_example.go) file.

## Testing

This project uses pre-commit hooks to ensure code quality and consistency. To set up pre-commit hooks, run:

```sh
pip install pre-commit
pre-commit install
```

To run tests:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes or enhancements. Make sure to run pre-commit and tests before submitting your pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
