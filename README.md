# Unofficial Bank of Canada Valet API Go Module

This is an **unofficial Go module** for interacting with the [Bank of Canada Valet API](https://www.bankofcanada.ca/valet/docs). The module allows developers to programmatically access a wide range of economic data including exchange rates, interest rates, and financial time series published by the Bank of Canada.

## Features

- Query and fetch the latest economic data from the Bank of Canada.
- Easily integrate financial data into Go applications.

## Prerequisites

- **Go 1.22 or higher**

## Installation

Install the module by running:

```bash
go get github.com/KubeFinancial/bankofcanada-go
```

## Usage

```go
package main

import (
	"log"

	"github.com/KubeFinancial/bankofcanada-go/valet"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	// Fetch the list of available series
	response, err := valet.Api("/lists/series/json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
```

## Example Output
The above code will return a JSON object with the list of available series from the Bank of Canada Valet API.

## Additional Endpoints
Refer to the official [Valet API Documentation](https://www.bankofcanada.ca/valet/docs)  for more available endpoints and options.

## License and Attribution

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

You are free to use, modify, and distribute the code as needed. However, please note the following attribution:

Source: [Bank of Canada](https://www.bankofcanada.ca/terms/)

Content has been modified from its original form. This project is not endorsed by or affiliated with the Bank of Canada.
