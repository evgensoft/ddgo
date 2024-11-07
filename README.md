# ddgo

`ddgo` is a Go package that provides a simple interface to query DuckDuckGo's search engine and retrieve search results. It allows you to perform searches and obtain relevant information such as titles, snippets, and URLs of the results.

## Installation

To use the `ddgo` package, you need to have Go installed on your machine. You can install the package using the following command:

```bash
go get github.com/evgensoft/ddgo
```

## Usage

Here is a basic example of how to use the `ddgo` package to perform a search query:

```go
package main

import (
	"fmt"
	"log"

	"github.com/evgensoft/ddgo"
)

func main() {
	query := "Go programming language"
	maxResults := 5

	results, err := ddgo.Query(query, maxResults)
	if err != nil {
		log.Fatalf("Error querying DuckDuckGo: %v", err)
	}

	for _, result := range results {
		fmt.Printf("Title: %s\nInfo: %s\nRef: %s\n\n", result.Title, result.Info, result.Ref)
	}
}
```

### Function: Query

The `Query` function is the main function of the package. It takes a search query and the maximum number of results to return.

#### Parameters

- `query` (string): The search query string.
- `maxResult` (int): The maximum number of results to return.

#### Returns

- `([]Result, error)`: A slice of `Result` structs containing the search results and an error if any occurred.

### Result Struct

The `Result` struct holds the data for each search result:

- `Title` (string): The title of the search result.
- `Info` (string): A snippet of information about the search result.
- `Ref` (string): The URL of the search result.

## Error Handling

The `Query` function returns an error if there are issues with the HTTP request or if the response status code is not 200. Make sure to handle errors appropriately in your application.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contributing

Contributions are welcome! If you have suggestions for improvements or find bugs, please open an issue or submit a pull request.
