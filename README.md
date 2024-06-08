# Telegraph API Client in Go

This is a Go client for the [Telegraph API](https://telegra.ph/api).

## Features

- Create and manage Telegraph accounts
- Create, edit, and retrieve Telegraph pages
- Get page views and list of pages

## Installation

To install the package, use the following command:

```sh
go get github.com/smirnoffmg/telegraph
```

## Usage

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

### Documentation

For detailed documentation, please refer to the [API Documentation](https://telegra.ph/api).

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
