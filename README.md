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
	"fmt"
	"log"

	"github.com/KubeFinancial/bankofcanada-go/valet"
)

func main() {
	log.SetPrefix("INFO\t")
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	// Fetch the list of available series
	apiResponse, err := valet.Api("/observations/group/FX_RATES_DAILY/json?recent=1")
	if err != nil {
		log.Fatal(err)
	}

	// Print the unmarshalled data
	for _, observation := range apiResponse.Observations {
		for pair, rate := range observation.Series {
			fmt.Println(observation.Date, pair, rate.Value)
		}
	}
}
```

## Example Output
The above code will return the most recent FX_RATES_DAILY data from the Bank of Canada Valet API.
```shell
DEBUG   2024/10/15 02:31:19.375883 GET https://www.bankofcanada.ca/valet/observations/group/FX_RATES_DAILY/json?recent=1 HTTP/1.1
DEBUG   2024/10/15 02:31:19.547594 200 OK
2019-12-31 FXMYRCAD 0.3175
2019-12-31 FXTHBCAD 0.04362
2019-12-31 FXVNDCAD 0.000056
2024-10-11 FXINRCAD 0.01636
2024-10-11 FXHKDCAD 0.1771
2024-10-11 FXSGDCAD 1.0546
2024-10-11 FXEURCAD 1.5056
2024-10-11 FXSEKCAD 0.1327
2024-10-11 FXMXNCAD 0.07099
2024-10-11 FXNOKCAD 0.1286
2024-10-11 FXPENCAD 0.3680
2024-10-11 FXRUBCAD 0.01436
2024-10-11 FXSARCAD 0.3665
2024-10-11 FXKRWCAD 0.001019
2024-10-11 FXIDRCAD 0.000088
2024-10-11 FXZARCAD 0.07903
2024-10-11 FXCHFCAD 1.6053
2024-10-11 FXBRLCAD 0.2447
2024-10-11 FXJPYCAD 0.009230
2024-10-11 FXUSDCAD 1.3761
2024-10-11 FXCNYCAD 0.1947
2024-10-11 FXTRYCAD 0.04010
2024-10-11 FXAUDCAD 0.9288
2024-10-11 FXNZDCAD 0.8403
2024-10-11 FXTWDCAD 0.04274
2024-10-11 FXGBPCAD 1.7988
```

## Additional Endpoints
Refer to the official [Valet API Documentation](https://www.bankofcanada.ca/valet/docs)  for more available endpoints and options.

## License and Attribution

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

You are free to use, modify, and distribute the code as needed. However, please note the following attribution:

Source: [Bank of Canada](https://www.bankofcanada.ca/terms/)

Content has been modified from its original form. This project is not endorsed by or affiliated with the Bank of Canada.
